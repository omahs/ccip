// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_v2plus_upgraded_version

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type VRFCoordinatorV2PlusUpgradedVersionFeeConfig struct {
	FulfillmentFlatFeeLinkPPM uint32
	FulfillmentFlatFeeEthPPM  uint32
}

type VRFCoordinatorV2PlusUpgradedVersionRequestCommitment struct {
	BlockNum         uint64
	SubId            *big.Int
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
	ExtraArgs        []byte
}

type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

type VRFV2PlusClientRandomWordsRequest struct {
	KeyHash              [32]byte
	SubId                *big.Int
	RequestConfirmations uint16
	CallbackGasLimit     uint32
	NumWords             uint32
	ExtraArgs            []byte
}

var VRFCoordinatorV2PlusUpgradedVersionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blockhashStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"BlockhashNotInStore\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToTransferLink\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"InsufficientGasForConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transferredValue\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"expectedValue\",\"type\":\"uint96\"}],\"name\":\"InvalidNativeBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"have\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"min\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"max\",\"type\":\"uint16\"}],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"requestVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"expectedVersion\",\"type\":\"uint8\"}],\"name\":\"InvalidVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoCorrespondingRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"NoSuchProvingKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"NumWordsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubscriptionIDCollisionFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structVRFCoordinatorV2PlusUpgradedVersion.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthFundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"MigrationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountEth\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldEthBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEthBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFundedWithEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCKHASH_STORE\",\"outputs\":[{\"internalType\":\"contractBlockhashStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK_ETH_FEED\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFCoordinatorV2PlusUpgradedVersion.RequestCommitment\",\"name\":\"rc\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"fundSubscriptionWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"getActiveSubscriptionIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"ethBalance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedData\",\"type\":\"bytes\"}],\"name\":\"onMigration\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverEthFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"registerMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFV2PlusClient.RandomWordsRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_config\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_currentSubNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fallbackWeiPerUnitLink\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_feeConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_provingKeyHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_provingKeys\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requestCommitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalEthBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"internalType\":\"structVRFCoordinatorV2PlusUpgradedVersion.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkEthFeed\",\"type\":\"address\"}],\"name\":\"setLINKAndLINKETHFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620060b9380380620060b9833981016040819052620000349162000183565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000d7565b50505060601b6001600160601b031916608052620001b5565b6001600160a01b038116331415620001325760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019657600080fd5b81516001600160a01b0381168114620001ae57600080fd5b9392505050565b60805160601c615ede620001db600039600081816104c401526136f60152615ede6000f3fe6080604052600436106102015760003560e01c806201229114610206578063043bd6ae14610233578063088070f5146102575780630ae09540146102d757806315c48b84146102f95780631b6b6d2314610321578063294daa491461034e578063330987b31461036a578063405b84fa146103a257806340d6bb82146103c257806341af6c87146103ed57806346d8d4861461041d57806357133e641461043d5780635d06b4ab1461045d57806364d51a2a1461047d57806366316d8d14610492578063689c4517146104b25780636b6feccc146104e65780636f64f03f1461051c57806379ba50971461053c57806386fe91c7146105515780638da5cb5b146105715780639b1c385e1461058f5780639d40a6fd146105af578063a21a23e4146105dc578063a4c0ed36146105f1578063a8cb447b14610611578063aa433aff14610631578063ad17836114610651578063aefb212f14610671578063b08c87951461069e578063b2a7cac5146106be578063bec4c08c146106de578063caf70c4a146106fe578063cb6317971461071e578063ce3f47191461073e578063d98e620e14610751578063da2f261014610771578063dac83d29146107a7578063dc311dd3146107c7578063e72f6e30146107f8578063e8509bff14610818578063e95704bd1461082b578063ee9d2d3814610852578063f2fde38b1461087f575b600080fd5b34801561021257600080fd5b5061021b61089f565b60405161022a939291906159d4565b60405180910390f35b34801561023f57600080fd5b5061024960115481565b60405190815260200161022a565b34801561026357600080fd5b50600d5461029f9061ffff81169063ffffffff62010000820481169160ff600160301b82041691600160381b8204811691600160581b90041685565b6040805161ffff909616865263ffffffff9485166020870152921515928501929092528216606084015216608082015260a00161022a565b3480156102e357600080fd5b506102f76102f2366004615655565b61091b565b005b34801561030557600080fd5b5061030e60c881565b60405161ffff909116815260200161022a565b34801561032d57600080fd5b50600254610341906001600160a01b031681565b60405161022a9190615878565b34801561035a57600080fd5b506040516001815260200161022a565b34801561037657600080fd5b5061038a6103853660046153c3565b6109e9565b6040516001600160601b03909116815260200161022a565b3480156103ae57600080fd5b506102f76103bd366004615655565b610ec4565b3480156103ce57600080fd5b506103d86101f481565b60405163ffffffff909116815260200161022a565b3480156103f957600080fd5b5061040d610408366004615305565b6112af565b604051901515815260200161022a565b34801561042957600080fd5b506102f76104383660046151c7565b611450565b34801561044957600080fd5b506102f76104583660046151fc565b6115cd565b34801561046957600080fd5b506102f76104783660046151aa565b61162d565b34801561048957600080fd5b5061030e606481565b34801561049e57600080fd5b506102f76104ad3660046151c7565b6116e4565b3480156104be57600080fd5b506103417f000000000000000000000000000000000000000000000000000000000000000081565b3480156104f257600080fd5b5060125461050e9063ffffffff80821691600160201b90041682565b60405161022a929190615b56565b34801561052857600080fd5b506102f7610537366004615235565b6118ac565b34801561054857600080fd5b506102f76119b3565b34801561055d57600080fd5b50600a5461038a906001600160601b031681565b34801561057d57600080fd5b506000546001600160a01b0316610341565b34801561059b57600080fd5b506102496105aa3660046154a0565b611a5d565b3480156105bb57600080fd5b506007546105cf906001600160401b031681565b60405161022a9190615b6d565b3480156105e857600080fd5b50610249611dcd565b3480156105fd57600080fd5b506102f761060c366004615271565b61201b565b34801561061d57600080fd5b506102f761062c3660046151aa565b6121b8565b34801561063d57600080fd5b506102f761064c366004615305565b6122c4565b34801561065d57600080fd5b50600354610341906001600160a01b031681565b34801561067d57600080fd5b5061069161068c36600461567a565b612327565b60405161022a91906158ef565b3480156106aa57600080fd5b506102f76106b93660046155b7565b612428565b3480156106ca57600080fd5b506102f76106d9366004615305565b61259c565b3480156106ea57600080fd5b506102f76106f9366004615655565b6126cc565b34801561070a57600080fd5b506102496107193660046152cc565b612863565b34801561072a57600080fd5b506102f7610739366004615655565b612893565b6102f761074c366004615337565b612b80565b34801561075d57600080fd5b5061024961076c366004615305565b612e91565b34801561077d57600080fd5b5061034161078c366004615305565b600e602052600090815260409020546001600160a01b031681565b3480156107b357600080fd5b506102f76107c2366004615655565b612eb2565b3480156107d357600080fd5b506107e76107e2366004615305565b612fc2565b60405161022a959493929190615b81565b34801561080457600080fd5b506102f76108133660046151aa565b6130bd565b6102f7610826366004615305565b613298565b34801561083757600080fd5b50600a5461038a90600160601b90046001600160601b031681565b34801561085e57600080fd5b5061024961086d366004615305565b60106020526000908152604090205481565b34801561088b57600080fd5b506102f761089a3660046151aa565b6133d0565b600d54600f805460408051602080840282018101909252828152600094859460609461ffff8316946201000090930463ffffffff1693919283919083018282801561090957602002820191906000526020600020905b8154815260200190600101908083116108f5575b50505050509050925092509250909192565b60008281526005602052604090205482906001600160a01b03168061095357604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146109875780604051636c51fda960e11b815260040161097e9190615878565b60405180910390fd5b600d54600160301b900460ff16156109b25760405163769dd35360e11b815260040160405180910390fd5b6109bb846112af565b156109d957604051631685ecdd60e31b815260040160405180910390fd5b6109e384846133e1565b50505050565b600d54600090600160301b900460ff1615610a175760405163769dd35360e11b815260040160405180910390fd5b60005a90506000610a28858561359c565b90506000846060015163ffffffff166001600160401b03811115610a4e57610a4e615e98565b604051908082528060200260200182016040528015610a77578160200160208202803683370190505b50905060005b856060015163ffffffff16811015610aee57826040015181604051602001610aa6929190615902565b6040516020818303038152906040528051906020012060001c828281518110610ad157610ad1615e82565b602090810291909101015280610ae681615dea565b915050610a7d565b5060208083018051600090815260109092526040808320839055905190518291631fe543e360e01b91610b2691908690602401615a5e565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252600d805460ff60301b1916600160301b179055908801516080890151919250600091610b8b9163ffffffff16908461380f565b600d805460ff60301b19169055602089810151600090815260069091526040902054909150600160c01b90046001600160401b0316610bcb816001615cfb565b6020808b0151600090815260069091526040812080546001600160401b0393909316600160c01b026001600160c01b039093169290921790915560a08a01518051610c1890600190615d7b565b81518110610c2857610c28615e82565b602091010151600d5460f89190911c6001149150600090610c59908a90600160581b900463ffffffff163a8561385d565b90508115610d62576020808c01516000908152600690915260409020546001600160601b03808316600160601b909204161015610ca957604051631e9acf1760e31b815260040160405180910390fd5b60208b81015160009081526006909152604090208054829190600c90610ce0908490600160601b90046001600160601b0316615d92565b82546101009290920a6001600160601b0381810219909316918316021790915589516000908152600e60209081526040808320546001600160a01b03168352600c909152812080548594509092610d3991859116615d26565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550610e4e565b6020808c01516000908152600690915260409020546001600160601b0380831691161015610da357604051631e9acf1760e31b815260040160405180910390fd5b6020808c015160009081526006909152604081208054839290610dd09084906001600160601b0316615d92565b82546101009290920a6001600160601b0381810219909316918316021790915589516000908152600e60209081526040808320546001600160a01b03168352600b909152812080548594509092610e2991859116615d26565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b8a6020015188602001517f49580fdfd9497e1ed5c1b1cec0495087ae8e3f1267470ec2fb015db32e3d6aa78a604001518488604051610eab939291909283526001600160601b039190911660208301521515604082015260600190565b60405180910390a3985050505050505050505b92915050565b600d54600160301b900460ff1615610eef5760405163769dd35360e11b815260040160405180910390fd5b610ef8816138ac565b610f175780604051635428d44960e01b815260040161097e9190615878565b600080600080610f2686612fc2565b945094505093509350336001600160a01b0316826001600160a01b031614610f895760405162461bcd60e51b81526020600482015260166024820152752737ba1039bab139b1b934b83a34b7b71037bbb732b960511b604482015260640161097e565b610f92866112af565b15610fd85760405162461bcd60e51b815260206004820152601660248201527550656e64696e6720726571756573742065786973747360501b604482015260640161097e565b60006040518060c00160405280610fed600190565b60ff168152602001888152602001846001600160a01b03168152602001838152602001866001600160601b03168152602001856001600160601b031681525090506000816040516020016110419190615941565b604051602081830303815290604052905061105b88613916565b505060405163ce3f471960e01b81526001600160a01b0388169063ce3f4719906001600160601b0388169061109490859060040161592e565b6000604051808303818588803b1580156110ad57600080fd5b505af11580156110c1573d6000803e3d6000fd5b50506002546001600160a01b0316158015935091506110ea905057506001600160601b03861615155b156111b45760025460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb90611121908a908a906004016158bf565b602060405180830381600087803b15801561113b57600080fd5b505af115801561114f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061117391906152e8565b6111b45760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b604482015260640161097e565b600d805460ff60301b1916600160301b17905560005b835181101561125d578381815181106111e5576111e5615e82565b60200260200101516001600160a01b0316638ea98117896040518263ffffffff1660e01b81526004016112189190615878565b600060405180830381600087803b15801561123257600080fd5b505af1158015611246573d6000803e3d6000fd5b50505050808061125590615dea565b9150506111ca565b50600d805460ff60301b191690556040517fd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be41879061129d9089908b9061588c565b60405180910390a15050505050505050565b6000818152600560209081526040808320815160608101835281546001600160a01b039081168252600183015416818501526002820180548451818702810187018652818152879693958601939092919083018282801561133957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161131b575b505050505081525050905060005b8160400151518110156114465760005b600f548110156114335760006113fc600f838154811061137957611379615e82565b90600052602060002001548560400151858151811061139a5761139a615e82565b60200260200101518860046000896040015189815181106113bd576113bd615e82565b6020908102919091018101516001600160a01b0316825281810192909252604090810160009081208d82529092529020546001600160401b0316613b64565b50600081815260106020526040902054909150156114205750600195945050505050565b508061142b81615dea565b915050611357565b508061143e81615dea565b915050611347565b5060009392505050565b600d54600160301b900460ff161561147b5760405163769dd35360e11b815260040160405180910390fd5b336000908152600c60205260409020546001600160601b03808316911610156114b757604051631e9acf1760e31b815260040160405180910390fd5b336000908152600c6020526040812080548392906114df9084906001600160601b0316615d92565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600a600c8282829054906101000a90046001600160601b03166115279190615d92565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506000826001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d80600081146115a1576040519150601f19603f3d011682016040523d82523d6000602084013e6115a6565b606091505b50509050806115c857604051630dcf35db60e41b815260040160405180910390fd5b505050565b6115d5613bed565b6002546001600160a01b0316156115ff57604051631688c53760e11b815260040160405180910390fd5b600280546001600160a01b039384166001600160a01b03199182161790915560038054929093169116179055565b611635613bed565b61163e816138ac565b1561165e578060405163ac8a27ef60e01b815260040161097e9190615878565b601380546001810182556000919091527f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a0900180546001600160a01b0319166001600160a01b0383161790556040517fb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625906116d9908390615878565b60405180910390a150565b600d54600160301b900460ff161561170f5760405163769dd35360e11b815260040160405180910390fd5b6002546001600160a01b03166117385760405163c1f0c0a160e01b815260040160405180910390fd5b336000908152600b60205260409020546001600160601b038083169116101561177457604051631e9acf1760e31b815260040160405180910390fd5b336000908152600b60205260408120805483929061179c9084906001600160601b0316615d92565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600a60008282829054906101000a90046001600160601b03166117e49190615d92565b82546001600160601b039182166101009390930a92830291909202199091161790555060025460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb9061183990859085906004016158bf565b602060405180830381600087803b15801561185357600080fd5b505af1158015611867573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061188b91906152e8565b6118a857604051631e9acf1760e31b815260040160405180910390fd5b5050565b6118b4613bed565b6040805180820182526000916118e3919084906002908390839080828437600092019190915250612863915050565b6000818152600e60205260409020549091506001600160a01b03161561191f57604051634a0b8fa760e01b81526004810182905260240161097e565b6000818152600e6020908152604080832080546001600160a01b0319166001600160a01b038816908117909155600f805460018101825594527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802909301849055518381527fe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8910160405180910390a2505050565b6001546001600160a01b03163314611a065760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b604482015260640161097e565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600d54600090600160301b900460ff1615611a8b5760405163769dd35360e11b815260040160405180910390fd5b6020808301356000908152600590915260409020546001600160a01b0316611ac657604051630fb532db60e11b815260040160405180910390fd5b3360009081526004602090815260408083208583013584529091529020546001600160401b031680611b13578260200135336040516379bfd40160e01b815260040161097e929190615a33565b600d5461ffff16611b2a606085016040860161559c565b61ffff161080611b4d575060c8611b47606085016040860161559c565b61ffff16115b15611b8757611b62606084016040850161559c565b600d5460405163539c34bb60e11b815261097e929161ffff169060c8906004016159b6565b600d5462010000900463ffffffff16611ba6608085016060860161569c565b63ffffffff161115611bec57611bc2608084016060850161569c565b600d54604051637aebf00f60e11b815261097e929162010000900463ffffffff1690600401615b56565b6101f4611bff60a085016080860161569c565b63ffffffff161115611c3957611c1b60a084016080850161569c565b6101f46040516311ce1afb60e21b815260040161097e929190615b56565b6000611c46826001615cfb565b9050600080611c5c863533602089013586613b64565b90925090506000611c78611c7360a0890189615bd6565b613c42565b90506000611c8582613cbf565b905083611c90613d30565b60208a0135611ca560808c0160608d0161569c565b611cb560a08d0160808e0161569c565b3386604051602001611ccd9796959493929190615ab6565b604051602081830303815290604052805190602001206010600086815260200190815260200160002081905550336001600160a01b0316886020013589600001357feb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e87878d6040016020810190611d44919061559c565b8e6060016020810190611d57919061569c565b8f6080016020810190611d6a919061569c565b89604051611d7d96959493929190615a77565b60405180910390a45050336000908152600460209081526040808320898301358452909152902080546001600160401b0319166001600160401b039490941693909317909255925050505b919050565b600d54600090600160301b900460ff1615611dfb5760405163769dd35360e11b815260040160405180910390fd5b600033611e09600143615d7b565b600754604051606093841b6001600160601b03199081166020830152924060348201523090931b909116605483015260c01b6001600160c01b031916606882015260700160408051601f198184030181529190528051602090910120600780549192506001600160401b03909116906000611e8383615e05565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550506000806001600160401b03811115611ec257611ec2615e98565b604051908082528060200260200182016040528015611eeb578160200160208202803683370190505b506040805160608082018352600080835260208084018281528486018381528984526006835286842095518654925191516001600160601b039182166001600160c01b031990941693909317600160601b9190921602176001600160c01b0316600160c01b6001600160401b039092169190910217909355835191820184523382528183018181528285018681528883526005855294909120825181546001600160a01b03199081166001600160a01b039283161783559251600183018054909416911617909155925180519495509093611fcc9260028501920190614e1b565b50611fdc91506008905083613dc9565b50817f1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d3360405161200d9190615878565b60405180910390a250905090565b600d54600160301b900460ff16156120465760405163769dd35360e11b815260040160405180910390fd5b6002546001600160a01b03163314612071576040516344b0e3c360e01b815260040160405180910390fd5b6020811461209257604051638129bbcd60e01b815260040160405180910390fd5b60006120a082840184615305565b6000818152600560205260409020549091506001600160a01b03166120d857604051630fb532db60e11b815260040160405180910390fd5b600081815260066020526040812080546001600160601b0316918691906120ff8385615d26565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555084600a60008282829054906101000a90046001600160601b03166121479190615d26565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a82878461219a9190615ce3565b6040516121a8929190615902565b60405180910390a2505050505050565b6121c0613bed565b600a544790600160601b90046001600160601b0316818111156121fa5780826040516354ced18160e11b815260040161097e929190615902565b818110156115c857600061220e8284615d7b565b90506000846001600160a01b03168260405160006040518083038185875af1925050503d806000811461225d576040519150601f19603f3d011682016040523d82523d6000602084013e612262565b606091505b505090508061228457604051630dcf35db60e41b815260040160405180910390fd5b7f879c9ea2b9d5345b84ccd12610b032602808517cebdb795007f3dcb4df37731785836040516122b592919061588c565b60405180910390a15050505050565b6122cc613bed565b6000818152600560205260409020546001600160a01b031661230157604051630fb532db60e11b815260040160405180910390fd5b6000818152600560205260409020546123249082906001600160a01b03166133e1565b50565b606060006123356008613dd5565b905080841061235757604051631390f2a160e01b815260040160405180910390fd5b60006123638486615ce3565b905081811180612371575083155b61237b578061237d565b815b9050600061238b8683615d7b565b6001600160401b038111156123a2576123a2615e98565b6040519080825280602002602001820160405280156123cb578160200160208202803683370190505b50905060005b815181101561241e576123ef6123e78883615ce3565b600890613ddf565b82828151811061240157612401615e82565b60209081029190910101528061241681615dea565b9150506123d1565b5095945050505050565b612430613bed565b60c861ffff8716111561245d57858660c860405163539c34bb60e11b815260040161097e939291906159b6565b60008213612481576040516321ea67b360e11b81526004810183905260240161097e565b6040805160a0808201835261ffff891680835263ffffffff89811660208086018290526000868801528a831660608088018290528b85166080988901819052600d805465ffffffffffff1916881762010000870217600160301b600160781b031916600160381b850263ffffffff60581b191617600160581b83021790558a51601280548d8701519289166001600160401b031990911617600160201b92891692909202919091179081905560118d90558a519788528785019590955298860191909152840196909652938201879052838116928201929092529190921c90911660c08201527f777357bb93f63d088f18112d3dba38457aec633eb8f1341e1d418380ad328e789060e00160405180910390a1505050505050565b600d54600160301b900460ff16156125c75760405163769dd35360e11b815260040160405180910390fd5b6000818152600560205260409020546001600160a01b03166125fc57604051630fb532db60e11b815260040160405180910390fd5b6000818152600560205260409020600101546001600160a01b03163314612653576000818152600560205260409081902060010154905163d084e97560e01b815261097e916001600160a01b031690600401615878565b6000818152600560205260409081902080546001600160a01b031980821633908117845560019093018054909116905591516001600160a01b039092169183917fd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c9386916126c09185916158a5565b60405180910390a25050565b60008281526005602052604090205482906001600160a01b03168061270457604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b0382161461272f5780604051636c51fda960e11b815260040161097e9190615878565b600d54600160301b900460ff161561275a5760405163769dd35360e11b815260040160405180910390fd5b6000848152600560205260409020600201546064141561278d576040516305a48e0f60e01b815260040160405180910390fd5b6001600160a01b03831660009081526004602090815260408083208784529091529020546001600160401b0316156127c4576109e3565b6001600160a01b0383166000818152600460209081526040808320888452825280832080546001600160401b031916600190811790915560058352818420600201805491820181558452919092200180546001600160a01b0319169092179091555184907f1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e190612855908690615878565b60405180910390a250505050565b60008160405160200161287691906158e1565b604051602081830303815290604052805190602001209050919050565b60008281526005602052604090205482906001600160a01b0316806128cb57604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146128f65780604051636c51fda960e11b815260040161097e9190615878565b600d54600160301b900460ff16156129215760405163769dd35360e11b815260040160405180910390fd5b61292a846112af565b1561294857604051631685ecdd60e31b815260040160405180910390fd5b6001600160a01b03831660009081526004602090815260408083208784529091529020546001600160401b03166129965783836040516379bfd40160e01b815260040161097e929190615a33565b6000848152600560209081526040808320600201805482518185028101850190935280835291929091908301828280156129f957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116129db575b50505050509050600060018251612a109190615d7b565b905060005b8251811015612b1c57856001600160a01b0316838281518110612a3a57612a3a615e82565b60200260200101516001600160a01b03161415612b0a576000838381518110612a6557612a65615e82565b6020026020010151905080600560008a81526020019081526020016000206002018381548110612a9757612a97615e82565b600091825260208083209190910180546001600160a01b0319166001600160a01b039490941693909317909255898152600590915260409020600201805480612ae257612ae2615e6c565b600082815260209020810160001990810180546001600160a01b031916905501905550612b1c565b80612b1481615dea565b915050612a15565b506001600160a01b03851660009081526004602090815260408083208984529091529081902080546001600160401b03191690555186907f32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a7906121a8908890615878565b6000612b8e828401846154da565b9050806000015160ff16600114612bc757805160405163237d181f60e21b815260ff90911660048201526001602482015260440161097e565b8060a001516001600160601b03163414612c0b5760a08101516040516306acf13560e41b81523460048201526001600160601b03909116602482015260440161097e565b6020808201516000908152600590915260409020546001600160a01b031615612c47576040516326afa43560e11b815260040160405180910390fd5b60005b816060015151811015612ce75760016004600084606001518481518110612c7357612c73615e82565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060008460200151815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b031602179055508080612cdf90615dea565b915050612c4a565b50604080516060808201835260808401516001600160601b03908116835260a0850151811660208085019182526000858701818152828901805183526006845288832097518854955192516001600160401b0316600160c01b026001600160c01b03938816600160601b026001600160c01b0319909716919097161794909417169390931790945584518084018652868601516001600160a01b03908116825281860184815294880151828801908152925184526005865295909220825181549087166001600160a01b0319918216178255935160018201805491909716941693909317909455925180519192612de692600285019290910190614e1b565b5050506080810151600a8054600090612e099084906001600160601b0316615d26565b92506101000a8154816001600160601b0302191690836001600160601b031602179055508060a00151600a600c8282829054906101000a90046001600160601b0316612e559190615d26565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506109e381602001516008613dc990919063ffffffff16565b600f8181548110612ea157600080fd5b600091825260209091200154905081565b60008281526005602052604090205482906001600160a01b031680612eea57604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b03821614612f155780604051636c51fda960e11b815260040161097e9190615878565b600d54600160301b900460ff1615612f405760405163769dd35360e11b815260040160405180910390fd5b6000848152600560205260409020600101546001600160a01b038481169116146109e3576000848152600560205260409081902060010180546001600160a01b0319166001600160a01b0386161790555184907f21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a19061285590339087906158a5565b6000818152600560205260408120548190819081906060906001600160a01b031661300057604051630fb532db60e11b815260040160405180910390fd5b60008681526006602090815260408083205460058352928190208054600290910180548351818602810186019094528084526001600160601b0380871696600160601b810490911695600160c01b9091046001600160401b0316946001600160a01b03909416939183918301828280156130a357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613085575b505050505090509450945094509450945091939590929450565b6130c5613bed565b6002546001600160a01b03166130ee5760405163c1f0c0a160e01b815260040160405180910390fd5b6002546040516370a0823160e01b81526000916001600160a01b0316906370a082319061311f903090600401615878565b60206040518083038186803b15801561313757600080fd5b505afa15801561314b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061316f919061531e565b600a549091506001600160601b0316818111156131a35780826040516354ced18160e11b815260040161097e929190615902565b818110156115c85760006131b78284615d7b565b60025460405163a9059cbb60e01b81529192506001600160a01b03169063a9059cbb906131ea908790859060040161588c565b602060405180830381600087803b15801561320457600080fd5b505af1158015613218573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061323c91906152e8565b61325957604051631f01ff1360e21b815260040160405180910390fd5b7f59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600848260405161328a92919061588c565b60405180910390a150505050565b600d54600160301b900460ff16156132c35760405163769dd35360e11b815260040160405180910390fd5b6000818152600560205260409020546001600160a01b03166132f857604051630fb532db60e11b815260040160405180910390fd5b60008181526006602052604090208054600160601b90046001600160601b0316903490600c6133278385615d26565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555034600a600c8282829054906101000a90046001600160601b031661336f9190615d26565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f3f1ddc3ab1bdb39001ad76ca51a0e6f57ce6627c69f251d1de41622847721cde8234846133c29190615ce3565b6040516126c0929190615902565b6133d8613bed565b61232481613deb565b6000806133ed84613916565b60025491935091506001600160a01b03161580159061341457506001600160601b03821615155b156134c35760025460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb906134549086906001600160601b0387169060040161588c565b602060405180830381600087803b15801561346e57600080fd5b505af1158015613482573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134a691906152e8565b6134c357604051631e9acf1760e31b815260040160405180910390fd5b6000836001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d8060008114613519576040519150601f19603f3d011682016040523d82523d6000602084013e61351e565b606091505b505090508061354057604051630dcf35db60e41b815260040160405180910390fd5b604080516001600160a01b03861681526001600160601b038581166020830152841681830152905186917f8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4919081900360600190a25050505050565b604080516060810182526000808252602082018190529181019190915260006135c88460000151612863565b6000818152600e60205260409020549091506001600160a01b03168061360457604051631dfd6e1360e21b81526004810183905260240161097e565b600082866080015160405160200161361d929190615902565b60408051601f198184030181529181528151602092830120600081815260109093529120549091508061366357604051631b44092560e11b815260040160405180910390fd5b85516020808801516040808a015160608b015160808c015160a08d01519351613692978a979096959101615b02565b6040516020818303038152906040528051906020012081146136c75760405163354a450b60e21b815260040160405180910390fd5b60006136d68760000151613e8f565b90508061379d578651604051631d2827a760e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163e9413d389161372a9190600401615b6d565b60206040518083038186803b15801561374257600080fd5b505afa158015613756573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061377a919061531e565b90508061379d57865160405163175dadad60e01b815261097e9190600401615b6d565b60008860800151826040516020016137bf929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c905060006137e68a83613f82565b604080516060810182529889526020890196909652948701949094525093979650505050505050565b60005a61138881101561382157600080fd5b61138881039050846040820482031161383957600080fd5b50823b61384557600080fd5b60008083516020850160008789f190505b9392505050565b6000811561388a576012546138839086908690600160201b900463ffffffff1686613fed565b90506138a4565b6012546138a1908690869063ffffffff1686614057565b90505b949350505050565b6000805b60135481101561390d57826001600160a01b0316601382815481106138d7576138d7615e82565b6000918252602090912001546001600160a01b031614156138fb5750600192915050565b8061390581615dea565b9150506138b0565b50600092915050565b6000818152600560209081526040808320815160608101835281546001600160a01b039081168252600183015416818501526002820180548451818702810187018652818152879687969495948601939192908301828280156139a257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613984575b505050919092525050506000858152600660209081526040808320815160608101835290546001600160601b03808216808452600160601b8304909116948301859052600160c01b9091046001600160401b0316928201929092529096509094509192505b826040015151811015613a7e576004600084604001518381518110613a2e57613a2e615e82565b6020908102919091018101516001600160a01b031682528181019290925260409081016000908120898252909252902080546001600160401b031916905580613a7681615dea565b915050613a07565b50600085815260056020526040812080546001600160a01b03199081168255600182018054909116905590613ab66002830182614e80565b5050600085815260066020526040812055613ad2600886614144565b50600a8054859190600090613af19084906001600160601b0316615d92565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555082600a600c8282829054906101000a90046001600160601b0316613b399190615d92565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505050915091565b60408051602081018690526001600160a01b03851691810191909152606081018390526001600160401b03821660808201526000908190819060a00160408051601f198184030181529082905280516020918201209250613bc9918991849101615902565b60408051808303601f19018152919052805160209091012097909650945050505050565b6000546001600160a01b03163314613c405760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015260640161097e565b565b60408051602081019091526000815281613c6b5750604080516020810190915260008152610ebe565b63125fa26760e31b613c7d8385615dba565b6001600160e01b03191614613ca557604051632923fee760e11b815260040160405180910390fd5b613cb28260048186615cb9565b8101906138569190615378565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa82604051602401613cf891511515815260200190565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915292915050565b60004661a4b1811480613d45575062066eed81145b15613dc25760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b158015613d8457600080fd5b505afa158015613d98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613dbc919061531e565b91505090565b4391505090565b60006138568383614150565b6000610ebe825490565b6000613856838361419f565b6001600160a01b038116331415613e3e5760405162461bcd60e51b815260206004820152601760248201527621b0b73737ba103a3930b739b332b9103a379039b2b63360491b604482015260640161097e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60004661a4b1811480613ea4575062066eed81145b80613eb1575062066eee81145b15613f7357610100836001600160401b0316613ecb613d30565b613ed59190615d7b565b1180613ef15750613ee4613d30565b836001600160401b031610155b15613eff5750600092915050565b6040516315a03d4160e11b8152606490632b407a8290613f23908690600401615b6d565b60206040518083038186803b158015613f3b57600080fd5b505afa158015613f4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613856919061531e565b50506001600160401b03164090565b6000613fb68360000151846020015185604001518660600151868860a001518960c001518a60e001518b61010001516141c9565b60038360200151604051602001613fce929190615a4a565b60408051601f1981840301815291905280516020909101209392505050565b600080613ff86143e4565b905060005a6140078888615ce3565b6140119190615d7b565b61401b9085615d5c565b9050600061403463ffffffff871664e8d4a51000615d5c565b9050826140418284615ce3565b61404b9190615ce3565b98975050505050505050565b600080614062614440565b905060008113614088576040516321ea67b360e11b81526004810182905260240161097e565b60006140926143e4565b9050600082825a6140a38b8b615ce3565b6140ad9190615d7b565b6140b79088615d5c565b6140c19190615ce3565b6140d390670de0b6b3a7640000615d5c565b6140dd9190615d48565b905060006140f663ffffffff881664e8d4a51000615d5c565b905061410d81676765c793fa10079d601b1b615d7b565b82111561412d5760405163e80fa38160e01b815260040160405180910390fd5b6141378183615ce3565b9998505050505050505050565b6000613856838361450b565b600081815260018301602052604081205461419757508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610ebe565b506000610ebe565b60008260000182815481106141b6576141b6615e82565b9060005260206000200154905092915050565b6141d2896145fe565b61421b5760405162461bcd60e51b815260206004820152601a6024820152797075626c6963206b6579206973206e6f74206f6e20637572766560301b604482015260640161097e565b614224886145fe565b6142685760405162461bcd60e51b815260206004820152601560248201527467616d6d61206973206e6f74206f6e20637572766560581b604482015260640161097e565b614271836145fe565b6142bd5760405162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e206375727665000000604482015260640161097e565b6142c6826145fe565b6143115760405162461bcd60e51b815260206004820152601c60248201527b73486173685769746e657373206973206e6f74206f6e20637572766560201b604482015260640161097e565b61431d878a88876146c1565b6143655760405162461bcd60e51b81526020600482015260196024820152786164647228632a706b2b732a6729213d5f755769746e65737360381b604482015260640161097e565b60006143718a876147d5565b90506000614384898b878b868989614839565b90506000614395838d8d8a8661494c565b9050808a146143d65760405162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b210383937b7b360991b604482015260640161097e565b505050505050505050505050565b60004661a4b18114806143f9575062066eed81145b1561443857606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b815260040160206040518083038186803b158015613d8457600080fd5b600091505090565b600d5460035460408051633fabe5a360e21b81529051600093600160381b900463ffffffff169283151592859283926001600160a01b03169163feaf968c9160048083019260a0929190829003018186803b15801561449e57600080fd5b505afa1580156144b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906144d691906156b7565b5094509092508491505080156144fa57506144f18242615d7b565b8463ffffffff16105b156138a45750601154949350505050565b600081815260018301602052604081205480156145f457600061452f600183615d7b565b855490915060009061454390600190615d7b565b90508181146145a857600086600001828154811061456357614563615e82565b906000526020600020015490508087600001848154811061458657614586615e82565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806145b9576145b9615e6c565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610ebe565b6000915050610ebe565b80516000906401000003d0191161464c5760405162461bcd60e51b8152602060048201526012602482015271696e76616c696420782d6f7264696e61746560701b604482015260640161097e565b60208201516401000003d0191161469a5760405162461bcd60e51b8152602060048201526012602482015271696e76616c696420792d6f7264696e61746560701b604482015260640161097e565b60208201516401000003d0199080096146ba8360005b602002015161498c565b1492915050565b60006001600160a01b0382166147075760405162461bcd60e51b815260206004820152600b60248201526a626164207769746e65737360a81b604482015260640161097e565b60208401516000906001161561471e57601c614721565b601b5b9050600070014551231950b75fc4402da1732fc9bebe1985876000602002015109865170014551231950b75fc4402da1732fc9bebe199182039250600091908909875160408051600080825260209091019182905292935060019161478b91869188918790615910565b6020604051602081039080840390855afa1580156147ad573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b6147dd614e9e565b61480a600184846040516020016147f693929190615857565b6040516020818303038152906040526149b0565b90505b614816816145fe565b610ebe57805160408051602081019290925261483291016147f6565b905061480d565b614841614e9e565b825186516401000003d01990819006910614156148a05760405162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e63740000604482015260640161097e565b6148ab8789886149fe565b6148f05760405162461bcd60e51b8152602060048201526016602482015275119a5c9cdd081b5d5b0818da1958dac819985a5b195960521b604482015260640161097e565b6148fb8486856149fe565b6149415760405162461bcd60e51b815260206004820152601760248201527614d958dbdb99081b5d5b0818da1958dac819985a5b1959604a1b604482015260640161097e565b61404b868484614b19565b60006002868686858760405160200161496a969594939291906157fd565b60408051601f1981840301815291905280516020909101209695505050505050565b6000806401000003d01980848509840990506401000003d019600782089392505050565b6149b8614e9e565b6149c182614bdc565b81526149d66149d18260006146b0565b614c17565b602082018190526002900660011415611dc8576020810180516401000003d019039052919050565b600082614a3b5760405162461bcd60e51b815260206004820152600b60248201526a3d32b9379039b1b0b630b960a91b604482015260640161097e565b83516020850151600090614a5190600290615e2c565b15614a5d57601c614a60565b601b5b9050600070014551231950b75fc4402da1732fc9bebe19838709604080516000808252602090910191829052919250600190614aa3908390869088908790615910565b6020604051602081039080840390855afa158015614ac5573d6000803e3d6000fd5b505050602060405103519050600086604051602001614ae491906157eb565b60408051601f1981840301815291905280516020909101206001600160a01b0392831692169190911498975050505050505050565b614b21614e9e565b835160208086015185519186015160009384938493614b4293909190614c37565b919450925090506401000003d019858209600114614b9e5760405162461bcd60e51b815260206004820152601960248201527834b73b2d1036bab9ba1031329034b73b32b939b29037b3103d60391b604482015260640161097e565b60405180604001604052806401000003d01980614bbd57614bbd615e56565b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d0198110611dc857604080516020808201939093528151808203840181529082019091528051910120614be4565b6000610ebe826002614c306401000003d0196001615ce3565b901c614d17565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a0890506000614c7783838585614dae565b9098509050614c8888828e88614dd2565b9098509050614c9988828c87614dd2565b90985090506000614cac8d878b85614dd2565b9098509050614cbd88828686614dae565b9098509050614cce88828e89614dd2565b9098509050818114614d03576401000003d019818a0998506401000003d01982890997506401000003d0198183099650614d07565b8196505b5050505050509450945094915050565b600080614d22614ebc565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a0820152614d54614eda565b60208160c0846005600019fa925082614da45760405162461bcd60e51b81526020600482015260126024820152716269674d6f64457870206661696c7572652160701b604482015260640161097e565b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b828054828255906000526020600020908101928215614e70579160200282015b82811115614e7057825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614e3b565b50614e7c929150614ef8565b5090565b50805460008255906000526020600020908101906123249190614ef8565b60405180604001604052806002906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5b80821115614e7c5760008155600101614ef9565b8035611dc881615eae565b600082601f830112614f2957600080fd5b813560206001600160401b03821115614f4457614f44615e98565b8160051b614f53828201615c89565b838152828101908684018388018501891015614f6e57600080fd5b600093505b85841015614f9a578035614f8681615eae565b835260019390930192918401918401614f73565b50979650505050505050565b600082601f830112614fb757600080fd5b614fbf615c1c565b808385604086011115614fd157600080fd5b60005b6002811015614ff3578135845260209384019390910190600101614fd4565b509095945050505050565b60008083601f84011261501057600080fd5b5081356001600160401b0381111561502757600080fd5b60208301915083602082850101111561503f57600080fd5b9250929050565b600082601f83011261505757600080fd5b81356001600160401b0381111561507057615070615e98565b615083601f8201601f1916602001615c89565b81815284602083860101111561509857600080fd5b816020850160208301376000918101602001919091529392505050565b600060c082840312156150c757600080fd5b6150cf615c44565b905081356001600160401b0380821682146150e957600080fd5b8183526020840135602084015261510260408501615168565b604084015261511360608501615168565b606084015261512460808501614f0d565b608084015260a084013591508082111561513d57600080fd5b5061514a84828501615046565b60a08301525092915050565b803561ffff81168114611dc857600080fd5b803563ffffffff81168114611dc857600080fd5b80516001600160501b0381168114611dc857600080fd5b80356001600160601b0381168114611dc857600080fd5b6000602082840312156151bc57600080fd5b813561385681615eae565b600080604083850312156151da57600080fd5b82356151e581615eae565b91506151f360208401615193565b90509250929050565b6000806040838503121561520f57600080fd5b823561521a81615eae565b9150602083013561522a81615eae565b809150509250929050565b6000806060838503121561524857600080fd5b823561525381615eae565b91506060830184101561526557600080fd5b50926020919091019150565b6000806000806060858703121561528757600080fd5b843561529281615eae565b93506020850135925060408501356001600160401b038111156152b457600080fd5b6152c087828801614ffe565b95989497509550505050565b6000604082840312156152de57600080fd5b6138568383614fa6565b6000602082840312156152fa57600080fd5b815161385681615ec3565b60006020828403121561531757600080fd5b5035919050565b60006020828403121561533057600080fd5b5051919050565b6000806020838503121561534a57600080fd5b82356001600160401b0381111561536057600080fd5b61536c85828601614ffe565b90969095509350505050565b60006020828403121561538a57600080fd5b604051602081016001600160401b03811182821017156153ac576153ac615e98565b60405282356153ba81615ec3565b81529392505050565b6000808284036101c08112156153d857600080fd5b6101a0808212156153e857600080fd5b6153f0615c66565b91506153fc8686614fa6565b825261540b8660408701614fa6565b60208301526080850135604083015260a0850135606083015260c0850135608083015261543a60e08601614f0d565b60a083015261010061544e87828801614fa6565b60c0840152615461876101408801614fa6565b60e0840152610180860135908301529092508301356001600160401b0381111561548a57600080fd5b615496858286016150b5565b9150509250929050565b6000602082840312156154b257600080fd5b81356001600160401b038111156154c857600080fd5b820160c0818503121561385657600080fd5b6000602082840312156154ec57600080fd5b81356001600160401b038082111561550357600080fd5b9083019060c0828603121561551757600080fd5b61551f615c44565b823560ff8116811461553057600080fd5b81526020838101359082015261554860408401614f0d565b604082015260608301358281111561555f57600080fd5b61556b87828601614f18565b60608301525061557d60808401615193565b608082015261558e60a08401615193565b60a082015295945050505050565b6000602082840312156155ae57600080fd5b61385682615156565b60008060008060008086880360e08112156155d157600080fd5b6155da88615156565b96506155e860208901615168565b95506155f660408901615168565b945061560460608901615168565b9350608088013592506040609f198201121561561f57600080fd5b50615628615c1c565b61563460a08901615168565b815261564260c08901615168565b6020820152809150509295509295509295565b6000806040838503121561566857600080fd5b82359150602083013561522a81615eae565b6000806040838503121561568d57600080fd5b50508035926020909101359150565b6000602082840312156156ae57600080fd5b61385682615168565b600080600080600060a086880312156156cf57600080fd5b6156d88661517c565b94506020860151935060408601519250606086015191506156fb6080870161517c565b90509295509295909350565b600081518084526020808501945080840160005b838110156157405781516001600160a01b03168752958201959082019060010161571b565b509495945050505050565b8060005b60028110156109e357815184526020938401939091019060010161574f565b600081518084526020808501945080840160005b8381101561574057815187529582019590820190600101615782565b6000815180845260005b818110156157c4576020818501810151868301820152016157a8565b818111156157d6576000602083870101525b50601f01601f19169290920160200192915050565b6157f5818361574b565b604001919050565b86815261580d602082018761574b565b61581a606082018661574b565b61582760a082018561574b565b61583460e082018461574b565b60609190911b6001600160601b0319166101208201526101340195945050505050565b838152615867602082018461574b565b606081019190915260800192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b039290921682526001600160601b0316602082015260400190565b60408101610ebe828461574b565b602081526000613856602083018461576e565b918252602082015260400190565b93845260ff9290921660208401526040830152606082015260800190565b602081526000613856602083018461579e565b6020815260ff82511660208201526020820151604082015260018060a01b0360408301511660608201526000606083015160c0608084015261598660e0840182615707565b60808501516001600160601b0390811660a0868101919091529095015190941660c0909301929092525090919050565b61ffff93841681529183166020830152909116604082015260600190565b60006060820161ffff86168352602063ffffffff86168185015260606040850152818551808452608086019150828701935060005b81811015615a2557845183529383019391830191600101615a09565b509098975050505050505050565b9182526001600160a01b0316602082015260400190565b82815260608101613856602083018461574b565b8281526040602082015260006138a4604083018461576e565b86815285602082015261ffff85166040820152600063ffffffff808616606084015280851660808401525060c060a083015261404b60c083018461579e565b878152602081018790526040810186905263ffffffff8581166060830152841660808201526001600160a01b03831660a082015260e060c082018190526000906141379083018461579e565b8781526001600160401b03871660208201526040810186905263ffffffff8581166060830152841660808201526001600160a01b03831660a082015260e060c082018190526000906141379083018461579e565b63ffffffff92831681529116602082015260400190565b6001600160401b0391909116815260200190565b6001600160601b038681168252851660208201526001600160401b03841660408201526001600160a01b038316606082015260a060808201819052600090615bcb90830184615707565b979650505050505050565b6000808335601e19843603018112615bed57600080fd5b8301803591506001600160401b03821115615c0757600080fd5b60200191503681900382131561503f57600080fd5b604080519081016001600160401b0381118282101715615c3e57615c3e615e98565b60405290565b60405160c081016001600160401b0381118282101715615c3e57615c3e615e98565b60405161012081016001600160401b0381118282101715615c3e57615c3e615e98565b604051601f8201601f191681016001600160401b0381118282101715615cb157615cb1615e98565b604052919050565b60008085851115615cc957600080fd5b83861115615cd657600080fd5b5050820193919092039150565b60008219821115615cf657615cf6615e40565b500190565b60006001600160401b03828116848216808303821115615d1d57615d1d615e40565b01949350505050565b60006001600160601b03828116848216808303821115615d1d57615d1d615e40565b600082615d5757615d57615e56565b500490565b6000816000190483118215151615615d7657615d76615e40565b500290565b600082821015615d8d57615d8d615e40565b500390565b60006001600160601b0383811690831681811015615db257615db2615e40565b039392505050565b6001600160e01b03198135818116916004851015615de25780818660040360031b1b83161692505b505092915050565b6000600019821415615dfe57615dfe615e40565b5060010190565b60006001600160401b0382811680821415615e2257615e22615e40565b6001019392505050565b600082615e3b57615e3b615e56565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461232457600080fd5b801515811461232457600080fdfea164736f6c6343000806000a",
}

var VRFCoordinatorV2PlusUpgradedVersionABI = VRFCoordinatorV2PlusUpgradedVersionMetaData.ABI

var VRFCoordinatorV2PlusUpgradedVersionBin = VRFCoordinatorV2PlusUpgradedVersionMetaData.Bin

func DeployVRFCoordinatorV2PlusUpgradedVersion(auth *bind.TransactOpts, backend bind.ContractBackend, blockhashStore common.Address) (common.Address, *types.Transaction, *VRFCoordinatorV2PlusUpgradedVersion, error) {
	parsed, err := VRFCoordinatorV2PlusUpgradedVersionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFCoordinatorV2PlusUpgradedVersionBin), backend, blockhashStore)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinatorV2PlusUpgradedVersion{VRFCoordinatorV2PlusUpgradedVersionCaller: VRFCoordinatorV2PlusUpgradedVersionCaller{contract: contract}, VRFCoordinatorV2PlusUpgradedVersionTransactor: VRFCoordinatorV2PlusUpgradedVersionTransactor{contract: contract}, VRFCoordinatorV2PlusUpgradedVersionFilterer: VRFCoordinatorV2PlusUpgradedVersionFilterer{contract: contract}}, nil
}

type VRFCoordinatorV2PlusUpgradedVersion struct {
	address common.Address
	abi     abi.ABI
	VRFCoordinatorV2PlusUpgradedVersionCaller
	VRFCoordinatorV2PlusUpgradedVersionTransactor
	VRFCoordinatorV2PlusUpgradedVersionFilterer
}

type VRFCoordinatorV2PlusUpgradedVersionCaller struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusUpgradedVersionTransactor struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusUpgradedVersionFilterer struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusUpgradedVersionSession struct {
	Contract     *VRFCoordinatorV2PlusUpgradedVersion
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV2PlusUpgradedVersionCallerSession struct {
	Contract *VRFCoordinatorV2PlusUpgradedVersionCaller
	CallOpts bind.CallOpts
}

type VRFCoordinatorV2PlusUpgradedVersionTransactorSession struct {
	Contract     *VRFCoordinatorV2PlusUpgradedVersionTransactor
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV2PlusUpgradedVersionRaw struct {
	Contract *VRFCoordinatorV2PlusUpgradedVersion
}

type VRFCoordinatorV2PlusUpgradedVersionCallerRaw struct {
	Contract *VRFCoordinatorV2PlusUpgradedVersionCaller
}

type VRFCoordinatorV2PlusUpgradedVersionTransactorRaw struct {
	Contract *VRFCoordinatorV2PlusUpgradedVersionTransactor
}

func NewVRFCoordinatorV2PlusUpgradedVersion(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV2PlusUpgradedVersion, error) {
	abi, err := abi.JSON(strings.NewReader(VRFCoordinatorV2PlusUpgradedVersionABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFCoordinatorV2PlusUpgradedVersion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersion{address: address, abi: abi, VRFCoordinatorV2PlusUpgradedVersionCaller: VRFCoordinatorV2PlusUpgradedVersionCaller{contract: contract}, VRFCoordinatorV2PlusUpgradedVersionTransactor: VRFCoordinatorV2PlusUpgradedVersionTransactor{contract: contract}, VRFCoordinatorV2PlusUpgradedVersionFilterer: VRFCoordinatorV2PlusUpgradedVersionFilterer{contract: contract}}, nil
}

func NewVRFCoordinatorV2PlusUpgradedVersionCaller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV2PlusUpgradedVersionCaller, error) {
	contract, err := bindVRFCoordinatorV2PlusUpgradedVersion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionCaller{contract: contract}, nil
}

func NewVRFCoordinatorV2PlusUpgradedVersionTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV2PlusUpgradedVersionTransactor, error) {
	contract, err := bindVRFCoordinatorV2PlusUpgradedVersion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionTransactor{contract: contract}, nil
}

func NewVRFCoordinatorV2PlusUpgradedVersionFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV2PlusUpgradedVersionFilterer, error) {
	contract, err := bindVRFCoordinatorV2PlusUpgradedVersion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionFilterer{contract: contract}, nil
}

func bindVRFCoordinatorV2PlusUpgradedVersion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorV2PlusUpgradedVersionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.VRFCoordinatorV2PlusUpgradedVersionCaller.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.VRFCoordinatorV2PlusUpgradedVersionTransactor.contract.Transfer(opts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.VRFCoordinatorV2PlusUpgradedVersionTransactor.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.contract.Transfer(opts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "BLOCKHASH_STORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.LINK(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.LINK(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) LINKETHFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "LINK_ETH_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.LINKETHFEED(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.LINKETHFEED(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.MAXCONSUMERS(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.MAXCONSUMERS(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.MAXNUMWORDS(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.MAXNUMWORDS(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) GetActiveSubscriptionIds(opts *bind.CallOpts, startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "getActiveSubscriptionIds", startIndex, maxCount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) GetActiveSubscriptionIds(startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.GetActiveSubscriptionIds(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, startIndex, maxCount)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) GetActiveSubscriptionIds(startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.GetActiveSubscriptionIds(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, startIndex, maxCount)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint16), *new(uint32), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.GetRequestConfig(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.GetRequestConfig(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) GetSubscription(opts *bind.CallOpts, subId *big.Int) (GetSubscription,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(GetSubscription)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EthBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) GetSubscription(subId *big.Int) (GetSubscription,

	error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.GetSubscription(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) GetSubscription(subId *big.Int) (GetSubscription,

	error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.GetSubscription(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.HashOfKey(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, publicKey)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.HashOfKey(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, publicKey)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) MigrationVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "migrationVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) MigrationVersion() (uint8, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.MigrationVersion(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) MigrationVersion() (uint8, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.MigrationVersion(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.Owner(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.Owner(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) PendingRequestExists(opts *bind.CallOpts, subId *big.Int) (bool, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) PendingRequestExists(subId *big.Int) (bool, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.PendingRequestExists(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) PendingRequestExists(subId *big.Int) (bool, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.PendingRequestExists(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) SConfig(opts *bind.CallOpts) (SConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_config")

	outstruct := new(SConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ReentrancyLock = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.StalenessSeconds = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SConfig(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SConfig(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) SCurrentSubNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_currentSubNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SCurrentSubNonce() (uint64, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SCurrentSubNonce(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) SCurrentSubNonce() (uint64, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SCurrentSubNonce(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_fallbackWeiPerUnitLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) SFeeConfig(opts *bind.CallOpts) (SFeeConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_feeConfig")

	outstruct := new(SFeeConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FulfillmentFlatFeeLinkPPM = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeEthPPM = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SFeeConfig() (SFeeConfig,

	error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SFeeConfig(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) SFeeConfig() (SFeeConfig,

	error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SFeeConfig(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_provingKeyHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SProvingKeyHashes(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, arg0)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SProvingKeyHashes(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, arg0)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_provingKeys", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SProvingKeys(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SProvingKeys(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, arg0)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) SProvingKeys(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SProvingKeys(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, arg0)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_requestCommitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SRequestCommitments(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, arg0)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SRequestCommitments(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts, arg0)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) STotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_totalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.STotalBalance(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.STotalBalance(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCaller) STotalEthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2PlusUpgradedVersion.contract.Call(opts, &out, "s_totalEthBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) STotalEthBalance() (*big.Int, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.STotalEthBalance(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionCallerSession) STotalEthBalance() (*big.Int, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.STotalEthBalance(&_VRFCoordinatorV2PlusUpgradedVersion.CallOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "acceptOwnership")
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.AcceptOwnership(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.AcceptOwnership(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.AddConsumer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.AddConsumer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.CancelSubscription(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.CancelSubscription(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "createSubscription")
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.CreateSubscription(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.CreateSubscription(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2PlusUpgradedVersionRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "fulfillRandomWords", proof, rc)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2PlusUpgradedVersionRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.FulfillRandomWords(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2PlusUpgradedVersionRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.FulfillRandomWords(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) FundSubscriptionWithEth(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "fundSubscriptionWithEth", subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) FundSubscriptionWithEth(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.FundSubscriptionWithEth(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) FundSubscriptionWithEth(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.FundSubscriptionWithEth(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) Migrate(opts *bind.TransactOpts, subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "migrate", subId, newCoordinator)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) Migrate(subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.Migrate(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, newCoordinator)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) Migrate(subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.Migrate(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, newCoordinator)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) OnMigration(opts *bind.TransactOpts, encodedData []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "onMigration", encodedData)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) OnMigration(encodedData []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OnMigration(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, encodedData)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) OnMigration(encodedData []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OnMigration(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, encodedData)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OnTokenTransfer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OnTokenTransfer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OracleWithdraw(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OracleWithdraw(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) OracleWithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "oracleWithdrawEth", recipient, amount)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) OracleWithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OracleWithdrawEth(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) OracleWithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OracleWithdrawEth(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) OwnerCancelSubscription(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "ownerCancelSubscription", subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) OwnerCancelSubscription(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OwnerCancelSubscription(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) OwnerCancelSubscription(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.OwnerCancelSubscription(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) RecoverEthFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "recoverEthFunds", to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) RecoverEthFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RecoverEthFunds(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) RecoverEthFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RecoverEthFunds(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "recoverFunds", to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RecoverFunds(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RecoverFunds(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) RegisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "registerMigratableCoordinator", target)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RegisterMigratableCoordinator(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, target)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RegisterMigratableCoordinator(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, target)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "registerProvingKey", oracle, publicProvingKey)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RegisterProvingKey(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RegisterProvingKey(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RemoveConsumer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RemoveConsumer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "requestRandomWords", req)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RequestRandomWords(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, req)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RequestRandomWords(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, req)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusUpgradedVersionFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusUpgradedVersionFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SetConfig(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusUpgradedVersionFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SetConfig(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) SetLINKAndLINKETHFeed(opts *bind.TransactOpts, link common.Address, linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "setLINKAndLINKETHFeed", link, linkEthFeed)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) SetLINKAndLINKETHFeed(link common.Address, linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SetLINKAndLINKETHFeed(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, link, linkEthFeed)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) SetLINKAndLINKETHFeed(link common.Address, linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.SetLINKAndLINKETHFeed(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, link, linkEthFeed)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.TransferOwnership(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, to)
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2PlusUpgradedVersion.Contract.TransferOwnership(&_VRFCoordinatorV2PlusUpgradedVersion.TransactOpts, to)
}

type VRFCoordinatorV2PlusUpgradedVersionConfigSetIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionConfigSet struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
	FallbackWeiPerUnitLink      *big.Int
	FeeConfig                   VRFCoordinatorV2PlusUpgradedVersionFeeConfig
	Raw                         types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionConfigSetIterator, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionConfigSetIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionConfigSet)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseConfigSet(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionConfigSet, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionConfigSet)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegisteredIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegisteredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegisteredIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "CoordinatorRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseCoordinatorRegistered(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionEthFundsRecoveredIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionEthFundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionEthFundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionEthFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterEthFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionEthFundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "EthFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionEthFundsRecoveredIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "EthFundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchEthFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "EthFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "EthFundsRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseEthFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "EthFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionFundsRecoveredIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionFundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionFundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionFundsRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionFundsRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionFundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionFundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionFundsRecoveredIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionFundsRecovered)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionFundsRecovered, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionFundsRecovered)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionMigrationCompletedIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionMigrationCompletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionMigrationCompletedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionMigrationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted struct {
	NewCoordinator common.Address
	SubId          *big.Int
	Raw            types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterMigrationCompleted(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionMigrationCompletedIterator, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "MigrationCompleted")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionMigrationCompletedIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "MigrationCompleted", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchMigrationCompleted(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "MigrationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseMigrationCompleted(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequestedIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequestedIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferredIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferredIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegisteredIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegisteredIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilledIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled struct {
	RequestId  *big.Int
	OutputSeed *big.Int
	SubID      *big.Int
	Payment    *big.Int
	Success    bool
	Raw        types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int, subID []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var subIDRule []interface{}
	for _, subIDItem := range subID {
		subIDRule = append(subIDRule, subIDItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule, subIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilledIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled, requestId []*big.Int, subID []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var subIDRule []interface{}
	for _, subIDItem := range subID {
		subIDRule = append(subIDRule, subIDItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule, subIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequestedIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       *big.Int
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	ExtraArgs                   []byte
	Sender                      common.Address
	Raw                         types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (*VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequestedIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceledIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled struct {
	SubId      *big.Int
	To         common.Address
	AmountLink *big.Int
	AmountEth  *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceledIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAddedIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAddedIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemovedIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemovedIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreatedIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreatedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated struct {
	SubId *big.Int
	Owner common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreatedIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded struct {
	SubId      *big.Int
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEthIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEthIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEthIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth struct {
	SubId         *big.Int
	OldEthBalance *big.Int
	NewEthBalance *big.Int
	Raw           types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterSubscriptionFundedWithEth(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEthIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "SubscriptionFundedWithEth", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEthIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "SubscriptionFundedWithEth", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchSubscriptionFundedWithEth(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "SubscriptionFundedWithEth", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionFundedWithEth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseSubscriptionFundedWithEth(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionFundedWithEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequestedIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequestedIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferredIterator struct {
	Event *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferredIterator{contract: _VRFCoordinatorV2PlusUpgradedVersion.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2PlusUpgradedVersion.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred)
				if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersionFilterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred, error) {
	event := new(VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred)
	if err := _VRFCoordinatorV2PlusUpgradedVersion.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetSubscription struct {
	Balance    *big.Int
	EthBalance *big.Int
	ReqCount   uint64
	Owner      common.Address
	Consumers  []common.Address
}
type SConfig struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	ReentrancyLock              bool
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
}
type SFeeConfig struct {
	FulfillmentFlatFeeLinkPPM uint32
	FulfillmentFlatFeeEthPPM  uint32
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersion) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["ConfigSet"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseConfigSet(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["CoordinatorRegistered"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseCoordinatorRegistered(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["EthFundsRecovered"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseEthFundsRecovered(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["FundsRecovered"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseFundsRecovered(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["MigrationCompleted"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseMigrationCompleted(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseOwnershipTransferRequested(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["OwnershipTransferred"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseOwnershipTransferred(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["ProvingKeyRegistered"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseProvingKeyRegistered(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["RandomWordsFulfilled"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseRandomWordsFulfilled(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["RandomWordsRequested"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseRandomWordsRequested(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["SubscriptionCanceled"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseSubscriptionCanceled(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["SubscriptionConsumerAdded"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseSubscriptionConsumerAdded(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["SubscriptionConsumerRemoved"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseSubscriptionConsumerRemoved(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["SubscriptionCreated"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseSubscriptionCreated(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["SubscriptionFunded"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseSubscriptionFunded(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["SubscriptionFundedWithEth"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseSubscriptionFundedWithEth(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["SubscriptionOwnerTransferRequested"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseSubscriptionOwnerTransferRequested(log)
	case _VRFCoordinatorV2PlusUpgradedVersion.abi.Events["SubscriptionOwnerTransferred"].ID:
		return _VRFCoordinatorV2PlusUpgradedVersion.ParseSubscriptionOwnerTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFCoordinatorV2PlusUpgradedVersionConfigSet) Topic() common.Hash {
	return common.HexToHash("0x777357bb93f63d088f18112d3dba38457aec633eb8f1341e1d418380ad328e78")
}

func (VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered) Topic() common.Hash {
	return common.HexToHash("0xb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625")
}

func (VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x879c9ea2b9d5345b84ccd12610b032602808517cebdb795007f3dcb4df377317")
}

func (VRFCoordinatorV2PlusUpgradedVersionFundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600")
}

func (VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted) Topic() common.Hash {
	return common.HexToHash("0xd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187")
}

func (VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered) Topic() common.Hash {
	return common.HexToHash("0xe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8")
}

func (VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled) Topic() common.Hash {
	return common.HexToHash("0x49580fdfd9497e1ed5c1b1cec0495087ae8e3f1267470ec2fb015db32e3d6aa7")
}

func (VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested) Topic() common.Hash {
	return common.HexToHash("0xeb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e")
}

func (VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled) Topic() common.Hash {
	return common.HexToHash("0x8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4")
}

func (VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded) Topic() common.Hash {
	return common.HexToHash("0x1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e1")
}

func (VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved) Topic() common.Hash {
	return common.HexToHash("0x32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a7")
}

func (VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated) Topic() common.Hash {
	return common.HexToHash("0x1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d")
}

func (VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded) Topic() common.Hash {
	return common.HexToHash("0x1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a")
}

func (VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth) Topic() common.Hash {
	return common.HexToHash("0x3f1ddc3ab1bdb39001ad76ca51a0e6f57ce6627c69f251d1de41622847721cde")
}

func (VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a1")
}

func (VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred) Topic() common.Hash {
	return common.HexToHash("0xd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c9386")
}

func (_VRFCoordinatorV2PlusUpgradedVersion *VRFCoordinatorV2PlusUpgradedVersion) Address() common.Address {
	return _VRFCoordinatorV2PlusUpgradedVersion.address
}

type VRFCoordinatorV2PlusUpgradedVersionInterface interface {
	BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error)

	LINK(opts *bind.CallOpts) (common.Address, error)

	LINKETHFEED(opts *bind.CallOpts) (common.Address, error)

	MAXCONSUMERS(opts *bind.CallOpts) (uint16, error)

	MAXNUMWORDS(opts *bind.CallOpts) (uint32, error)

	MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error)

	GetActiveSubscriptionIds(opts *bind.CallOpts, startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error)

	GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error)

	GetSubscription(opts *bind.CallOpts, subId *big.Int) (GetSubscription,

		error)

	HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error)

	MigrationVersion(opts *bind.CallOpts) (uint8, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingRequestExists(opts *bind.CallOpts, subId *big.Int) (bool, error)

	SConfig(opts *bind.CallOpts) (SConfig,

		error)

	SCurrentSubNonce(opts *bind.CallOpts) (uint64, error)

	SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error)

	SFeeConfig(opts *bind.CallOpts) (SFeeConfig,

		error)

	SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error)

	SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	STotalBalance(opts *bind.CallOpts) (*big.Int, error)

	STotalEthBalance(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error)

	CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error)

	CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error)

	FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2PlusUpgradedVersionRequestCommitment) (*types.Transaction, error)

	FundSubscriptionWithEth(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	Migrate(opts *bind.TransactOpts, subId *big.Int, newCoordinator common.Address) (*types.Transaction, error)

	OnMigration(opts *bind.TransactOpts, encodedData []byte) (*types.Transaction, error)

	OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OracleWithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OwnerCancelSubscription(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	RecoverEthFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RegisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error)

	RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error)

	RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusUpgradedVersionFeeConfig) (*types.Transaction, error)

	SetLINKAndLINKETHFeed(opts *bind.TransactOpts, link common.Address, linkEthFeed common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionConfigSet, error)

	FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegisteredIterator, error)

	WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered) (event.Subscription, error)

	ParseCoordinatorRegistered(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionCoordinatorRegistered, error)

	FilterEthFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionEthFundsRecoveredIterator, error)

	WatchEthFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered) (event.Subscription, error)

	ParseEthFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionEthFundsRecovered, error)

	FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionFundsRecoveredIterator, error)

	WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionFundsRecovered) (event.Subscription, error)

	ParseFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionFundsRecovered, error)

	FilterMigrationCompleted(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusUpgradedVersionMigrationCompletedIterator, error)

	WatchMigrationCompleted(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted) (event.Subscription, error)

	ParseMigrationCompleted(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionMigrationCompleted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionOwnershipTransferred, error)

	FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegisteredIterator, error)

	WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionProvingKeyRegistered, error)

	FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int, subID []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilledIterator, error)

	WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled, requestId []*big.Int, subID []*big.Int) (event.Subscription, error)

	ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled, error)

	FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (*VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequestedIterator, error)

	WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (event.Subscription, error)

	ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested, error)

	FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceledIterator, error)

	WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionCanceled, error)

	FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAddedIterator, error)

	WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerAdded, error)

	FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemovedIterator, error)

	WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionConsumerRemoved, error)

	FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreatedIterator, error)

	WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionCreated, error)

	FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedIterator, error)

	WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionFunded, error)

	FilterSubscriptionFundedWithEth(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEthIterator, error)

	WatchSubscriptionFundedWithEth(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionFundedWithEth(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionFundedWithEth, error)

	FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequestedIterator, error)

	WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferRequested, error)

	FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferredIterator, error)

	WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV2PlusUpgradedVersionSubscriptionOwnerTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
