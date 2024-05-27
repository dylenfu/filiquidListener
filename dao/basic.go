package dao

import (
	"fmt"
	"math/big"

	"github.com/filiquid/listener/eth/abi/fetcher"
	"github.com/filiquid/listener/model"
)

type BasicData struct {
	BlockHeight    uint64 `gorm:"primaryKey"`
	BlockTimeStamp uint64

	FitTotalSupply    string
	FigTotalSupply    string
	TotalFIL          string
	AvailableFIL      string
	UtilizedLiquidity string

	AccumulatedDeposit         string
	AccumulatedRedeem          string
	AccumulatedBurntFILTrust   string
	AccumulatedMintFILTrust    string
	AccumulatedBorrow          string
	AccumulatedPayback         string
	AccumulatedInterest        string
	AccumulatedRedeemFee       string
	AccumulatedBorrowFee       string
	AccumulatedLiquidateReward string
	AccumulatedLiquidateFee    string
	AccumulatedDeposits        string
	AccumulatedBorrows         string
	UtilizationRate            string
	ExchangeRate               string
	InterestRate               string
	CollateralizedMiner        string
	MinerWithBorrows           string
	RateBase                   string

	AccumulatedStake         string
	AccumulatedStakeDuration string
	AccumulatedInterestMint  string
	AccumulatedStakeMint     string
	AccumulatedWithdrawn     string
	NextStakeID              string
	ReleasedFigStake         string

	Bonders               string
	TotalBondedAmount     string
	FirstActiveProposalId string

	U1                    string
	R0                    string
	R1                    string
	RM                    string
	CollateralRate        string
	MaxFamilySize         string
	AlertThreshold        string
	LiquidateThreshold    string
	MaxLiquidations       string
	MinLiquidateInterval  string
	LiquidateDiscountRate string
	LiquidateFeeRate      string
	MaxExistingBorrows    string
	MinBorrowAmount       string
	MinDepositAmount      string
	N                     string
	NInterest             string
	NStake                string
}

func (b *BasicData) Up(basicData *struct {
	BlockHeight                *big.Int
	BlockTimeStamp             *big.Int
	FitTotalSupply             *big.Int
	FigTotalSupply             *big.Int
	FilLiquidInfo              fetcher.FILLiquidInterfaceFILLiquidInfo
	FilStakeInfo               fetcher.FILStakeFILStakeInfo
	GovernanceInfo             fetcher.GovernanceGovernanceInfo
	FiLLiquidGovernanceFactors fetcher.DataFetcherFiLLiquidGovernanceFactors
	FiLStakeGovernanceFactors  fetcher.DataFetcherFiLStakeGovernanceFactors
}) *BasicData {
	b.BlockHeight = basicData.BlockHeight.Uint64()
	b.BlockTimeStamp = basicData.BlockTimeStamp.Uint64()
	b.FitTotalSupply = basicData.FitTotalSupply.String()
	b.FigTotalSupply = basicData.FigTotalSupply.String()
	b.TotalFIL = basicData.FilLiquidInfo.TotalFIL.String()
	b.AvailableFIL = basicData.FilLiquidInfo.AvailableFIL.String()
	b.UtilizedLiquidity = basicData.FilLiquidInfo.UtilizedLiquidity.String()
	b.AccumulatedDeposit = basicData.FilLiquidInfo.AccumulatedDeposit.String()
	b.AccumulatedRedeem = basicData.FilLiquidInfo.AccumulatedRedeem.String()
	b.AccumulatedBurntFILTrust = basicData.FilLiquidInfo.AccumulatedBurntFILTrust.String()
	b.AccumulatedMintFILTrust = basicData.FilLiquidInfo.AccumulatedMintFILTrust.String()
	b.AccumulatedBorrow = basicData.FilLiquidInfo.AccumulatedBorrow.String()
	b.AccumulatedPayback = basicData.FilLiquidInfo.AccumulatedPayback.String()
	b.AccumulatedInterest = basicData.FilLiquidInfo.AccumulatedInterest.String()
	b.AccumulatedRedeemFee = basicData.FilLiquidInfo.AccumulatedRedeemFee.String()
	b.AccumulatedBorrowFee = basicData.FilLiquidInfo.AccumulatedBorrowFee.String()
	b.AccumulatedLiquidateReward = basicData.FilLiquidInfo.AccumulatedLiquidateReward.String()
	b.AccumulatedLiquidateFee = basicData.FilLiquidInfo.AccumulatedLiquidateFee.String()
	b.AccumulatedDeposits = basicData.FilLiquidInfo.AccumulatedDeposits.String()
	b.AccumulatedBorrows = basicData.FilLiquidInfo.AccumulatedBorrows.String()
	b.UtilizationRate = basicData.FilLiquidInfo.UtilizationRate.String()
	b.ExchangeRate = basicData.FilLiquidInfo.ExchangeRate.String()
	b.InterestRate = basicData.FilLiquidInfo.InterestRate.String()
	b.CollateralizedMiner = basicData.FilLiquidInfo.CollateralizedMiner.String()
	b.MinerWithBorrows = basicData.FilLiquidInfo.MinerWithBorrows.String()
	b.RateBase = basicData.FilLiquidInfo.RateBase.String()
	b.AccumulatedStake = basicData.FilStakeInfo.AccumulatedStake.String()
	b.AccumulatedStakeDuration = basicData.FilStakeInfo.AccumulatedStakeDuration.String()
	b.AccumulatedInterestMint = basicData.FilStakeInfo.AccumulatedInterestMint.String()
	b.AccumulatedStakeMint = basicData.FilStakeInfo.AccumulatedStakeMint.String()
	b.AccumulatedWithdrawn = basicData.FilStakeInfo.AccumulatedWithdrawn.String()
	b.NextStakeID = basicData.FilStakeInfo.NextStakeID.String()
	b.ReleasedFigStake = basicData.FilStakeInfo.ReleasedFigStake.String()
	b.Bonders = basicData.GovernanceInfo.Bonders.String()
	b.TotalBondedAmount = basicData.GovernanceInfo.TotalBondedAmount.String()
	b.FirstActiveProposalId = basicData.GovernanceInfo.FirstActiveProposalId.String()
	b.U1 = basicData.FiLLiquidGovernanceFactors.U1.String()
	b.R0 = basicData.FiLLiquidGovernanceFactors.R0.String()
	b.R1 = basicData.FiLLiquidGovernanceFactors.R1.String()
	b.RM = basicData.FiLLiquidGovernanceFactors.RM.String()
	b.CollateralRate = basicData.FiLLiquidGovernanceFactors.CollateralRate.String()
	b.MaxFamilySize = basicData.FiLLiquidGovernanceFactors.MaxFamilySize.String()
	b.AlertThreshold = basicData.FiLLiquidGovernanceFactors.AlertThreshold.String()
	b.LiquidateThreshold = basicData.FiLLiquidGovernanceFactors.LiquidateThreshold.String()
	b.MaxLiquidations = basicData.FiLLiquidGovernanceFactors.MaxLiquidations.String()
	b.MinLiquidateInterval = basicData.FiLLiquidGovernanceFactors.MinLiquidateInterval.String()
	b.LiquidateDiscountRate = basicData.FiLLiquidGovernanceFactors.LiquidateDiscountRate.String()
	b.LiquidateFeeRate = basicData.FiLLiquidGovernanceFactors.LiquidateFeeRate.String()
	b.MaxExistingBorrows = basicData.FiLLiquidGovernanceFactors.MaxExistingBorrows.String()
	b.MinBorrowAmount = basicData.FiLLiquidGovernanceFactors.MinBorrowAmount.String()
	b.MinDepositAmount = basicData.FiLLiquidGovernanceFactors.MinDepositAmount.String()
	b.N = basicData.FiLLiquidGovernanceFactors.N.String()
	b.NInterest = basicData.FiLStakeGovernanceFactors.NInterest.String()
	b.NStake = basicData.FiLStakeGovernanceFactors.NStake.String()

	return b
}

func (b *BasicData) Down() *model.BasicDataStruct {
	d := new(model.BasicDataStruct)

	d.BlockHeight = fmt.Sprintf("%d", b.BlockHeight)
	d.BlockTimeStamp = fmt.Sprintf("%d", b.BlockTimeStamp)
	d.FitTotalSupply = b.FitTotalSupply
	d.FigTotalSupply = b.FigTotalSupply
	d.TotalFIL = b.TotalFIL
	d.AvailableFIL = b.AvailableFIL
	d.UtilizedLiquidity = b.UtilizedLiquidity
	d.AccumulatedDeposit = b.AccumulatedDeposit
	d.AccumulatedRedeem = b.AccumulatedRedeem
	d.AccumulatedBurntFILTrust = b.AccumulatedBurntFILTrust
	d.AccumulatedMintFILTrust = b.AccumulatedMintFILTrust
	d.AccumulatedBorrow = b.AccumulatedBorrow
	d.AccumulatedPayback = b.AccumulatedPayback
	d.AccumulatedInterest = b.AccumulatedInterest
	d.AccumulatedRedeemFee = b.AccumulatedRedeemFee
	d.AccumulatedBorrowFee = b.AccumulatedBorrowFee
	d.AccumulatedLiquidateReward = b.AccumulatedLiquidateReward
	d.AccumulatedLiquidateFee = b.AccumulatedLiquidateFee
	d.AccumulatedDeposits = b.AccumulatedDeposits
	d.AccumulatedBorrows = b.AccumulatedBorrows
	d.UtilizationRate = b.UtilizationRate
	d.ExchangeRate = b.ExchangeRate
	d.InterestRate = b.InterestRate
	d.CollateralizedMiner = b.CollateralizedMiner
	d.MinerWithBorrows = b.MinerWithBorrows
	d.RateBase = b.RateBase
	d.AccumulatedStake = b.AccumulatedStake
	d.AccumulatedStakeDuration = b.AccumulatedStakeDuration
	d.AccumulatedInterestMint = b.AccumulatedInterestMint
	d.AccumulatedStakeMint = b.AccumulatedStakeMint
	d.AccumulatedWithdrawn = b.AccumulatedWithdrawn
	d.NextStakeID = b.NextStakeID
	d.ReleasedFigStake = b.ReleasedFigStake
	d.Bonders = b.Bonders
	d.TotalBondedAmount = b.TotalBondedAmount
	d.FirstActiveProposalId = b.FirstActiveProposalId
	d.U1 = b.U1
	d.R0 = b.R0
	d.R1 = b.R1
	d.RM = b.RM
	d.CollateralRate = b.CollateralRate
	d.MaxFamilySize = b.MaxFamilySize
	d.AlertThreshold = b.AlertThreshold
	d.LiquidateThreshold = b.LiquidateThreshold
	d.MaxLiquidations = b.MaxLiquidations
	d.MinLiquidateInterval = b.MinLiquidateInterval
	d.LiquidateDiscountRate = b.LiquidateDiscountRate
	d.LiquidateFeeRate = b.LiquidateFeeRate
	d.MaxExistingBorrows = b.MaxExistingBorrows
	d.MinBorrowAmount = b.MinBorrowAmount
	d.MinDepositAmount = b.MinDepositAmount
	d.N = b.N
	d.NInterest = b.NInterest
	d.NStake = b.NStake

	return d
}

func (s *Dao) InsertBasic(data *BasicData) error {
	return s.db.Model(&BasicData{}).Create(data).Error
}

func (s *Dao) GetLatestBlockHeight() (uint64, error) {
	var height uint64
	if err := s.db.Model(&BasicData{}).Order("block_height desc").Select("block_height").Limit(1).Scan(&height).Error; err != nil {
		return 0, err
	}
	return height, nil
}

func (s *Dao) GetBasicDataCount() (int64, error) {
	var num int64
	if err := s.db.Model(&BasicData{}).Count(&num).Error; err != nil {
		return 0, err
	}
	return num, nil
}

func (s *Dao) GetLatestBasicData() (*BasicData, error) {
	data := new(BasicData)
	if err := s.db.Model(&BasicData{}).Order("block_height desc").Limit(1).Scan(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Dao) GetBasicDataAll() ([]BasicData, error) {
	var list []BasicData
	if err := s.db.Model(&BasicData{}).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Dao) GetBasicDataInterval(interval int64) ([]BasicData, error) {
	var list []BasicData

	if err := s.db.Find(&list).Where("block_height % ? = 0", interval).Order("block_height asc").Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (s *Dao) GetBasicDataLowerTimeCount(lowerBond int64) (int64, error) {
	var num int64
	if err := s.db.Model(&BasicData{}).Select("count(*)").Where("block_time_stamp >= ?", lowerBond).Count(&num).Error; err != nil {
		return 0, err
	}
	return num, nil
}

func (s *Dao) GetBasicDataLowerTimeList(lowerBond int64) ([]BasicData, error) {
	var list []BasicData
	if err := s.db.Model(&BasicData{}).Select("*").Where("block_time_stamp >= ?", lowerBond).Order("block_height asc").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Dao) GetBasicDataLowerTimeIntervalList(lowerBond int64, interval int64) ([]BasicData, error) {
	var list []BasicData
	if err := s.db.Model(&BasicData{}).Select("*").Where("block_time_stamp >= ? && block_height % ? = 0", lowerBond, interval).Order("block_height asc").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
