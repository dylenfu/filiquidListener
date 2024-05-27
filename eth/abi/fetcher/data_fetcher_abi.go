// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fetcher

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
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

// ConversionInteger is an auto generated low-level Go binding around an user-defined struct.
type ConversionInteger struct {
	Value *big.Int
	Neg   bool
}

// DataFetcherFiLLiquidGovernanceFactors is an auto generated low-level Go binding around an user-defined struct.
type DataFetcherFiLLiquidGovernanceFactors struct {
	U1                    *big.Int
	R0                    *big.Int
	R1                    *big.Int
	RM                    *big.Int
	CollateralRate        *big.Int
	MaxFamilySize         *big.Int
	AlertThreshold        *big.Int
	LiquidateThreshold    *big.Int
	MaxLiquidations       *big.Int
	MinLiquidateInterval  *big.Int
	LiquidateDiscountRate *big.Int
	LiquidateFeeRate      *big.Int
	MaxExistingBorrows    *big.Int
	MinBorrowAmount       *big.Int
	MinDepositAmount      *big.Int
	N                     *big.Int
}

// DataFetcherFiLStakeGovernanceFactors is an auto generated low-level Go binding around an user-defined struct.
type DataFetcherFiLStakeGovernanceFactors struct {
	NInterest *big.Int
	NStake    *big.Int
}

// DataFetcherMinerBorrowable is an auto generated low-level Go binding around an user-defined struct.
type DataFetcherMinerBorrowable struct {
	Borrowable bool
	Reason     string
}

// FILLiquidInterfaceBorrowInfo is an auto generated low-level Go binding around an user-defined struct.
type FILLiquidInterfaceBorrowInfo struct {
	Id                      *big.Int
	BorrowAmount            *big.Int
	RemainingOriginalAmount *big.Int
	InterestRate            *big.Int
	DatedTime               *big.Int
	InitialTime             *big.Int
}

// FILLiquidInterfaceBorrowInterestInfo is an auto generated low-level Go binding around an user-defined struct.
type FILLiquidInterfaceBorrowInterestInfo struct {
	Borrow   FILLiquidInterfaceBorrowInfo
	Interest *big.Int
}

// FILLiquidInterfaceFILLiquidInfo is an auto generated low-level Go binding around an user-defined struct.
type FILLiquidInterfaceFILLiquidInfo struct {
	TotalFIL                   *big.Int
	AvailableFIL               *big.Int
	UtilizedLiquidity          *big.Int
	AccumulatedDeposit         *big.Int
	AccumulatedRedeem          *big.Int
	AccumulatedBurntFILTrust   *big.Int
	AccumulatedMintFILTrust    *big.Int
	AccumulatedBorrow          *big.Int
	AccumulatedPayback         *big.Int
	AccumulatedInterest        *big.Int
	AccumulatedRedeemFee       *big.Int
	AccumulatedBorrowFee       *big.Int
	AccumulatedBadDebt         *big.Int
	AccumulatedLiquidateReward *big.Int
	AccumulatedLiquidateFee    *big.Int
	AccumulatedDeposits        *big.Int
	AccumulatedBorrows         *big.Int
	UtilizationRate            *big.Int
	ExchangeRate               *big.Int
	InterestRate               *big.Int
	CollateralizedMiner        *big.Int
	MinerWithBorrows           *big.Int
	RateBase                   *big.Int
}

// FILLiquidInterfaceLiquidateConditionInfo is an auto generated low-level Go binding around an user-defined struct.
type FILLiquidInterfaceLiquidateConditionInfo struct {
	Rate         *big.Int
	Alertable    bool
	Liquidatable bool
}

// FILLiquidInterfaceMinerBorrowInfo is an auto generated low-level Go binding around an user-defined struct.
type FILLiquidInterfaceMinerBorrowInfo struct {
	MinerId          uint64
	DebtOutStanding  *big.Int
	Balance          *big.Int
	BorrowSum        *big.Int
	AvailableBalance ConversionInteger
	Borrows          []FILLiquidInterfaceBorrowInterestInfo
}

// FILLiquidInterfaceUserInfo is an auto generated low-level Go binding around an user-defined struct.
type FILLiquidInterfaceUserInfo struct {
	User                   common.Address
	DebtOutStanding        *big.Int
	AvailableCredit        *big.Int
	BalanceSum             *big.Int
	BorrowSum              *big.Int
	LiquidateConditionInfo FILLiquidInterfaceLiquidateConditionInfo
	MinerBorrowInfo        []FILLiquidInterfaceMinerBorrowInfo
}

// FILStakeFILStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type FILStakeFILStakeInfo struct {
	AccumulatedInterest      *big.Int
	AccumulatedStake         *big.Int
	AccumulatedStakeDuration *big.Int
	AccumulatedInterestMint  *big.Int
	AccumulatedStakeMint     *big.Int
	AccumulatedWithdrawn     *big.Int
	NextStakeID              *big.Int
	ReleasedFigStake         *big.Int
}

// GovernanceGovernanceInfo is an auto generated low-level Go binding around an user-defined struct.
type GovernanceGovernanceInfo struct {
	Bonders               *big.Int
	TotalBondedAmount     *big.Int
	FirstActiveProposalId *big.Int
}

// DataMetaData contains all meta data concerning the Data contract.
var DataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractFILLiquid\",\"name\":\"filLiquidAddr\",\"type\":\"address\"},{\"internalType\":\"contractFILTrust\",\"name\":\"filTrustAddr\",\"type\":\"address\"},{\"internalType\":\"contractFILStake\",\"name\":\"filStakeAddr\",\"type\":\"address\"},{\"internalType\":\"contractFILGovernance\",\"name\":\"filGovernanceAddr\",\"type\":\"address\"},{\"internalType\":\"contractGovernance\",\"name\":\"governanceAddr\",\"type\":\"address\"},{\"internalType\":\"contractFilecoinAPI\",\"name\":\"filecoinAPIAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"fetchData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fitTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"figTotalSupply\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalFIL\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableFIL\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBurntFILTrust\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedMintFILTrust\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedPayback\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRedeemFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBorrowFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBadDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedLiquidateReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedLiquidateFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralizedMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minerWithBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateBase\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.FILLiquidInfo\",\"name\":\"filLiquidInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accumulatedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedStakeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedInterestMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedStakeMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedWithdrawn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextStakeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releasedFigStake\",\"type\":\"uint256\"}],\"internalType\":\"structFILStake.FILStakeInfo\",\"name\":\"filStakeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bonders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBondedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstActiveProposalId\",\"type\":\"uint256\"}],\"internalType\":\"structGovernance.GovernanceInfo\",\"name\":\"governanceInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"u_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFamilySize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alertThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidateInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateDiscountRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExistingBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"internalType\":\"structDataFetcher.FiLLiquidGovernanceFactors\",\"name\":\"fiLLiquidGovernanceFactors\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"n_interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n_stake\",\"type\":\"uint256\"}],\"internalType\":\"structDataFetcher.FiLStakeGovernanceFactors\",\"name\":\"fiLStakeGovernanceFactors\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchGovernanceFactors\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"u_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFamilySize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alertThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidateInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateDiscountRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExistingBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"internalType\":\"structDataFetcher.FiLLiquidGovernanceFactors\",\"name\":\"fiLLiquidGovernanceFactors\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"n_interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n_stake\",\"type\":\"uint256\"}],\"internalType\":\"structDataFetcher.FiLStakeGovernanceFactors\",\"name\":\"fiLStakeGovernanceFactors\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"fetchPersonalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filTrustBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filBalance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alertable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"liquidatable\",\"type\":\"bool\"}],\"internalType\":\"structFILLiquidInterface.LiquidateConditionInfo\",\"name\":\"liquidateConditionInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structConversion.Integer\",\"name\":\"availableBalance\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingOriginalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"datedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInfo\",\"name\":\"borrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInterestInfo[]\",\"name\":\"borrows\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.MinerBorrowInfo[]\",\"name\":\"minerBorrowInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.UserInfo\",\"name\":\"userInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"fetchStakerData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filTrustBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filTrustFixed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filTrustVariable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filGovernanceBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[6]\",\"name\":\"\",\"type\":\"address[6]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"getBatchedUserBorrows\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alertable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"liquidatable\",\"type\":\"bool\"}],\"internalType\":\"structFILLiquidInterface.LiquidateConditionInfo\",\"name\":\"liquidateConditionInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structConversion.Integer\",\"name\":\"availableBalance\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingOriginalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"datedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInfo\",\"name\":\"borrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInterestInfo[]\",\"name\":\"borrows\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.MinerBorrowInfo[]\",\"name\":\"minerBorrowInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.UserInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBorrowExpecting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sixMonthInterest\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBorrowable\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"borrowable\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structDataFetcher.MinerBorrowable[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountFIL\",\"type\":\"uint256\"}],\"name\":\"getDepositExpecting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedAmountFILTrust\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"}],\"name\":\"getPendingBeneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountFILTrust\",\"type\":\"uint256\"}],\"name\":\"getRedeemExpecting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedAmountFIL\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPendingInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPendingInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFIL\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingAndPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedPayback\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestExp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserBorrowsAndBorrowable\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alertable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"liquidatable\",\"type\":\"bool\"}],\"internalType\":\"structFILLiquidInterface.LiquidateConditionInfo\",\"name\":\"liquidateConditionInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structConversion.Integer\",\"name\":\"availableBalance\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingOriginalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"datedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInfo\",\"name\":\"borrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInterestInfo[]\",\"name\":\"borrows\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.MinerBorrowInfo[]\",\"name\":\"minerBorrowInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.UserInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"borrowable\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structDataFetcher.MinerBorrowable[]\",\"name\":\"borrowables\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"}],\"name\":\"getUserBorrowsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alertable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"liquidatable\",\"type\":\"bool\"}],\"internalType\":\"structFILLiquidInterface.LiquidateConditionInfo\",\"name\":\"liquidateConditionInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structConversion.Integer\",\"name\":\"availableBalance\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingOriginalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"datedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInfo\",\"name\":\"borrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInterestInfo[]\",\"name\":\"borrows\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.MinerBorrowInfo[]\",\"name\":\"minerBorrowInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.UserInfo\",\"name\":\"infos\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"}],\"name\":\"maxBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"afterBlocks\",\"type\":\"uint256\"}],\"name\":\"maxBorrowAllowedInAdvance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620031c9380380620031c98339810160408190526200003391620000bd565b5f80546001600160a01b03199081166001600160a01b03988916179091556001805482169688169690961790955560028054861694871694909417909355600380548516928616929092179091556004805484169185169190911790556005805490921692169190911790556200014d565b6001600160a01b0381168114620000ba575f80fd5b50565b5f805f805f8060c08789031215620000d3575f80fd5b8651620000e081620000a5565b6020880151909650620000f381620000a5565b60408801519095506200010681620000a5565b60608801519094506200011981620000a5565b60808801519093506200012c81620000a5565b60a08801519092506200013f81620000a5565b809150509295509295509295565b61306e806200015b5f395ff3fe608060405234801561000f575f80fd5b50600436106100fb575f3560e01c8063a39fac1211610093578063c459110011610063578063c4591100146102ab578063cc798ad5146102cb578063ea5565a7146102de578063ec12bf50146102f1575f80fd5b8063a39fac1214610242578063a405a88f14610257578063a71d4d7614610277578063b8ce16e41461028a575f80fd5b8063855c55f3116100ce578063855c55f3146101a857806391546352146101eb5780639685e7161461020c5780639fd676d21461022c575f80fd5b80632b3c4cac146100ff578063314befa3146101375780636536c50f1461016957806367845e161461018b575b5f80fd5b61011261010d366004611e79565b610319565b6040805194855260208501939093529183015260608201526080015b60405180910390f35b61014a610145366004611eaf565b610474565b604080516001600160a01b03909316835290151560208301520161012e565b61017c610177366004611e79565b6104f3565b60405161012e93929190612061565b6101936105e6565b60405161012e99989796959493929190612122565b6101b061096f565b604080519889526020890197909752958701949094526060860192909252608085015260a084015260c083015260e08201526101000161012e565b6101fe6101f9366004611eaf565b610cf4565b60405190815260200161012e565b61021f61021a366004611e79565b610f05565b60405161012e91906123a9565b6102346110d2565b60405161012e9291906123bb565b61024a611317565b60405161012e91906123e2565b61026a610265366004612412565b611371565b60405161012e9190612480565b6101fe6102853660046124e2565b61149d565b61029d610298366004611e79565b611647565b60405161012e92919061250c565b6102be6102b9366004611eaf565b6116cf565b60405161012e9190612530565b6101fe6102d9366004612542565b6117b9565b6101fe6102ec366004612542565b611825565b6103046102ff366004612542565b611856565b6040805192835260208301919091520161012e565b6001546040516370a0823160e01b81526001600160a01b0383811660048301525f9283928392839216906370a0823190602401602060405180830381865afa158015610367573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061038b9190612559565b600254604051638c04942b60e01b81526001600160a01b038881166004830152929650911690638c04942b906024016040805180830381865afa1580156103d4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103f89190612570565b6003546040516370a0823160e01b81526001600160a01b03898116600483015293965091945091909116906370a0823190602401602060405180830381865afa158015610447573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046b9190612559565b90509193509193565b60055460405163513e132d60e11b81526001600160401b03831660048201525f9182916001600160a01b039091169063a27c265a906024016040805180830381865afa1580156104c6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104ea91906125b1565b91509150915091565b5f806104fd611d73565b6001546040516370a0823160e01b81526001600160a01b038681166004830152909116906370a0823190602401602060405180830381865afa158015610545573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105699190612559565b5f54604051631dd64f3160e21b81526001600160a01b0387811660048301819052939650923194509116906377593cc4906024015f60405180830381865afa1580156105b7573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105de9190810190612937565b929491935050565b5f805f80610684604051806102e001604052805f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b6106c46040518061010001604052805f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b6106e560405180606001604052805f81526020015f81526020015f81525090565b6106ed611dd7565b604080518082019091525f808252602082015243985042975060015f9054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610756573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077a9190612559565b965060035f9054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107cc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107f09190612559565b95505f8054906101000a90046001600160a01b03166001600160a01b0316634e69d5606040518163ffffffff1660e01b81526004016102e060405180830381865afa158015610841573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061086591906129eb565b945060025f9054906101000a90046001600160a01b03166001600160a01b0316634e69d5606040518163ffffffff1660e01b815260040161010060405180830381865afa1580156108b8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108dc9190612afc565b935060045f9054906101000a90046001600160a01b03166001600160a01b0316634e69d5606040518163ffffffff1660e01b8152600401606060405180830381865afa15801561092e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109529190612b89565b925061095c6110d2565b8092508193505050909192939495969798565b5f805f805f805f804397504296505f80805f9054906101000a90046001600160a01b03166001600160a01b031663764f3f326040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109ce573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109f29190612559565b90505f805f9054906101000a90046001600160a01b03166001600160a01b0316634e69d5606040518163ffffffff1660e01b81526004016102e060405180830381865afa158015610a45573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a6991906129eb565b9050818314610cd2575f80546040516354399efd60e01b815260048101869052602481018590526001600160a01b03909116906354399efd906044015f60405180830381865afa158015610abf573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610ae69190810190612bc3565b90505f5b8151811015610ccf575f828281518110610b0657610b06612cb1565b60200260200101519050806020015160200151610b235750610cc7565b5f80548251604051630244e01b60e41b81526001600160401b0390911660048201526001600160a01b039091169063244e01b0906024015f60405180830381865afa158015610b74573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610b9b9190810190612cc5565b905080606001518160200151610bb19190612d12565b610bbb908e612d25565b9c505f5b8160a0015151811015610cc3575f8260a001518281518110610be357610be3612cb1565b60200260200101515f015190508060a0015142610c009190612d12565b8160400151610c0f9190612d38565b610c19908d612d25565b60208201515f54606084015160405163c9abb10f60e01b8152600481018490526301e1338060248201526044810191909152929e5090916001600160a01b039091169063c9abb10f90606401602060405180830381865afa158015610c80573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ca49190612559565b610cae9190612d12565b610cb8908b612d25565b995050600101610bbf565b5050505b600101610aea565b50505b805f015197508060400151965080610100015194505050509091929394959697565b5f8054604051631e310edd60e31b81526001600160401b038416600482015282916001600160a01b03169063f18876e8906024015f60405180830381865afa158015610d42573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d699190810190612d4f565b50905080610d7957505f92915050565b5f8054906101000a90046001600160a01b03166001600160a01b031663df4b0aef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dc7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610deb9190612559565b5f80546040516348aa31a960e11b81526001600160401b038716600482015292945090916001600160a01b0390911690639154635290602401602060405180830381865afa158015610e3f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e639190612559565b905080831115610e71578092505b5f805f9054906101000a90046001600160a01b03166001600160a01b03166354397e5c6040518163ffffffff1660e01b815260040161014060405180830381865afa158015610ec2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ee69190612df0565b505050509550505050505080841015610efd575f93505b505050919050565b5f8054604051633c20dad160e21b81526001600160a01b03848116600483015260609392169063f0836b44906024015f60405180830381865afa158015610f4e573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610f759190810190612e6e565b905080516001600160401b03811115610f9057610f906125e4565b604051908082528060200260200182016040528015610fd557816020015b604080518082019091525f815260606020820152815260200190600190039081610fae5790505b5091505f5b81518110156110cb575f5482516001600160a01b039091169063f18876e89084908490811061100b5761100b612cb1565b60200260200101516040518263ffffffff1660e01b815260040161103e91906001600160401b0391909116815260200190565b5f60405180830381865afa158015611058573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261107f9190810190612d4f565b84838151811061109157611091612cb1565b60200260200101515f018584815181106110ad576110ad612cb1565b60209081029190910181015101919091529015159052600101610fda565b5050919050565b6110da611dd7565b604080518082019091525f80825260208201525f8054906101000a90046001600160a01b03166001600160a01b031663b77d061c6040518163ffffffff1660e01b815260040160a060405180830381865afa15801561113b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061115f9190612f02565b6101e0870152606086015260408086019190915260208501919091529083525f54815163150e5f9760e21b815291516001600160a01b03909116916354397e5c916004808301926101409291908290030181865afa1580156111c3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111e79190612df0565b505060a08a01526101808901526101a08801526101c0870152608086015250505f546040805163bd56162b60e01b815290516001600160a01b03909216925063bd56162b9160048083019260c09291908290030181865afa15801561124e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112729190612f3e565b61016088015261014087015260e086015260c085015261012080850191909152610100840191909152600254604080516334dfbea760e21b815290515f936001600160a01b039093169263d37efa9c92600480820193918290030181865afa1580156112e0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113049190612f84565b8051835260209081015190830152509091565b61131f611e44565b506040805160c0810182525f546001600160a01b03908116825260015481166020830152600254811692820192909252600354821660608201526004548216608082015260055490911660a082015290565b6060816001600160401b0381111561138b5761138b6125e4565b6040519080825280602002602001820160405280156113c457816020015b6113b1611d73565b8152602001906001900390816113a95790505b5090505f5b8151811015611496575f546001600160a01b03166377593cc48585848181106113f4576113f4612cb1565b90506020020160208101906114099190611e79565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024015f60405180830381865afa15801561144a573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526114719190810190612937565b82828151811061148357611483612cb1565b60209081029190910101526001016113c9565b5092915050565b5f8054604051631e310edd60e31b81526001600160401b038516600482015282916001600160a01b03169063f18876e8906024015f60405180830381865afa1580156114eb573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526115129190810190612d4f565b50905080611523575f915050611641565b5f805f9054906101000a90046001600160a01b03166001600160a01b031663df4b0aef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611573573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115979190612559565b90506115a3858561194a565b9250808311156115b1578092505b5f805f9054906101000a90046001600160a01b03166001600160a01b03166354397e5c6040518163ffffffff1660e01b815260040161014060405180830381865afa158015611602573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116269190612df0565b50505050955050505050508084101561163d575f93505b5050505b92915050565b61164f611d73565b5f54604051631dd64f3160e21b81526001600160a01b03848116600483015260609216906377593cc4906024015f60405180830381865afa158015611696573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526116bd9190810190612937565b91506116c883610f05565b9050915091565b6116d7611d73565b5f5460405163090a55c560e21b81526001600160401b03841660048201526001600160a01b03909116906377593cc4908290632429571490602401602060405180830381865afa15801561172d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117519190612ffe565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024015f60405180830381865afa158015611792573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526116419190810190612937565b5f80546040516337b6fa7560e01b8152600481018490526001600160a01b03909116906337b6fa75906024015b602060405180830381865afa158015611801573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116419190612559565b5f805460405163d6167ce560e01b8152600481018490526001600160a01b039091169063d6167ce5906024016117e6565b5f8054604051630ba0a8b160e31b81526004810184905282916001600160a01b031690635d05458890602401602060405180830381865afa15801561189d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118c19190612559565b5f5460405163c9abb10f60e01b81526004810186905262ed4e0060248201526044810183905291935084916001600160a01b039091169063c9abb10f90606401602060405180830381865afa15801561191c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119409190612559565b6116c89190612d12565b5f805460405163090a55c560e21b81526001600160401b038516600482015282916001600160a01b03169063f0836b44908290632429571490602401602060405180830381865afa1580156119a1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119c59190612ffe565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024015f60405180830381865afa158015611a06573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611a2d9190810190612e6e565b90505f80805b8351811015611c10575f805485516001600160a01b039091169063244e01b090879085908110611a6557611a65612cb1565b60200260200101516040518263ffffffff1660e01b8152600401611a9891906001600160401b0391909116815260200190565b5f60405180830381865afa158015611ab2573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611ad99190810190612cc5565b60a001519050611b01858381518110611af457611af4612cb1565b6020026020010151611c25565b611b15906001600160a01b03163185612d25565b93505f5b8151811015611c06575f828281518110611b3557611b35612cb1565b602090810291909101810151515f549181015160808201519193506001600160a01b039092169163c9abb10f91611b6d8d601e612d38565b611b779042612d25565b611b819190612d12565b60608501516040516001600160e01b031960e086901b168152600481019390935260248301919091526044820152606401602060405180830381865afa158015611bcd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bf19190612559565b611bfb9086612d25565b945050600101611b19565b5050600101611a33565b50611c1b8282611c98565b9695505050505050565b60055460405163b6ecd87f60e01b81526001600160401b03831660048201525f916001600160a01b03169063b6ecd87f90602401602060405180830381865afa158015611c74573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116419190612ffe565b5f805f805f9054906101000a90046001600160a01b03166001600160a01b03166354397e5c6040518163ffffffff1660e01b815260040161014060405180830381865afa158015611ceb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d0f9190612df0565b5050505050509350505091505f8186611d289190612d38565b90505f611d358487612d38565b9050808211611d4a575f945050505050611641565b611d548385612d12565b611d5e8284612d12565b611d689190613019565b945050505050611641565b6040518060e001604052805f6001600160a01b031681526020015f81526020015f81526020015f81526020015f8152602001611dca60405180606001604052805f81526020015f151581526020015f151581525090565b8152602001606081525090565b6040518061020001604052805f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b6040518060c001604052806006906020820280368337509192915050565b6001600160a01b0381168114611e76575f80fd5b50565b5f60208284031215611e89575f80fd5b8135611e9481611e62565b9392505050565b6001600160401b0381168114611e76575f80fd5b5f60208284031215611ebf575f80fd5b8135611e9481611e9b565b5f82825180855260208086019550808260051b8401018186015f5b84811015611fd657858303601f19018952815180516001600160401b0316845284810151858501526040808201518186015260608083015181870152608080840151805182890152880151151560a0808901919091529384015160e060c089018190528151818a01819052918a019590945f949093909290916101008b01915b80871015611fb85788518051805185528e8101518f860152878101518886015286810151878601528581015186860152830151838501528d015160c0840152978c01976001969096019591870191611f65565b50509e8a019e98505050948701945050506001919091019050611ee5565b5090979650505050505050565b5f61012060018060a01b0383511684526020830151602085015260408301516040850152606083015160608501526080830151608085015260a0830151805160a08601526020810151151560c08601526040810151151560e08601525060c08301518161010086015261205882860182611eca565b95945050505050565b838152826020820152606060408201525f6120586060830184611fe3565b805182526020808201519083015260408082015190830152606080820151908301526080808201519083015260a0808201519083015260c0808201519083015260e08082015190830152610100808201519083015261012080820151908301526101408082015190830152610160808201519083015261018080820151908301526101a080820151908301526101c080820151908301526101e090810151910152565b89815260208082018a905260408083018a905260608084018a905288516080808601919091529289015160a0808601919091529189015160c0808601919091529089015160e08086019190915292890151610100808601919091529189015161012080860191909152908901516101408086019190915292890151610160808601919091529189015161018080860191909152908901516101a080860191909152928901516101c080860191909152918901516101e08086019190915290890151610200808601919091529289015161022080860191909152918901516102408086019190915290890151610260808601919091529289015161028080860191909152918901516102a080860191909152908901516102c080860191909152928901516102e08501529088015161030084015287015161032083015286015161034082015261070081016122c3610360830187805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b8451610460830152602085015161048083015260408501516104a08301526122ef6104c083018561207f565b82516106c083015260208301516106e08301529a9950505050505050505050565b5f5b8381101561232a578181015183820152602001612312565b50505f910152565b5f82825180855260208086019550808260051b8401018186015f5b84811015611fd657601f1986840381018a5282518051151585528501516040868601819052815190860181905260609061238c81838901858b01612310565b9b87019b601f01909216949094010192509083019060010161234d565b602081525f611e946020830184612332565b61024081016123ca828561207f565b82516102008301526020830151610220830152611e94565b60c0810181835f5b600681101561163d5781516001600160a01b03168352602092830192909101906001016123ea565b5f8060208385031215612423575f80fd5b82356001600160401b0380821115612439575f80fd5b818501915085601f83011261244c575f80fd5b81358181111561245a575f80fd5b8660208260051b850101111561246e575f80fd5b60209290920196919550909350505050565b5f60208083016020845280855180835260408601915060408160051b8701019250602087015f5b828110156124d557603f198886030184526124c3858351611fe3565b945092850192908501906001016124a7565b5092979650505050505050565b5f80604083850312156124f3575f80fd5b82356124fe81611e9b565b946020939093013593505050565b604081525f61251e6040830185611fe3565b82810360208401526120588185612332565b602081525f611e946020830184611fe3565b5f60208284031215612552575f80fd5b5035919050565b5f60208284031215612569575f80fd5b5051919050565b5f8060408385031215612581575f80fd5b505080516020909101519092909150565b805161259d81611e62565b919050565b8051801515811461259d575f80fd5b5f80604083850312156125c2575f80fd5b82516125cd81611e62565b91506125db602084016125a2565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561261a5761261a6125e4565b60405290565b604080519081016001600160401b038111828210171561261a5761261a6125e4565b60405160c081016001600160401b038111828210171561261a5761261a6125e4565b60405160e081016001600160401b038111828210171561261a5761261a6125e4565b6040516102e081016001600160401b038111828210171561261a5761261a6125e4565b604051601f8201601f191681016001600160401b03811182821017156126d1576126d16125e4565b604052919050565b5f606082840312156126e9575f80fd5b6126f16125f8565b905081518152612703602083016125a2565b6020820152612714604083016125a2565b604082015292915050565b5f6001600160401b03821115612737576127376125e4565b5060051b60200190565b5f82601f830112612750575f80fd5b815160206127656127608361271f565b6126a9565b82815260e09283028501820192828201919087851115612783575f80fd5b8387015b85811015611fd6578089038281121561279e575f80fd5b6127a6612620565b60c0808312156127b4575f80fd5b6127bc612642565b84518152888501518982015260408086015190820152606080860151908201526080808601519082015260a08086015190820152825283015181880152855250928401928101612787565b5f81830360e0811215612818575f80fd5b612820612642565b9150825161282d81611e9b565b8252602083810151908301526040808401518184015260608085015190840152607f198201121561285c575f80fd5b50612865612620565b6080830151815261287860a084016125a2565b6020820152608082015260c08201516001600160401b0381111561289a575f80fd5b6128a684828501612741565b60a08301525092915050565b5f82601f8301126128c1575f80fd5b815160206128d16127608361271f565b82815260059290921b840181019181810190868411156128ef575f80fd5b8286015b8481101561292c5780516001600160401b03811115612910575f80fd5b61291e8986838b0101612807565b8452509183019183016128f3565b509695505050505050565b5f60208284031215612947575f80fd5b81516001600160401b038082111561295d575f80fd5b908301906101208286031215612971575f80fd5b612979612664565b61298283612592565b8152602083015160208201526040830151604082015260608301516060820152608083015160808201526129b98660a085016126d9565b60a0820152610100830151828111156129d0575f80fd5b6129dc878286016128b2565b60c08301525095945050505050565b5f6102e082840312156129fc575f80fd5b612a04612686565b825181526020808401519082015260408084015190820152606080840151908201526080808401519082015260a0808401519082015260c0808401519082015260e08084015190820152610100808401519082015261012080840151908201526101408084015190820152610160808401519082015261018080840151908201526101a080840151908201526101c080840151908201526101e08084015190820152610200808401519082015261022080840151908201526102408084015190820152610260808401519082015261028080840151908201526102a080840151908201526102c0928301519281019290925250919050565b5f610100808385031215612b0e575f80fd5b604051908101906001600160401b0382118183101715612b3057612b306125e4565b81604052835181526020840151602082015260408401516040820152606084015160608201526080840151608082015260a084015160a082015260c084015160c082015260e084015160e0820152809250505092915050565b5f60608284031215612b99575f80fd5b612ba16125f8565b8251815260208301516020820152604083015160408201528091505092915050565b5f6020808385031215612bd4575f80fd5b82516001600160401b03811115612be9575f80fd5b8301601f81018513612bf9575f80fd5b8051612c076127608261271f565b81815260609182028301840191848201919088841115612c25575f80fd5b938501935b83851015612ca55784890381811215612c41575f80fd5b612c49612620565b8651612c5481611e9b565b81526040601f198301811315612c68575f80fd5b612c70612620565b9250612c7d8989016125a2565b8352612c8a8189016125a2565b838a0152508088019190915283529384019391850191612c2a565b50979650505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215612cd5575f80fd5b81516001600160401b03811115612cea575f80fd5b612cf684828501612807565b949350505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561164157611641612cfe565b8082018082111561164157611641612cfe565b808202811582820484141761164157611641612cfe565b5f8060408385031215612d60575f80fd5b612d69836125a2565b915060208301516001600160401b0380821115612d84575f80fd5b818501915085601f830112612d97575f80fd5b815181811115612da957612da96125e4565b612dbc601f8201601f19166020016126a9565b9150808252866020828501011115612dd2575f80fd5b612de3816020840160208601612310565b5080925050509250929050565b5f805f805f805f805f806101408b8d031215612e0a575f80fd5b8a51995060208b0151985060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935060e08b015192506101008b015191506101208b01518060070b8114612e5b575f80fd5b809150509295989b9194979a5092959850565b5f6020808385031215612e7f575f80fd5b82516001600160401b03811115612e94575f80fd5b8301601f81018513612ea4575f80fd5b8051612eb26127608261271f565b81815260059190911b82018301908381019087831115612ed0575f80fd5b928401925b82841015612ef7578351612ee881611e9b565b82529284019290840190612ed5565b979650505050505050565b5f805f805f60a08688031215612f16575f80fd5b5050835160208501516040860151606087015160809097015192989197509594509092509050565b5f805f805f8060c08789031215612f53575f80fd5b865195506020870151945060408701519350606087015192506080870151915060a087015190509295509295509295565b5f610120808385031215612f96575f80fd5b83601f840112612fa4575f80fd5b6040518181018181106001600160401b0382111715612fc557612fc56125e4565b604052908301908085831115612fd9575f80fd5b845b83811015612ff3578051825260209182019101612fdb565b509095945050505050565b5f6020828403121561300e575f80fd5b8151611e9481611e62565b5f8261303357634e487b7160e01b5f52601260045260245ffd5b50049056fea26469706673582212203683fbc3a998fa887934524a71bcff13ffddbe3f8aa9d1ca58bfcecd46a2b30e64736f6c63430008160033",
}

// DataABI is the input ABI used to generate the binding from.
// Deprecated: Use DataMetaData.ABI instead.
var DataABI = DataMetaData.ABI

// DataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DataMetaData.Bin instead.
var DataBin = DataMetaData.Bin

// DeployData deploys a new Ethereum contract, binding an instance of Data to it.
func DeployData(auth *bind.TransactOpts, backend bind.ContractBackend, filLiquidAddr common.Address, filTrustAddr common.Address, filStakeAddr common.Address, filGovernanceAddr common.Address, governanceAddr common.Address, filecoinAPIAddr common.Address) (common.Address, *types.Transaction, *Data, error) {
	parsed, err := DataMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DataBin), backend, filLiquidAddr, filTrustAddr, filStakeAddr, filGovernanceAddr, governanceAddr, filecoinAPIAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}, DataFilterer: DataFilterer{contract: contract}}, nil
}

// Data is an auto generated Go binding around an Ethereum contract.
type Data struct {
	DataCaller     // Read-only binding to the contract
	DataTransactor // Write-only binding to the contract
	DataFilterer   // Log filterer for contract events
}

// DataCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataSession struct {
	Contract     *Data             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataCallerSession struct {
	Contract *DataCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTransactorSession struct {
	Contract     *DataTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataRaw struct {
	Contract *Data // Generic contract binding to access the raw methods on
}

// DataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataCallerRaw struct {
	Contract *DataCaller // Generic read-only contract binding to access the raw methods on
}

// DataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTransactorRaw struct {
	Contract *DataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewData creates a new instance of Data, bound to a specific deployed contract.
func NewData(address common.Address, backend bind.ContractBackend) (*Data, error) {
	contract, err := bindData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}, DataFilterer: DataFilterer{contract: contract}}, nil
}

// NewDataCaller creates a new read-only instance of Data, bound to a specific deployed contract.
func NewDataCaller(address common.Address, caller bind.ContractCaller) (*DataCaller, error) {
	contract, err := bindData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataCaller{contract: contract}, nil
}

// NewDataTransactor creates a new write-only instance of Data, bound to a specific deployed contract.
func NewDataTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTransactor, error) {
	contract, err := bindData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTransactor{contract: contract}, nil
}

// NewDataFilterer creates a new log filterer instance of Data, bound to a specific deployed contract.
func NewDataFilterer(address common.Address, filterer bind.ContractFilterer) (*DataFilterer, error) {
	contract, err := bindData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataFilterer{contract: contract}, nil
}

// bindData binds a generic wrapper to an already deployed contract.
func bindData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Data.Contract.DataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Data.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.contract.Transact(opts, method, params...)
}

// FetchData is a free data retrieval call binding the contract method 0x67845e16.
//
// Solidity: function fetchData() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 fitTotalSupply, uint256 figTotalSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filLiquidInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filStakeInfo, (uint256,uint256,uint256) governanceInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fiLStakeGovernanceFactors)
func (_Data *DataCaller) FetchData(opts *bind.CallOpts) (struct {
	BlockHeight                *big.Int
	BlockTimeStamp             *big.Int
	FitTotalSupply             *big.Int
	FigTotalSupply             *big.Int
	FilLiquidInfo              FILLiquidInterfaceFILLiquidInfo
	FilStakeInfo               FILStakeFILStakeInfo
	GovernanceInfo             GovernanceGovernanceInfo
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FiLStakeGovernanceFactors  DataFetcherFiLStakeGovernanceFactors
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "fetchData")

	outstruct := new(struct {
		BlockHeight                *big.Int
		BlockTimeStamp             *big.Int
		FitTotalSupply             *big.Int
		FigTotalSupply             *big.Int
		FilLiquidInfo              FILLiquidInterfaceFILLiquidInfo
		FilStakeInfo               FILStakeFILStakeInfo
		GovernanceInfo             GovernanceGovernanceInfo
		FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
		FiLStakeGovernanceFactors  DataFetcherFiLStakeGovernanceFactors
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockTimeStamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FitTotalSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FigTotalSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FilLiquidInfo = *abi.ConvertType(out[4], new(FILLiquidInterfaceFILLiquidInfo)).(*FILLiquidInterfaceFILLiquidInfo)
	outstruct.FilStakeInfo = *abi.ConvertType(out[5], new(FILStakeFILStakeInfo)).(*FILStakeFILStakeInfo)
	outstruct.GovernanceInfo = *abi.ConvertType(out[6], new(GovernanceGovernanceInfo)).(*GovernanceGovernanceInfo)
	outstruct.FiLLiquidGovernanceFactors = *abi.ConvertType(out[7], new(DataFetcherFiLLiquidGovernanceFactors)).(*DataFetcherFiLLiquidGovernanceFactors)
	outstruct.FiLStakeGovernanceFactors = *abi.ConvertType(out[8], new(DataFetcherFiLStakeGovernanceFactors)).(*DataFetcherFiLStakeGovernanceFactors)

	return *outstruct, err

}

// FetchData is a free data retrieval call binding the contract method 0x67845e16.
//
// Solidity: function fetchData() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 fitTotalSupply, uint256 figTotalSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filLiquidInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filStakeInfo, (uint256,uint256,uint256) governanceInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fiLStakeGovernanceFactors)
func (_Data *DataSession) FetchData() (struct {
	BlockHeight                *big.Int
	BlockTimeStamp             *big.Int
	FitTotalSupply             *big.Int
	FigTotalSupply             *big.Int
	FilLiquidInfo              FILLiquidInterfaceFILLiquidInfo
	FilStakeInfo               FILStakeFILStakeInfo
	GovernanceInfo             GovernanceGovernanceInfo
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FiLStakeGovernanceFactors  DataFetcherFiLStakeGovernanceFactors
}, error) {
	return _Data.Contract.FetchData(&_Data.CallOpts)
}

// FetchData is a free data retrieval call binding the contract method 0x67845e16.
//
// Solidity: function fetchData() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 fitTotalSupply, uint256 figTotalSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filLiquidInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filStakeInfo, (uint256,uint256,uint256) governanceInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fiLStakeGovernanceFactors)
func (_Data *DataCallerSession) FetchData() (struct {
	BlockHeight                *big.Int
	BlockTimeStamp             *big.Int
	FitTotalSupply             *big.Int
	FigTotalSupply             *big.Int
	FilLiquidInfo              FILLiquidInterfaceFILLiquidInfo
	FilStakeInfo               FILStakeFILStakeInfo
	GovernanceInfo             GovernanceGovernanceInfo
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FiLStakeGovernanceFactors  DataFetcherFiLStakeGovernanceFactors
}, error) {
	return _Data.Contract.FetchData(&_Data.CallOpts)
}

// FetchGovernanceFactors is a free data retrieval call binding the contract method 0x9fd676d2.
//
// Solidity: function fetchGovernanceFactors() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fiLStakeGovernanceFactors)
func (_Data *DataCaller) FetchGovernanceFactors(opts *bind.CallOpts) (struct {
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FiLStakeGovernanceFactors  DataFetcherFiLStakeGovernanceFactors
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "fetchGovernanceFactors")

	outstruct := new(struct {
		FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
		FiLStakeGovernanceFactors  DataFetcherFiLStakeGovernanceFactors
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FiLLiquidGovernanceFactors = *abi.ConvertType(out[0], new(DataFetcherFiLLiquidGovernanceFactors)).(*DataFetcherFiLLiquidGovernanceFactors)
	outstruct.FiLStakeGovernanceFactors = *abi.ConvertType(out[1], new(DataFetcherFiLStakeGovernanceFactors)).(*DataFetcherFiLStakeGovernanceFactors)

	return *outstruct, err

}

// FetchGovernanceFactors is a free data retrieval call binding the contract method 0x9fd676d2.
//
// Solidity: function fetchGovernanceFactors() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fiLStakeGovernanceFactors)
func (_Data *DataSession) FetchGovernanceFactors() (struct {
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FiLStakeGovernanceFactors  DataFetcherFiLStakeGovernanceFactors
}, error) {
	return _Data.Contract.FetchGovernanceFactors(&_Data.CallOpts)
}

// FetchGovernanceFactors is a free data retrieval call binding the contract method 0x9fd676d2.
//
// Solidity: function fetchGovernanceFactors() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fiLStakeGovernanceFactors)
func (_Data *DataCallerSession) FetchGovernanceFactors() (struct {
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FiLStakeGovernanceFactors  DataFetcherFiLStakeGovernanceFactors
}, error) {
	return _Data.Contract.FetchGovernanceFactors(&_Data.CallOpts)
}

// FetchPersonalData is a free data retrieval call binding the contract method 0x6536c50f.
//
// Solidity: function fetchPersonalData(address player) view returns(uint256 filTrustBalance, uint256 filBalance, (address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) userInfo)
func (_Data *DataCaller) FetchPersonalData(opts *bind.CallOpts, player common.Address) (struct {
	FilTrustBalance *big.Int
	FilBalance      *big.Int
	UserInfo        FILLiquidInterfaceUserInfo
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "fetchPersonalData", player)

	outstruct := new(struct {
		FilTrustBalance *big.Int
		FilBalance      *big.Int
		UserInfo        FILLiquidInterfaceUserInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FilTrustBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FilBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UserInfo = *abi.ConvertType(out[2], new(FILLiquidInterfaceUserInfo)).(*FILLiquidInterfaceUserInfo)

	return *outstruct, err

}

// FetchPersonalData is a free data retrieval call binding the contract method 0x6536c50f.
//
// Solidity: function fetchPersonalData(address player) view returns(uint256 filTrustBalance, uint256 filBalance, (address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) userInfo)
func (_Data *DataSession) FetchPersonalData(player common.Address) (struct {
	FilTrustBalance *big.Int
	FilBalance      *big.Int
	UserInfo        FILLiquidInterfaceUserInfo
}, error) {
	return _Data.Contract.FetchPersonalData(&_Data.CallOpts, player)
}

// FetchPersonalData is a free data retrieval call binding the contract method 0x6536c50f.
//
// Solidity: function fetchPersonalData(address player) view returns(uint256 filTrustBalance, uint256 filBalance, (address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) userInfo)
func (_Data *DataCallerSession) FetchPersonalData(player common.Address) (struct {
	FilTrustBalance *big.Int
	FilBalance      *big.Int
	UserInfo        FILLiquidInterfaceUserInfo
}, error) {
	return _Data.Contract.FetchPersonalData(&_Data.CallOpts, player)
}

// FetchStakerData is a free data retrieval call binding the contract method 0x2b3c4cac.
//
// Solidity: function fetchStakerData(address staker) view returns(uint256 filTrustBalance, uint256 filTrustFixed, uint256 filTrustVariable, uint256 filGovernanceBalance)
func (_Data *DataCaller) FetchStakerData(opts *bind.CallOpts, staker common.Address) (struct {
	FilTrustBalance      *big.Int
	FilTrustFixed        *big.Int
	FilTrustVariable     *big.Int
	FilGovernanceBalance *big.Int
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "fetchStakerData", staker)

	outstruct := new(struct {
		FilTrustBalance      *big.Int
		FilTrustFixed        *big.Int
		FilTrustVariable     *big.Int
		FilGovernanceBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FilTrustBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FilTrustFixed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FilTrustVariable = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FilGovernanceBalance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FetchStakerData is a free data retrieval call binding the contract method 0x2b3c4cac.
//
// Solidity: function fetchStakerData(address staker) view returns(uint256 filTrustBalance, uint256 filTrustFixed, uint256 filTrustVariable, uint256 filGovernanceBalance)
func (_Data *DataSession) FetchStakerData(staker common.Address) (struct {
	FilTrustBalance      *big.Int
	FilTrustFixed        *big.Int
	FilTrustVariable     *big.Int
	FilGovernanceBalance *big.Int
}, error) {
	return _Data.Contract.FetchStakerData(&_Data.CallOpts, staker)
}

// FetchStakerData is a free data retrieval call binding the contract method 0x2b3c4cac.
//
// Solidity: function fetchStakerData(address staker) view returns(uint256 filTrustBalance, uint256 filTrustFixed, uint256 filTrustVariable, uint256 filGovernanceBalance)
func (_Data *DataCallerSession) FetchStakerData(staker common.Address) (struct {
	FilTrustBalance      *big.Int
	FilTrustFixed        *big.Int
	FilTrustVariable     *big.Int
	FilGovernanceBalance *big.Int
}, error) {
	return _Data.Contract.FetchStakerData(&_Data.CallOpts, staker)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[6])
func (_Data *DataCaller) GetAddresses(opts *bind.CallOpts) ([6]common.Address, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getAddresses")

	if err != nil {
		return *new([6]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([6]common.Address)).(*[6]common.Address)

	return out0, err

}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[6])
func (_Data *DataSession) GetAddresses() ([6]common.Address, error) {
	return _Data.Contract.GetAddresses(&_Data.CallOpts)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[6])
func (_Data *DataCallerSession) GetAddresses() ([6]common.Address, error) {
	return _Data.Contract.GetAddresses(&_Data.CallOpts)
}

// GetBatchedUserBorrows is a free data retrieval call binding the contract method 0xa405a88f.
//
// Solidity: function getBatchedUserBorrows(address[] accounts) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[])[] infos)
func (_Data *DataCaller) GetBatchedUserBorrows(opts *bind.CallOpts, accounts []common.Address) ([]FILLiquidInterfaceUserInfo, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getBatchedUserBorrows", accounts)

	if err != nil {
		return *new([]FILLiquidInterfaceUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]FILLiquidInterfaceUserInfo)).(*[]FILLiquidInterfaceUserInfo)

	return out0, err

}

// GetBatchedUserBorrows is a free data retrieval call binding the contract method 0xa405a88f.
//
// Solidity: function getBatchedUserBorrows(address[] accounts) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[])[] infos)
func (_Data *DataSession) GetBatchedUserBorrows(accounts []common.Address) ([]FILLiquidInterfaceUserInfo, error) {
	return _Data.Contract.GetBatchedUserBorrows(&_Data.CallOpts, accounts)
}

// GetBatchedUserBorrows is a free data retrieval call binding the contract method 0xa405a88f.
//
// Solidity: function getBatchedUserBorrows(address[] accounts) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[])[] infos)
func (_Data *DataCallerSession) GetBatchedUserBorrows(accounts []common.Address) ([]FILLiquidInterfaceUserInfo, error) {
	return _Data.Contract.GetBatchedUserBorrows(&_Data.CallOpts, accounts)
}

// GetBorrowExpecting is a free data retrieval call binding the contract method 0xec12bf50.
//
// Solidity: function getBorrowExpecting(uint256 amount) view returns(uint256 expectedInterestRate, uint256 sixMonthInterest)
func (_Data *DataCaller) GetBorrowExpecting(opts *bind.CallOpts, amount *big.Int) (struct {
	ExpectedInterestRate *big.Int
	SixMonthInterest     *big.Int
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getBorrowExpecting", amount)

	outstruct := new(struct {
		ExpectedInterestRate *big.Int
		SixMonthInterest     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExpectedInterestRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SixMonthInterest = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBorrowExpecting is a free data retrieval call binding the contract method 0xec12bf50.
//
// Solidity: function getBorrowExpecting(uint256 amount) view returns(uint256 expectedInterestRate, uint256 sixMonthInterest)
func (_Data *DataSession) GetBorrowExpecting(amount *big.Int) (struct {
	ExpectedInterestRate *big.Int
	SixMonthInterest     *big.Int
}, error) {
	return _Data.Contract.GetBorrowExpecting(&_Data.CallOpts, amount)
}

// GetBorrowExpecting is a free data retrieval call binding the contract method 0xec12bf50.
//
// Solidity: function getBorrowExpecting(uint256 amount) view returns(uint256 expectedInterestRate, uint256 sixMonthInterest)
func (_Data *DataCallerSession) GetBorrowExpecting(amount *big.Int) (struct {
	ExpectedInterestRate *big.Int
	SixMonthInterest     *big.Int
}, error) {
	return _Data.Contract.GetBorrowExpecting(&_Data.CallOpts, amount)
}

// GetBorrowable is a free data retrieval call binding the contract method 0x9685e716.
//
// Solidity: function getBorrowable(address account) view returns((bool,string)[] result)
func (_Data *DataCaller) GetBorrowable(opts *bind.CallOpts, account common.Address) ([]DataFetcherMinerBorrowable, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getBorrowable", account)

	if err != nil {
		return *new([]DataFetcherMinerBorrowable), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataFetcherMinerBorrowable)).(*[]DataFetcherMinerBorrowable)

	return out0, err

}

// GetBorrowable is a free data retrieval call binding the contract method 0x9685e716.
//
// Solidity: function getBorrowable(address account) view returns((bool,string)[] result)
func (_Data *DataSession) GetBorrowable(account common.Address) ([]DataFetcherMinerBorrowable, error) {
	return _Data.Contract.GetBorrowable(&_Data.CallOpts, account)
}

// GetBorrowable is a free data retrieval call binding the contract method 0x9685e716.
//
// Solidity: function getBorrowable(address account) view returns((bool,string)[] result)
func (_Data *DataCallerSession) GetBorrowable(account common.Address) ([]DataFetcherMinerBorrowable, error) {
	return _Data.Contract.GetBorrowable(&_Data.CallOpts, account)
}

// GetDepositExpecting is a free data retrieval call binding the contract method 0xcc798ad5.
//
// Solidity: function getDepositExpecting(uint256 amountFIL) view returns(uint256 expectedAmountFILTrust)
func (_Data *DataCaller) GetDepositExpecting(opts *bind.CallOpts, amountFIL *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getDepositExpecting", amountFIL)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositExpecting is a free data retrieval call binding the contract method 0xcc798ad5.
//
// Solidity: function getDepositExpecting(uint256 amountFIL) view returns(uint256 expectedAmountFILTrust)
func (_Data *DataSession) GetDepositExpecting(amountFIL *big.Int) (*big.Int, error) {
	return _Data.Contract.GetDepositExpecting(&_Data.CallOpts, amountFIL)
}

// GetDepositExpecting is a free data retrieval call binding the contract method 0xcc798ad5.
//
// Solidity: function getDepositExpecting(uint256 amountFIL) view returns(uint256 expectedAmountFILTrust)
func (_Data *DataCallerSession) GetDepositExpecting(amountFIL *big.Int) (*big.Int, error) {
	return _Data.Contract.GetDepositExpecting(&_Data.CallOpts, amountFIL)
}

// GetPendingBeneficiary is a free data retrieval call binding the contract method 0x314befa3.
//
// Solidity: function getPendingBeneficiary(uint64 minerId) view returns(address, bool)
func (_Data *DataCaller) GetPendingBeneficiary(opts *bind.CallOpts, minerId uint64) (common.Address, bool, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getPendingBeneficiary", minerId)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetPendingBeneficiary is a free data retrieval call binding the contract method 0x314befa3.
//
// Solidity: function getPendingBeneficiary(uint64 minerId) view returns(address, bool)
func (_Data *DataSession) GetPendingBeneficiary(minerId uint64) (common.Address, bool, error) {
	return _Data.Contract.GetPendingBeneficiary(&_Data.CallOpts, minerId)
}

// GetPendingBeneficiary is a free data retrieval call binding the contract method 0x314befa3.
//
// Solidity: function getPendingBeneficiary(uint64 minerId) view returns(address, bool)
func (_Data *DataCallerSession) GetPendingBeneficiary(minerId uint64) (common.Address, bool, error) {
	return _Data.Contract.GetPendingBeneficiary(&_Data.CallOpts, minerId)
}

// GetRedeemExpecting is a free data retrieval call binding the contract method 0xea5565a7.
//
// Solidity: function getRedeemExpecting(uint256 amountFILTrust) view returns(uint256 expectedAmountFIL)
func (_Data *DataCaller) GetRedeemExpecting(opts *bind.CallOpts, amountFILTrust *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getRedeemExpecting", amountFILTrust)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedeemExpecting is a free data retrieval call binding the contract method 0xea5565a7.
//
// Solidity: function getRedeemExpecting(uint256 amountFILTrust) view returns(uint256 expectedAmountFIL)
func (_Data *DataSession) GetRedeemExpecting(amountFILTrust *big.Int) (*big.Int, error) {
	return _Data.Contract.GetRedeemExpecting(&_Data.CallOpts, amountFILTrust)
}

// GetRedeemExpecting is a free data retrieval call binding the contract method 0xea5565a7.
//
// Solidity: function getRedeemExpecting(uint256 amountFILTrust) view returns(uint256 expectedAmountFIL)
func (_Data *DataCallerSession) GetRedeemExpecting(amountFILTrust *big.Int) (*big.Int, error) {
	return _Data.Contract.GetRedeemExpecting(&_Data.CallOpts, amountFILTrust)
}

// GetTotalPendingInterest is a free data retrieval call binding the contract method 0x855c55f3.
//
// Solidity: function getTotalPendingInterest() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 totalPendingInterest, uint256 totalFIL, uint256 borrowing, uint256 borrowingAndPeriod, uint256 accumulatedPayback, uint256 interestExp)
func (_Data *DataCaller) GetTotalPendingInterest(opts *bind.CallOpts) (struct {
	BlockHeight          *big.Int
	BlockTimeStamp       *big.Int
	TotalPendingInterest *big.Int
	TotalFIL             *big.Int
	Borrowing            *big.Int
	BorrowingAndPeriod   *big.Int
	AccumulatedPayback   *big.Int
	InterestExp          *big.Int
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getTotalPendingInterest")

	outstruct := new(struct {
		BlockHeight          *big.Int
		BlockTimeStamp       *big.Int
		TotalPendingInterest *big.Int
		TotalFIL             *big.Int
		Borrowing            *big.Int
		BorrowingAndPeriod   *big.Int
		AccumulatedPayback   *big.Int
		InterestExp          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockTimeStamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalPendingInterest = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalFIL = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Borrowing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BorrowingAndPeriod = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.AccumulatedPayback = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.InterestExp = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTotalPendingInterest is a free data retrieval call binding the contract method 0x855c55f3.
//
// Solidity: function getTotalPendingInterest() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 totalPendingInterest, uint256 totalFIL, uint256 borrowing, uint256 borrowingAndPeriod, uint256 accumulatedPayback, uint256 interestExp)
func (_Data *DataSession) GetTotalPendingInterest() (struct {
	BlockHeight          *big.Int
	BlockTimeStamp       *big.Int
	TotalPendingInterest *big.Int
	TotalFIL             *big.Int
	Borrowing            *big.Int
	BorrowingAndPeriod   *big.Int
	AccumulatedPayback   *big.Int
	InterestExp          *big.Int
}, error) {
	return _Data.Contract.GetTotalPendingInterest(&_Data.CallOpts)
}

// GetTotalPendingInterest is a free data retrieval call binding the contract method 0x855c55f3.
//
// Solidity: function getTotalPendingInterest() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 totalPendingInterest, uint256 totalFIL, uint256 borrowing, uint256 borrowingAndPeriod, uint256 accumulatedPayback, uint256 interestExp)
func (_Data *DataCallerSession) GetTotalPendingInterest() (struct {
	BlockHeight          *big.Int
	BlockTimeStamp       *big.Int
	TotalPendingInterest *big.Int
	TotalFIL             *big.Int
	Borrowing            *big.Int
	BorrowingAndPeriod   *big.Int
	AccumulatedPayback   *big.Int
	InterestExp          *big.Int
}, error) {
	return _Data.Contract.GetTotalPendingInterest(&_Data.CallOpts)
}

// GetUserBorrowsAndBorrowable is a free data retrieval call binding the contract method 0xb8ce16e4.
//
// Solidity: function getUserBorrowsAndBorrowable(address account) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) info, (bool,string)[] borrowables)
func (_Data *DataCaller) GetUserBorrowsAndBorrowable(opts *bind.CallOpts, account common.Address) (struct {
	Info        FILLiquidInterfaceUserInfo
	Borrowables []DataFetcherMinerBorrowable
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getUserBorrowsAndBorrowable", account)

	outstruct := new(struct {
		Info        FILLiquidInterfaceUserInfo
		Borrowables []DataFetcherMinerBorrowable
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Info = *abi.ConvertType(out[0], new(FILLiquidInterfaceUserInfo)).(*FILLiquidInterfaceUserInfo)
	outstruct.Borrowables = *abi.ConvertType(out[1], new([]DataFetcherMinerBorrowable)).(*[]DataFetcherMinerBorrowable)

	return *outstruct, err

}

// GetUserBorrowsAndBorrowable is a free data retrieval call binding the contract method 0xb8ce16e4.
//
// Solidity: function getUserBorrowsAndBorrowable(address account) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) info, (bool,string)[] borrowables)
func (_Data *DataSession) GetUserBorrowsAndBorrowable(account common.Address) (struct {
	Info        FILLiquidInterfaceUserInfo
	Borrowables []DataFetcherMinerBorrowable
}, error) {
	return _Data.Contract.GetUserBorrowsAndBorrowable(&_Data.CallOpts, account)
}

// GetUserBorrowsAndBorrowable is a free data retrieval call binding the contract method 0xb8ce16e4.
//
// Solidity: function getUserBorrowsAndBorrowable(address account) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) info, (bool,string)[] borrowables)
func (_Data *DataCallerSession) GetUserBorrowsAndBorrowable(account common.Address) (struct {
	Info        FILLiquidInterfaceUserInfo
	Borrowables []DataFetcherMinerBorrowable
}, error) {
	return _Data.Contract.GetUserBorrowsAndBorrowable(&_Data.CallOpts, account)
}

// GetUserBorrowsByMiner is a free data retrieval call binding the contract method 0xc4591100.
//
// Solidity: function getUserBorrowsByMiner(uint64 minerId) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) infos)
func (_Data *DataCaller) GetUserBorrowsByMiner(opts *bind.CallOpts, minerId uint64) (FILLiquidInterfaceUserInfo, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getUserBorrowsByMiner", minerId)

	if err != nil {
		return *new(FILLiquidInterfaceUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FILLiquidInterfaceUserInfo)).(*FILLiquidInterfaceUserInfo)

	return out0, err

}

// GetUserBorrowsByMiner is a free data retrieval call binding the contract method 0xc4591100.
//
// Solidity: function getUserBorrowsByMiner(uint64 minerId) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) infos)
func (_Data *DataSession) GetUserBorrowsByMiner(minerId uint64) (FILLiquidInterfaceUserInfo, error) {
	return _Data.Contract.GetUserBorrowsByMiner(&_Data.CallOpts, minerId)
}

// GetUserBorrowsByMiner is a free data retrieval call binding the contract method 0xc4591100.
//
// Solidity: function getUserBorrowsByMiner(uint64 minerId) view returns((address,uint256,uint256,uint256,uint256,(uint256,bool,bool),(uint64,uint256,uint256,uint256,(uint256,bool),((uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])[]) infos)
func (_Data *DataCallerSession) GetUserBorrowsByMiner(minerId uint64) (FILLiquidInterfaceUserInfo, error) {
	return _Data.Contract.GetUserBorrowsByMiner(&_Data.CallOpts, minerId)
}

// MaxBorrowAllowed is a free data retrieval call binding the contract method 0x91546352.
//
// Solidity: function maxBorrowAllowed(uint64 minerId) view returns(uint256 amount)
func (_Data *DataCaller) MaxBorrowAllowed(opts *bind.CallOpts, minerId uint64) (*big.Int, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "maxBorrowAllowed", minerId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowAllowed is a free data retrieval call binding the contract method 0x91546352.
//
// Solidity: function maxBorrowAllowed(uint64 minerId) view returns(uint256 amount)
func (_Data *DataSession) MaxBorrowAllowed(minerId uint64) (*big.Int, error) {
	return _Data.Contract.MaxBorrowAllowed(&_Data.CallOpts, minerId)
}

// MaxBorrowAllowed is a free data retrieval call binding the contract method 0x91546352.
//
// Solidity: function maxBorrowAllowed(uint64 minerId) view returns(uint256 amount)
func (_Data *DataCallerSession) MaxBorrowAllowed(minerId uint64) (*big.Int, error) {
	return _Data.Contract.MaxBorrowAllowed(&_Data.CallOpts, minerId)
}

// MaxBorrowAllowedInAdvance is a free data retrieval call binding the contract method 0xa71d4d76.
//
// Solidity: function maxBorrowAllowedInAdvance(uint64 minerId, uint256 afterBlocks) view returns(uint256 amount)
func (_Data *DataCaller) MaxBorrowAllowedInAdvance(opts *bind.CallOpts, minerId uint64, afterBlocks *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "maxBorrowAllowedInAdvance", minerId, afterBlocks)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowAllowedInAdvance is a free data retrieval call binding the contract method 0xa71d4d76.
//
// Solidity: function maxBorrowAllowedInAdvance(uint64 minerId, uint256 afterBlocks) view returns(uint256 amount)
func (_Data *DataSession) MaxBorrowAllowedInAdvance(minerId uint64, afterBlocks *big.Int) (*big.Int, error) {
	return _Data.Contract.MaxBorrowAllowedInAdvance(&_Data.CallOpts, minerId, afterBlocks)
}

// MaxBorrowAllowedInAdvance is a free data retrieval call binding the contract method 0xa71d4d76.
//
// Solidity: function maxBorrowAllowedInAdvance(uint64 minerId, uint256 afterBlocks) view returns(uint256 amount)
func (_Data *DataCallerSession) MaxBorrowAllowedInAdvance(minerId uint64, afterBlocks *big.Int) (*big.Int, error) {
	return _Data.Contract.MaxBorrowAllowedInAdvance(&_Data.CallOpts, minerId, afterBlocks)
}