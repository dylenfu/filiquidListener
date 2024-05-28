package dao

import (
	"fmt"
	"math/big"

	"github.com/filiquid/listener/eth/abi/data"
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
	//N                     string
	NInterest string
	NStake    string
}

func (b *BasicData) Up(basicData *struct {
	BlockHeight                *big.Int
	BlockTimeStamp             *big.Int
	FitTotalSupply             *big.Int
	FigTotalSupply             *big.Int
	FilLiquidInfo              data.FILLiquidInterfaceFILLiquidInfo
	FitStakeInfo               data.FITStakeFITStakeInfo
	GovernanceInfo             data.GovernanceGovernanceInfo
	FiLLiquidGovernanceFactors data.DataFetcherFiLLiquidGovernanceFactors
	FitStakeGovernanceFactors  data.DataFetcherFITStakeGovernanceFactors
}) *BasicData {
	if basicData.BlockHeight != nil {
		b.BlockHeight = basicData.BlockHeight.Uint64()
	}
	if basicData.BlockTimeStamp != nil {
		b.BlockTimeStamp = basicData.BlockTimeStamp.Uint64()
	}
	if basicData.FitTotalSupply != nil {
		b.FitTotalSupply = basicData.FitTotalSupply.String()
	}
	if basicData.FigTotalSupply != nil {
		b.FigTotalSupply = basicData.FigTotalSupply.String()
	}
	if basicData.FilLiquidInfo.TotalFIL != nil {
		b.TotalFIL = basicData.FilLiquidInfo.TotalFIL.String()
	}
	if basicData.FilLiquidInfo.AvailableFIL != nil {
		b.AvailableFIL = basicData.FilLiquidInfo.AvailableFIL.String()
	}
	if basicData.FilLiquidInfo.UtilizedLiquidity != nil {
		b.UtilizedLiquidity = basicData.FilLiquidInfo.UtilizedLiquidity.String()
	}
	if basicData.FilLiquidInfo.AccumulatedDeposit != nil {
		b.AccumulatedDeposit = basicData.FilLiquidInfo.AccumulatedDeposit.String()
	}
	if basicData.FilLiquidInfo.AccumulatedRedeem != nil {
		b.AccumulatedRedeem = basicData.FilLiquidInfo.AccumulatedRedeem.String()
	}
	if basicData.FilLiquidInfo.AccumulatedBurntFILTrust != nil {
		b.AccumulatedBurntFILTrust = basicData.FilLiquidInfo.AccumulatedBurntFILTrust.String()
	}
	if basicData.FilLiquidInfo.AccumulatedMintFILTrust != nil {
		b.AccumulatedMintFILTrust = basicData.FilLiquidInfo.AccumulatedMintFILTrust.String()
	}
	if basicData.FilLiquidInfo.AccumulatedBorrow != nil {
		b.AccumulatedBorrow = basicData.FilLiquidInfo.AccumulatedBorrow.String()
	}
	if basicData.FilLiquidInfo.AccumulatedPayback != nil {
		b.AccumulatedPayback = basicData.FilLiquidInfo.AccumulatedPayback.String()
	}
	if basicData.FilLiquidInfo.AccumulatedInterest != nil {
		b.AccumulatedInterest = basicData.FilLiquidInfo.AccumulatedInterest.String()
	}
	if basicData.FilLiquidInfo.AccumulatedRedeemFee != nil {
		b.AccumulatedRedeemFee = basicData.FilLiquidInfo.AccumulatedRedeemFee.String()
	}
	if basicData.FilLiquidInfo.AccumulatedBorrowFee != nil {
		b.AccumulatedBorrowFee = basicData.FilLiquidInfo.AccumulatedBorrowFee.String()
	}
	if basicData.FilLiquidInfo.AccumulatedLiquidateReward != nil {
		b.AccumulatedLiquidateReward = basicData.FilLiquidInfo.AccumulatedLiquidateReward.String()
	}
	if basicData.FilLiquidInfo.AccumulatedLiquidateFee != nil {
		b.AccumulatedLiquidateFee = basicData.FilLiquidInfo.AccumulatedLiquidateFee.String()
	}
	if basicData.FilLiquidInfo.AccumulatedDeposits != nil {
		b.AccumulatedDeposits = basicData.FilLiquidInfo.AccumulatedDeposits.String()
	}
	if basicData.FilLiquidInfo.AccumulatedBorrows != nil {
		b.AccumulatedBorrows = basicData.FilLiquidInfo.AccumulatedBorrows.String()
	}
	if basicData.FilLiquidInfo.UtilizationRate != nil {
		b.UtilizationRate = basicData.FilLiquidInfo.UtilizationRate.String()
	}
	if basicData.FilLiquidInfo.ExchangeRate != nil {
		b.ExchangeRate = basicData.FilLiquidInfo.ExchangeRate.String()
	}
	if basicData.FilLiquidInfo.InterestRate != nil {
		b.InterestRate = basicData.FilLiquidInfo.InterestRate.String()
	}
	if basicData.FilLiquidInfo.CollateralizedMiner != nil {
		b.CollateralizedMiner = basicData.FilLiquidInfo.CollateralizedMiner.String()
	}
	if basicData.FilLiquidInfo.MinerWithBorrows != nil {
		b.MinerWithBorrows = basicData.FilLiquidInfo.MinerWithBorrows.String()
	}
	if basicData.FilLiquidInfo.RateBase != nil {
		b.RateBase = basicData.FilLiquidInfo.RateBase.String()
	}
	if basicData.FitStakeInfo.AccumulatedStake != nil {
		b.AccumulatedStake = basicData.FitStakeInfo.AccumulatedStake.String()
	}
	if basicData.FitStakeInfo.AccumulatedStakeDuration != nil {
		b.AccumulatedStakeDuration = basicData.FitStakeInfo.AccumulatedStakeDuration.String()
	}
	if basicData.FitStakeInfo.AccumulatedInterestMint != nil {
		b.AccumulatedInterestMint = basicData.FitStakeInfo.AccumulatedInterestMint.String()
	}
	if basicData.FitStakeInfo.AccumulatedStakeMint != nil {
		b.AccumulatedStakeMint = basicData.FitStakeInfo.AccumulatedStakeMint.String()
	}
	if basicData.FitStakeInfo.AccumulatedWithdrawn != nil {
		b.AccumulatedWithdrawn = basicData.FitStakeInfo.AccumulatedWithdrawn.String()
	}
	if basicData.FitStakeInfo.NextStakeID != nil {
		b.NextStakeID = basicData.FitStakeInfo.NextStakeID.String()
	}
	if basicData.FitStakeInfo.ReleasedFIGStake != nil {
		b.ReleasedFigStake = basicData.FitStakeInfo.ReleasedFIGStake.String()
	}
	if basicData.GovernanceInfo.Bonders != nil {
		b.Bonders = basicData.GovernanceInfo.Bonders.String()
	}
	if basicData.GovernanceInfo.TotalBondedAmount != nil {
		b.TotalBondedAmount = basicData.GovernanceInfo.TotalBondedAmount.String()
	}
	if basicData.GovernanceInfo.FirstActiveProposalId != nil {
		b.FirstActiveProposalId = basicData.GovernanceInfo.FirstActiveProposalId.String()
	}
	if basicData.FiLLiquidGovernanceFactors.U1 != nil {
		b.U1 = basicData.FiLLiquidGovernanceFactors.U1.String()
	}
	if basicData.FiLLiquidGovernanceFactors.R0 != nil {
		b.R0 = basicData.FiLLiquidGovernanceFactors.R0.String()
	}
	if basicData.FiLLiquidGovernanceFactors.R1 != nil {
		b.R1 = basicData.FiLLiquidGovernanceFactors.R1.String()
	}
	if basicData.FiLLiquidGovernanceFactors.RM != nil {
		b.RM = basicData.FiLLiquidGovernanceFactors.RM.String()
	}
	if basicData.FiLLiquidGovernanceFactors.CollateralRate != nil {
		b.CollateralRate = basicData.FiLLiquidGovernanceFactors.CollateralRate.String()
	}
	if basicData.FiLLiquidGovernanceFactors.MaxFamilySize != nil {
		b.MaxFamilySize = basicData.FiLLiquidGovernanceFactors.MaxFamilySize.String()
	}
	if basicData.FiLLiquidGovernanceFactors.AlertThreshold != nil {
		b.AlertThreshold = basicData.FiLLiquidGovernanceFactors.AlertThreshold.String()
	}
	if basicData.FiLLiquidGovernanceFactors.LiquidateThreshold != nil {
		b.LiquidateThreshold = basicData.FiLLiquidGovernanceFactors.LiquidateThreshold.String()
	}
	if basicData.FiLLiquidGovernanceFactors.MaxLiquidations != nil {
		b.MaxLiquidations = basicData.FiLLiquidGovernanceFactors.MaxLiquidations.String()
	}
	if basicData.FiLLiquidGovernanceFactors.MinLiquidateInterval != nil {
		b.MinLiquidateInterval = basicData.FiLLiquidGovernanceFactors.MinLiquidateInterval.String()
	}
	if basicData.FiLLiquidGovernanceFactors.LiquidateDiscountRate != nil {
		b.LiquidateDiscountRate = basicData.FiLLiquidGovernanceFactors.LiquidateDiscountRate.String()
	}
	if basicData.FiLLiquidGovernanceFactors.LiquidateFeeRate != nil {
		b.LiquidateFeeRate = basicData.FiLLiquidGovernanceFactors.LiquidateFeeRate.String()
	}
	if basicData.FiLLiquidGovernanceFactors.MaxExistingBorrows != nil {
		b.MaxExistingBorrows = basicData.FiLLiquidGovernanceFactors.MaxExistingBorrows.String()
	}
	if basicData.FiLLiquidGovernanceFactors.MinBorrowAmount != nil {
		b.MinBorrowAmount = basicData.FiLLiquidGovernanceFactors.MinBorrowAmount.String()
	}
	if basicData.FiLLiquidGovernanceFactors.MinDepositAmount != nil {
		b.MinDepositAmount = basicData.FiLLiquidGovernanceFactors.MinDepositAmount.String()
	}

	//b.N = basicData.FiLLiquidGovernanceFactors.N.String()
	if basicData.FitStakeGovernanceFactors.NInterest != nil {
		b.NInterest = basicData.FitStakeGovernanceFactors.NInterest.String()
	}
	if basicData.FitStakeGovernanceFactors.NStake != nil {
		b.NStake = basicData.FitStakeGovernanceFactors.NStake.String()
	}

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
	//d.N = b.N
	d.NInterest = b.NInterest
	d.NStake = b.NStake

	return d
}

func (s *Dao) InsertBasic(data *BasicData) error {
	return s.db.Model(&BasicData{}).Create(data).Error
}

func (s *Dao) GetLatestBlockHeight() (uint64, error) {
	var height uint64
	if err := s.db.Model(&BasicData{}).Order("block_height desc").Select("block_height").First(&height).Error; err != nil {
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

	if err := s.db.Model(&BasicData{}).Order("block_height desc").First(data).Error; err != nil {
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
	if err := s.db.Model(&BasicData{}).Select("*").Where("block_time_stamp >= ?", lowerBond).Order("block_height asc").Find(&list).Error; err != nil {
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
