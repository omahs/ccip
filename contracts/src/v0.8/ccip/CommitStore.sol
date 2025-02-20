// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {TypeAndVersionInterface} from "../interfaces/TypeAndVersionInterface.sol";
import {ICommitStore} from "./interfaces/ICommitStore.sol";
import {IARM} from "./interfaces/IARM.sol";
import {IPriceRegistry} from "./interfaces/IPriceRegistry.sol";

import {OCR2Base} from "./ocr/OCR2Base.sol";
import {Internal} from "./libraries/Internal.sol";
import {MerkleMultiProof} from "./libraries/MerkleMultiProof.sol";

contract CommitStore is ICommitStore, TypeAndVersionInterface, OCR2Base {
  error StaleReport();
  error PausedError();
  error InvalidInterval(Interval interval);
  error InvalidRoot();
  error InvalidCommitStoreConfig();
  error BadARMSignal();
  error RootAlreadyCommitted();

  event Paused(address account);
  event Unpaused(address account);
  event ReportAccepted(CommitReport report);
  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event RootRemoved(bytes32 root);

  /// @notice Static commit store config
  struct StaticConfig {
    uint64 chainSelector; // -------┐  Destination chainSelector
    uint64 sourceChainSelector; // -┘  Source chainSelector
    address onRamp; // OnRamp address on the source chain
    address armProxy; // ARM proxy address
  }

  /// @notice Dynamic commit store config
  struct DynamicConfig {
    address priceRegistry; // Price registry address on the destination chain
  }

  /// @notice a sequenceNumber interval
  struct Interval {
    uint64 min; // ---┐ Minimum sequence number, inclusive
    uint64 max; // ---┘ Maximum sequence number, inclusive
  }

  /// @notice Report that is committed by the observing DON at the committing phase
  struct CommitReport {
    Internal.PriceUpdates priceUpdates;
    Interval interval;
    bytes32 merkleRoot;
  }

  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "CommitStore 1.2.0";
  // Chain ID of this chain
  uint64 internal immutable i_chainSelector;
  // Chain ID of the source chain
  uint64 internal immutable i_sourceChainSelector;
  // The onRamp address on the source chain
  address internal immutable i_onRamp;
  // The address of the arm proxy
  address internal immutable i_armProxy;

  // DYNAMIC CONFIG
  // The dynamic commitStore config
  DynamicConfig internal s_dynamicConfig;

  // STATE
  // The min sequence number expected for future messages
  uint64 private s_minSeqNr = 1;
  /// @dev The epoch and round of the last report
  uint40 private s_latestPriceEpochAndRound;
  /// @dev Whether this OnRamp is paused or not
  bool private s_paused = false;
  // merkleRoot => timestamp when received
  mapping(bytes32 merkleRoot => uint256 timestamp) private s_roots;

  /// @param staticConfig Containing the static part of the commitStore config
  /// @dev When instantiating OCR2Base we set UNIQUE_REPORTS to false, which means
  /// that we do not require 2f+1 signatures on a report, only f+1 to save gas. 2f+1 is required
  /// only if one must strictly ensure that for a given round there is only one valid report ever generated by
  /// the DON. In our case additional valid reports (i.e. approved by >= f+1 oracles) are not a problem, as they will
  /// will either be ignored (reverted as an invalid interval) or will be accepted as an additional valid price update.
  constructor(StaticConfig memory staticConfig) OCR2Base(false) {
    if (
      staticConfig.onRamp == address(0) ||
      staticConfig.chainSelector == 0 ||
      staticConfig.sourceChainSelector == 0 ||
      staticConfig.armProxy == address(0)
    ) revert InvalidCommitStoreConfig();

    i_chainSelector = staticConfig.chainSelector;
    i_sourceChainSelector = staticConfig.sourceChainSelector;
    i_onRamp = staticConfig.onRamp;
    i_armProxy = staticConfig.armProxy;
  }

  // ================================================================
  // |                        Verification                          |
  // ================================================================

  /// @notice Returns the next expected sequence number.
  /// @return the next expected sequenceNumber.
  function getExpectedNextSequenceNumber() external view returns (uint64) {
    return s_minSeqNr;
  }

  /// @notice Sets the minimum sequence number.
  /// @param minSeqNr The new minimum sequence number.
  function setMinSeqNr(uint64 minSeqNr) external onlyOwner {
    s_minSeqNr = minSeqNr;
  }

  /// @notice Returns the epoch and round of the last price update.
  /// @return the latest price epoch and round.
  function getLatestPriceEpochAndRound() public view returns (uint64) {
    return s_latestPriceEpochAndRound;
  }

  /// @notice Sets the latest epoch and round for price update.
  /// @param latestPriceEpochAndRound The new epoch and round for prices.
  function setLatestPriceEpochAndRound(uint40 latestPriceEpochAndRound) external onlyOwner {
    s_latestPriceEpochAndRound = latestPriceEpochAndRound;
  }

  /// @notice Returns the timestamp of a potentially previously committed merkle root.
  /// If the root was never committed 0 will be returned.
  /// @param root The merkle root to check the commit status for.
  /// @return the timestamp of the committed root or zero in the case that it was never
  /// committed.
  function getMerkleRoot(bytes32 root) external view returns (uint256) {
    return s_roots[root];
  }

  /// @notice Returns if a root is blessed or not.
  /// @param root The merkle root to check the blessing status for.
  /// @return whether the root is blessed or not.
  function isBlessed(bytes32 root) public view returns (bool) {
    return IARM(i_armProxy).isBlessed(IARM.TaggedRoot({commitStore: address(this), root: root}));
  }

  /// @notice Used by the owner in case an invalid sequence of roots has been
  /// posted and needs to be removed. The interval in the report is trusted.
  /// @param rootToReset The roots that will be reset. This function will only
  /// reset roots that are not blessed.
  function resetUnblessedRoots(bytes32[] calldata rootToReset) external onlyOwner {
    for (uint256 i = 0; i < rootToReset.length; ++i) {
      bytes32 root = rootToReset[i];
      if (!isBlessed(root)) {
        delete s_roots[root];
        emit RootRemoved(root);
      }
    }
  }

  /// @inheritdoc ICommitStore
  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata proofs,
    uint256 proofFlagBits
  ) external view override whenNotPaused returns (uint256 timestamp) {
    bytes32 root = MerkleMultiProof.merkleRoot(hashedLeaves, proofs, proofFlagBits);
    // Only return non-zero if present and blessed.
    if (!isBlessed(root)) {
      return 0;
    }
    return s_roots[root];
  }

  /// @inheritdoc OCR2Base
  /// @dev A commitReport can have two distinct parts (batched together to amortize the cost of checking sigs):
  /// 1. Price updates
  /// 2. A merkle root and sequence number interval
  /// Both have their own, separate, staleness checks, with price updates using the epoch and round
  /// number of the latest price update. The merkle root checks for staleness based on the seqNums.
  /// They need to be separate because a price report for round t+2 might be included before a report
  /// containing a merkle root for round t+1. This merkle root report for round t+1 is still valid
  /// and should not be rejected. When a report with a stale root but valid price updates is submitted,
  /// we are OK to revert to preserve the invariant that we always revert on invalid sequence number ranges.
  /// If that happens, prices will be updates in later rounds.
  function _report(bytes calldata encodedReport, uint40 epochAndRound) internal override whenNotPaused whenHealthy {
    CommitReport memory report = abi.decode(encodedReport, (CommitReport));

    // Check if the report contains price updates
    if (report.priceUpdates.tokenPriceUpdates.length > 0 || report.priceUpdates.destChainSelector != 0) {
      // Check for price staleness based on the epoch and round
      if (s_latestPriceEpochAndRound < epochAndRound) {
        // If prices are not stale, update the latest epoch and round
        s_latestPriceEpochAndRound = epochAndRound;
        // And update the prices in the price registry
        IPriceRegistry(s_dynamicConfig.priceRegistry).updatePrices(report.priceUpdates);

        // If there is no root, the report only contained fee updated and
        // we return to not revert on the empty root check below.
        if (report.merkleRoot == bytes32(0)) return;
      } else {
        // If prices are stale and the report doesn't contain a root, this report
        // does not have any valid information and we revert.
        // If it does contain a merkle root, continue to the root checking section.
        if (report.merkleRoot == bytes32(0)) revert StaleReport();
      }
    }

    // If we reached this section, the report should contain a valid root
    if (s_minSeqNr != report.interval.min || report.interval.min > report.interval.max)
      revert InvalidInterval(report.interval);

    if (report.merkleRoot == bytes32(0)) revert InvalidRoot();
    // Disallow duplicate roots as that would reset the timestamp and
    // delay potential manual execution.
    if (s_roots[report.merkleRoot] != 0) revert RootAlreadyCommitted();

    s_minSeqNr = report.interval.max + 1;
    s_roots[report.merkleRoot] = block.timestamp;
    emit ReportAccepted(report);
  }

  // ================================================================
  // |                           Config                             |
  // ================================================================

  /// @notice Returns the static commit store config.
  /// @return the configuration.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return
      StaticConfig({
        chainSelector: i_chainSelector,
        sourceChainSelector: i_sourceChainSelector,
        onRamp: i_onRamp,
        armProxy: i_armProxy
      });
  }

  /// @notice Returns the dynamic commit store config.
  /// @return the configuration.
  function getDynamicConfig() external view returns (DynamicConfig memory) {
    return s_dynamicConfig;
  }

  /// @notice Sets the dynamic config. This function is called during `setOCR2Config` flow
  function _beforeSetConfig(bytes memory onchainConfig) internal override {
    DynamicConfig memory dynamicConfig = abi.decode(onchainConfig, (DynamicConfig));

    if (dynamicConfig.priceRegistry == address(0)) revert InvalidCommitStoreConfig();

    s_dynamicConfig = dynamicConfig;
    // When the OCR config changes, we reset the price epoch and round
    // since epoch and rounds are scoped per config digest.
    // Note that s_minSeqNr/roots do not need to be reset as the roots persist
    // across reconfigurations and are de-duplicated separately.
    s_latestPriceEpochAndRound = 0;

    emit ConfigSet(
      StaticConfig({
        chainSelector: i_chainSelector,
        sourceChainSelector: i_sourceChainSelector,
        onRamp: i_onRamp,
        armProxy: i_armProxy
      }),
      dynamicConfig
    );
  }

  // ================================================================
  // |                        Access and ARM                        |
  // ================================================================

  /// @notice Single function to check the status of the commitStore.
  function isUnpausedAndARMHealthy() external view returns (bool) {
    return !IARM(i_armProxy).isCursed() && !s_paused;
  }

  /// @notice Support querying whether health checker is healthy.
  function isARMHealthy() external view returns (bool) {
    return !IARM(i_armProxy).isCursed();
  }

  /// @notice Ensure that the ARM has not emitted a bad signal, and that the latest heartbeat is not stale.
  modifier whenHealthy() {
    if (IARM(i_armProxy).isCursed()) revert BadARMSignal();
    _;
  }

  /// @notice Modifier to make a function callable only when the contract is not paused.
  modifier whenNotPaused() {
    if (paused()) revert PausedError();
    _;
  }

  /// @notice Returns true if the contract is paused, and false otherwise.
  function paused() public view returns (bool) {
    return s_paused;
  }

  /// @notice Pause the contract
  /// @dev only callable by the owner
  function pause() external onlyOwner {
    s_paused = true;
    emit Paused(msg.sender);
  }

  /// @notice Unpause the contract
  /// @dev only callable by the owner
  function unpause() external onlyOwner {
    s_paused = false;
    emit Unpaused(msg.sender);
  }
}
