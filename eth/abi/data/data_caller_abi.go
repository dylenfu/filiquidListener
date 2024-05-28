// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package data

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

// DataFetcherFITStakeGovernanceFactors is an auto generated low-level Go binding around an user-defined struct.
type DataFetcherFITStakeGovernanceFactors struct {
	NInterest *big.Int
	NStake    *big.Int
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
}

// DataFetcherMinerBorrowable is an auto generated low-level Go binding around an user-defined struct.
type DataFetcherMinerBorrowable struct {
	Borrowable bool
	Reason     string
}

// DataFetcherVotingResult is an auto generated low-level Go binding around an user-defined struct.
type DataFetcherVotingResult struct {
	AmountTotal      *big.Int
	AmountYes        *big.Int
	AmountNo         *big.Int
	AmountNoWithVeto *big.Int
	AmountAbstain    *big.Int
	Result           uint8
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

// FITStakeFITStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type FITStakeFITStakeInfo struct {
	AccumulatedInterest      *big.Int
	AccumulatedStake         *big.Int
	AccumulatedStakeDuration *big.Int
	AccumulatedInterestMint  *big.Int
	AccumulatedStakeMint     *big.Int
	AccumulatedWithdrawn     *big.Int
	NextStakeID              *big.Int
	ReleasedFIGStake         *big.Int
}

// GovernanceGovernanceInfo is an auto generated low-level Go binding around an user-defined struct.
type GovernanceGovernanceInfo struct {
	Bonders               *big.Int
	TotalBondedAmount     *big.Int
	FirstActiveProposalId *big.Int
}

// GovernanceProposalInfo is an auto generated low-level Go binding around an user-defined struct.
type GovernanceProposalInfo struct {
	Category        uint8
	Start           *big.Int
	Deadline        *big.Int
	Deposited       *big.Int
	DiscussionIndex *big.Int
	Executed        bool
	Result          uint8
	Text            string
	Proposer        common.Address
	Values          []*big.Int
}

// DataMetaData contains all meta data concerning the Data contract.
var DataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractFILLiquid\",\"name\":\"filLiquidAddr\",\"type\":\"address\"},{\"internalType\":\"contractFILTrust\",\"name\":\"filTrustAddr\",\"type\":\"address\"},{\"internalType\":\"contractFITStake\",\"name\":\"fitStakeAddr\",\"type\":\"address\"},{\"internalType\":\"contractFILGovernance\",\"name\":\"filGovernanceAddr\",\"type\":\"address\"},{\"internalType\":\"contractGovernance\",\"name\":\"governanceAddr\",\"type\":\"address\"},{\"internalType\":\"contractFilecoinAPI\",\"name\":\"filecoinAPIAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"fetchData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fitTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"figTotalSupply\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalFIL\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableFIL\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBurntFILTrust\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedMintFILTrust\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedPayback\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRedeemFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBorrowFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBadDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedLiquidateReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedLiquidateFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralizedMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minerWithBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateBase\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.FILLiquidInfo\",\"name\":\"filLiquidInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accumulatedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedStakeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedInterestMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedStakeMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedWithdrawn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextStakeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releasedFIGStake\",\"type\":\"uint256\"}],\"internalType\":\"structFITStake.FITStakeInfo\",\"name\":\"fitStakeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bonders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBondedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstActiveProposalId\",\"type\":\"uint256\"}],\"internalType\":\"structGovernance.GovernanceInfo\",\"name\":\"governanceInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"u_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFamilySize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alertThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidateInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateDiscountRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExistingBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDepositAmount\",\"type\":\"uint256\"}],\"internalType\":\"structDataFetcher.FiLLiquidGovernanceFactors\",\"name\":\"fiLLiquidGovernanceFactors\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"n_interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n_stake\",\"type\":\"uint256\"}],\"internalType\":\"structDataFetcher.FITStakeGovernanceFactors\",\"name\":\"fitStakeGovernanceFactors\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchGovernanceFactors\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"u_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFamilySize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alertThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidateInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateDiscountRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidateFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExistingBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDepositAmount\",\"type\":\"uint256\"}],\"internalType\":\"structDataFetcher.FiLLiquidGovernanceFactors\",\"name\":\"fiLLiquidGovernanceFactors\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"n_interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n_stake\",\"type\":\"uint256\"}],\"internalType\":\"structDataFetcher.FITStakeGovernanceFactors\",\"name\":\"fitStakeGovernanceFactors\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"fetchPersonalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filTrustBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filBalance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alertable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"liquidatable\",\"type\":\"bool\"}],\"internalType\":\"structFILLiquidInterface.LiquidateConditionInfo\",\"name\":\"liquidateConditionInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structConversion.Integer\",\"name\":\"availableBalance\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingOriginalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"datedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInfo\",\"name\":\"borrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInterestInfo[]\",\"name\":\"borrows\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.MinerBorrowInfo[]\",\"name\":\"minerBorrowInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.UserInfo\",\"name\":\"userInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"fetchStakerData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filTrustBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFIGSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releasedFIGSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filTrustVariable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"canWithdrawFIGSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filTrustLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filGovernanceBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[6]\",\"name\":\"\",\"type\":\"address[6]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"getBatchedUserBorrows\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alertable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"liquidatable\",\"type\":\"bool\"}],\"internalType\":\"structFILLiquidInterface.LiquidateConditionInfo\",\"name\":\"liquidateConditionInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structConversion.Integer\",\"name\":\"availableBalance\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingOriginalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"datedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInfo\",\"name\":\"borrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInterestInfo[]\",\"name\":\"borrows\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.MinerBorrowInfo[]\",\"name\":\"minerBorrowInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.UserInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBondsOnProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"votedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBondedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBorrowExpecting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sixMonthInterest\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBorrowable\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"borrowable\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structDataFetcher.MinerBorrowable[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountFIL\",\"type\":\"uint256\"}],\"name\":\"getDepositExpecting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedAmountFILTrust\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getFigAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"figBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondableAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"}],\"name\":\"getPendingBeneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"enumGovernance.proposalCategory\",\"name\":\"category\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discussionIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"enumGovernance.voteResult\",\"name\":\"result\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"internalType\":\"structGovernance.ProposalInfo[]\",\"name\":\"proposals\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountNoWithVeto\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAbstain\",\"type\":\"uint256\"},{\"internalType\":\"enumGovernance.voteResult\",\"name\":\"result\",\"type\":\"uint8\"}],\"internalType\":\"structDataFetcher.VotingResult[]\",\"name\":\"votings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountFILTrust\",\"type\":\"uint256\"}],\"name\":\"getRedeemExpecting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedAmountFIL\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPendingInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPendingInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFIL\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingAndPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedPayback\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestExp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserBorrowsAndBorrowable\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alertable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"liquidatable\",\"type\":\"bool\"}],\"internalType\":\"structFILLiquidInterface.LiquidateConditionInfo\",\"name\":\"liquidateConditionInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structConversion.Integer\",\"name\":\"availableBalance\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingOriginalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"datedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInfo\",\"name\":\"borrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInterestInfo[]\",\"name\":\"borrows\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.MinerBorrowInfo[]\",\"name\":\"minerBorrowInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.UserInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"borrowable\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structDataFetcher.MinerBorrowable[]\",\"name\":\"borrowables\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"}],\"name\":\"getUserBorrowsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableCredit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alertable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"liquidatable\",\"type\":\"bool\"}],\"internalType\":\"structFILLiquidInterface.LiquidateConditionInfo\",\"name\":\"liquidateConditionInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"debtOutStanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structConversion.Integer\",\"name\":\"availableBalance\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingOriginalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"datedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInfo\",\"name\":\"borrow\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"internalType\":\"structFILLiquidInterface.BorrowInterestInfo[]\",\"name\":\"borrows\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.MinerBorrowInfo[]\",\"name\":\"minerBorrowInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structFILLiquidInterface.UserInfo\",\"name\":\"infos\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"}],\"name\":\"maxBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"afterBlocks\",\"type\":\"uint256\"}],\"name\":\"maxBorrowAllowedInAdvance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DataABI is the input ABI used to generate the binding from.
// Deprecated: Use DataMetaData.ABI instead.
var DataABI = DataMetaData.ABI

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
// Solidity: function fetchData() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 fitTotalSupply, uint256 figTotalSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filLiquidInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fitStakeInfo, (uint256,uint256,uint256) governanceInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fitStakeGovernanceFactors)
func (_Data *DataCaller) FetchData(opts *bind.CallOpts) (struct {
	BlockHeight                *big.Int
	BlockTimeStamp             *big.Int
	FitTotalSupply             *big.Int
	FigTotalSupply             *big.Int
	FilLiquidInfo              FILLiquidInterfaceFILLiquidInfo
	FitStakeInfo               FITStakeFITStakeInfo
	GovernanceInfo             GovernanceGovernanceInfo
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FitStakeGovernanceFactors  DataFetcherFITStakeGovernanceFactors
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "fetchData")

	outstruct := new(struct {
		BlockHeight                *big.Int
		BlockTimeStamp             *big.Int
		FitTotalSupply             *big.Int
		FigTotalSupply             *big.Int
		FilLiquidInfo              FILLiquidInterfaceFILLiquidInfo
		FitStakeInfo               FITStakeFITStakeInfo
		GovernanceInfo             GovernanceGovernanceInfo
		FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
		FitStakeGovernanceFactors  DataFetcherFITStakeGovernanceFactors
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockTimeStamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FitTotalSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FigTotalSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FilLiquidInfo = *abi.ConvertType(out[4], new(FILLiquidInterfaceFILLiquidInfo)).(*FILLiquidInterfaceFILLiquidInfo)
	outstruct.FitStakeInfo = *abi.ConvertType(out[5], new(FITStakeFITStakeInfo)).(*FITStakeFITStakeInfo)
	outstruct.GovernanceInfo = *abi.ConvertType(out[6], new(GovernanceGovernanceInfo)).(*GovernanceGovernanceInfo)
	outstruct.FiLLiquidGovernanceFactors = *abi.ConvertType(out[7], new(DataFetcherFiLLiquidGovernanceFactors)).(*DataFetcherFiLLiquidGovernanceFactors)
	outstruct.FitStakeGovernanceFactors = *abi.ConvertType(out[8], new(DataFetcherFITStakeGovernanceFactors)).(*DataFetcherFITStakeGovernanceFactors)

	return *outstruct, err

}

// FetchData is a free data retrieval call binding the contract method 0x67845e16.
//
// Solidity: function fetchData() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 fitTotalSupply, uint256 figTotalSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filLiquidInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fitStakeInfo, (uint256,uint256,uint256) governanceInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fitStakeGovernanceFactors)
func (_Data *DataSession) FetchData() (struct {
	BlockHeight                *big.Int
	BlockTimeStamp             *big.Int
	FitTotalSupply             *big.Int
	FigTotalSupply             *big.Int
	FilLiquidInfo              FILLiquidInterfaceFILLiquidInfo
	FitStakeInfo               FITStakeFITStakeInfo
	GovernanceInfo             GovernanceGovernanceInfo
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FitStakeGovernanceFactors  DataFetcherFITStakeGovernanceFactors
}, error) {
	return _Data.Contract.FetchData(&_Data.CallOpts)
}

// FetchData is a free data retrieval call binding the contract method 0x67845e16.
//
// Solidity: function fetchData() view returns(uint256 blockHeight, uint256 blockTimeStamp, uint256 fitTotalSupply, uint256 figTotalSupply, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) filLiquidInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fitStakeInfo, (uint256,uint256,uint256) governanceInfo, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fitStakeGovernanceFactors)
func (_Data *DataCallerSession) FetchData() (struct {
	BlockHeight                *big.Int
	BlockTimeStamp             *big.Int
	FitTotalSupply             *big.Int
	FigTotalSupply             *big.Int
	FilLiquidInfo              FILLiquidInterfaceFILLiquidInfo
	FitStakeInfo               FITStakeFITStakeInfo
	GovernanceInfo             GovernanceGovernanceInfo
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FitStakeGovernanceFactors  DataFetcherFITStakeGovernanceFactors
}, error) {
	return _Data.Contract.FetchData(&_Data.CallOpts)
}

// FetchGovernanceFactors is a free data retrieval call binding the contract method 0x9fd676d2.
//
// Solidity: function fetchGovernanceFactors() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fitStakeGovernanceFactors)
func (_Data *DataCaller) FetchGovernanceFactors(opts *bind.CallOpts) (struct {
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FitStakeGovernanceFactors  DataFetcherFITStakeGovernanceFactors
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "fetchGovernanceFactors")

	outstruct := new(struct {
		FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
		FitStakeGovernanceFactors  DataFetcherFITStakeGovernanceFactors
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FiLLiquidGovernanceFactors = *abi.ConvertType(out[0], new(DataFetcherFiLLiquidGovernanceFactors)).(*DataFetcherFiLLiquidGovernanceFactors)
	outstruct.FitStakeGovernanceFactors = *abi.ConvertType(out[1], new(DataFetcherFITStakeGovernanceFactors)).(*DataFetcherFITStakeGovernanceFactors)

	return *outstruct, err

}

// FetchGovernanceFactors is a free data retrieval call binding the contract method 0x9fd676d2.
//
// Solidity: function fetchGovernanceFactors() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fitStakeGovernanceFactors)
func (_Data *DataSession) FetchGovernanceFactors() (struct {
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FitStakeGovernanceFactors  DataFetcherFITStakeGovernanceFactors
}, error) {
	return _Data.Contract.FetchGovernanceFactors(&_Data.CallOpts)
}

// FetchGovernanceFactors is a free data retrieval call binding the contract method 0x9fd676d2.
//
// Solidity: function fetchGovernanceFactors() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) fiLLiquidGovernanceFactors, (uint256,uint256) fitStakeGovernanceFactors)
func (_Data *DataCallerSession) FetchGovernanceFactors() (struct {
	FiLLiquidGovernanceFactors DataFetcherFiLLiquidGovernanceFactors
	FitStakeGovernanceFactors  DataFetcherFITStakeGovernanceFactors
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
// Solidity: function fetchStakerData(address staker) view returns(uint256 filTrustBalance, uint256 stakeSum, uint256 totalFIGSum, uint256 releasedFIGSum, uint256 filTrustVariable, uint256 canWithdrawFIGSum, uint256 filTrustLocked, uint256 filGovernanceBalance)
func (_Data *DataCaller) FetchStakerData(opts *bind.CallOpts, staker common.Address) (struct {
	FilTrustBalance      *big.Int
	StakeSum             *big.Int
	TotalFIGSum          *big.Int
	ReleasedFIGSum       *big.Int
	FilTrustVariable     *big.Int
	CanWithdrawFIGSum    *big.Int
	FilTrustLocked       *big.Int
	FilGovernanceBalance *big.Int
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "fetchStakerData", staker)

	outstruct := new(struct {
		FilTrustBalance      *big.Int
		StakeSum             *big.Int
		TotalFIGSum          *big.Int
		ReleasedFIGSum       *big.Int
		FilTrustVariable     *big.Int
		CanWithdrawFIGSum    *big.Int
		FilTrustLocked       *big.Int
		FilGovernanceBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FilTrustBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakeSum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalFIGSum = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReleasedFIGSum = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FilTrustVariable = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CanWithdrawFIGSum = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.FilTrustLocked = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.FilGovernanceBalance = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FetchStakerData is a free data retrieval call binding the contract method 0x2b3c4cac.
//
// Solidity: function fetchStakerData(address staker) view returns(uint256 filTrustBalance, uint256 stakeSum, uint256 totalFIGSum, uint256 releasedFIGSum, uint256 filTrustVariable, uint256 canWithdrawFIGSum, uint256 filTrustLocked, uint256 filGovernanceBalance)
func (_Data *DataSession) FetchStakerData(staker common.Address) (struct {
	FilTrustBalance      *big.Int
	StakeSum             *big.Int
	TotalFIGSum          *big.Int
	ReleasedFIGSum       *big.Int
	FilTrustVariable     *big.Int
	CanWithdrawFIGSum    *big.Int
	FilTrustLocked       *big.Int
	FilGovernanceBalance *big.Int
}, error) {
	return _Data.Contract.FetchStakerData(&_Data.CallOpts, staker)
}

// FetchStakerData is a free data retrieval call binding the contract method 0x2b3c4cac.
//
// Solidity: function fetchStakerData(address staker) view returns(uint256 filTrustBalance, uint256 stakeSum, uint256 totalFIGSum, uint256 releasedFIGSum, uint256 filTrustVariable, uint256 canWithdrawFIGSum, uint256 filTrustLocked, uint256 filGovernanceBalance)
func (_Data *DataCallerSession) FetchStakerData(staker common.Address) (struct {
	FilTrustBalance      *big.Int
	StakeSum             *big.Int
	TotalFIGSum          *big.Int
	ReleasedFIGSum       *big.Int
	FilTrustVariable     *big.Int
	CanWithdrawFIGSum    *big.Int
	FilTrustLocked       *big.Int
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

// GetBondsOnProposal is a free data retrieval call binding the contract method 0xf807bb70.
//
// Solidity: function getBondsOnProposal(uint256 proposalId, address account) view returns(uint256 votedAmount, uint256 bondedAmount, uint256 totalBondedAmount)
func (_Data *DataCaller) GetBondsOnProposal(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (struct {
	VotedAmount       *big.Int
	BondedAmount      *big.Int
	TotalBondedAmount *big.Int
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getBondsOnProposal", proposalId, account)

	outstruct := new(struct {
		VotedAmount       *big.Int
		BondedAmount      *big.Int
		TotalBondedAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VotedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BondedAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBondedAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBondsOnProposal is a free data retrieval call binding the contract method 0xf807bb70.
//
// Solidity: function getBondsOnProposal(uint256 proposalId, address account) view returns(uint256 votedAmount, uint256 bondedAmount, uint256 totalBondedAmount)
func (_Data *DataSession) GetBondsOnProposal(proposalId *big.Int, account common.Address) (struct {
	VotedAmount       *big.Int
	BondedAmount      *big.Int
	TotalBondedAmount *big.Int
}, error) {
	return _Data.Contract.GetBondsOnProposal(&_Data.CallOpts, proposalId, account)
}

// GetBondsOnProposal is a free data retrieval call binding the contract method 0xf807bb70.
//
// Solidity: function getBondsOnProposal(uint256 proposalId, address account) view returns(uint256 votedAmount, uint256 bondedAmount, uint256 totalBondedAmount)
func (_Data *DataCallerSession) GetBondsOnProposal(proposalId *big.Int, account common.Address) (struct {
	VotedAmount       *big.Int
	BondedAmount      *big.Int
	TotalBondedAmount *big.Int
}, error) {
	return _Data.Contract.GetBondsOnProposal(&_Data.CallOpts, proposalId, account)
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

// GetProposals is a free data retrieval call binding the contract method 0xf5f3bf23.
//
// Solidity: function getProposals(uint256[] ids) view returns((uint8,uint256,uint256,uint256,uint256,bool,uint8,string,address,uint256[])[] proposals, (uint256,uint256,uint256,uint256,uint256,uint8)[] votings)
func (_Data *DataCaller) GetProposals(opts *bind.CallOpts, ids []*big.Int) (struct {
	Proposals []GovernanceProposalInfo
	Votings   []DataFetcherVotingResult
}, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getProposals", ids)

	outstruct := new(struct {
		Proposals []GovernanceProposalInfo
		Votings   []DataFetcherVotingResult
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposals = *abi.ConvertType(out[0], new([]GovernanceProposalInfo)).(*[]GovernanceProposalInfo)
	outstruct.Votings = *abi.ConvertType(out[1], new([]DataFetcherVotingResult)).(*[]DataFetcherVotingResult)

	return *outstruct, err

}

// GetProposals is a free data retrieval call binding the contract method 0xf5f3bf23.
//
// Solidity: function getProposals(uint256[] ids) view returns((uint8,uint256,uint256,uint256,uint256,bool,uint8,string,address,uint256[])[] proposals, (uint256,uint256,uint256,uint256,uint256,uint8)[] votings)
func (_Data *DataSession) GetProposals(ids []*big.Int) (struct {
	Proposals []GovernanceProposalInfo
	Votings   []DataFetcherVotingResult
}, error) {
	return _Data.Contract.GetProposals(&_Data.CallOpts, ids)
}

// GetProposals is a free data retrieval call binding the contract method 0xf5f3bf23.
//
// Solidity: function getProposals(uint256[] ids) view returns((uint8,uint256,uint256,uint256,uint256,bool,uint8,string,address,uint256[])[] proposals, (uint256,uint256,uint256,uint256,uint256,uint8)[] votings)
func (_Data *DataCallerSession) GetProposals(ids []*big.Int) (struct {
	Proposals []GovernanceProposalInfo
	Votings   []DataFetcherVotingResult
}, error) {
	return _Data.Contract.GetProposals(&_Data.CallOpts, ids)
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

// GetFigAmounts is a paid mutator transaction binding the contract method 0xb29a9641.
//
// Solidity: function getFigAmounts(address account) returns(uint256 figBalance, uint256 bondedAmount, uint256 unbondableAmount)
func (_Data *DataTransactor) GetFigAmounts(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "getFigAmounts", account)
}

// GetFigAmounts is a paid mutator transaction binding the contract method 0xb29a9641.
//
// Solidity: function getFigAmounts(address account) returns(uint256 figBalance, uint256 bondedAmount, uint256 unbondableAmount)
func (_Data *DataSession) GetFigAmounts(account common.Address) (*types.Transaction, error) {
	return _Data.Contract.GetFigAmounts(&_Data.TransactOpts, account)
}

// GetFigAmounts is a paid mutator transaction binding the contract method 0xb29a9641.
//
// Solidity: function getFigAmounts(address account) returns(uint256 figBalance, uint256 bondedAmount, uint256 unbondableAmount)
func (_Data *DataTransactorSession) GetFigAmounts(account common.Address) (*types.Transaction, error) {
	return _Data.Contract.GetFigAmounts(&_Data.TransactOpts, account)
}
