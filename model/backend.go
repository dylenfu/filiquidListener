package model

type BasicDataStruct struct {
	BlockHeight                string
	BlockTimeStamp             string
	FitTotalSupply             string
	FigTotalSupply             string
	TotalFIL                   string
	AvailableFIL               string
	UtilizedLiquidity          string
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
	AccumulatedStake           string
	AccumulatedStakeDuration   string
	AccumulatedInterestMint    string
	AccumulatedStakeMint       string
	AccumulatedWithdrawn       string
	NextStakeID                string
	ReleasedFigStake           string
	Bonders                    string
	TotalBondedAmount          string
	FirstActiveProposalId      string
	U1                         string
	R0                         string
	R1                         string
	RM                         string
	CollateralRate             string
	MaxFamilySize              string
	AlertThreshold             string
	LiquidateThreshold         string
	MaxLiquidations            string
	MinLiquidateInterval       string
	LiquidateDiscountRate      string
	LiquidateFeeRate           string
	MaxExistingBorrows         string
	MinBorrowAmount            string
	MinDepositAmount           string
	N                          string
	NInterest                  string
	NStake                     string
}

type SeniorDataStruct struct {
	BlockHeight          string
	BlockTimeStamp       string
	TotalPendingInterest string
	TotalFIL             string
	Borrowing            string
	BorrowingAndPeriod   string
	AccumulatedPayback   string
	InterestExp          string
}

type BasicDataStructFront struct {
	BlockTimeStamp          uint64
	UtilizationRate         float64
	FIL_FIT                 float64
	InterestRate            float64
	AccumulatedInterestMint string
	AccumulatedStakeMint    string
}

type SeniorDataStructFront struct {
	BlockTimeStamp uint64
	APY            float64
}

type PanelStructFront struct {
	TotalFIL                string
	FitTotalSupply          string
	UtilizationRate         float64
	FIL_FIT                 float64
	AvailableFIL            string
	AvailableFILRedeem      string
	UtilizedLiquidity       string
	InterestRate            float64
	FigTotalSupply          string
	FigReleased             string
	AccumulatedStakeMint    string
	AccumulatedInterestMint string
	APY                     float64
}

type DataStructFront struct {
	Basic        []BasicDataStructFront
	Basic1Day    []BasicDataStructFront
	Basic7Day    []BasicDataStructFront
	Basic1Month  []BasicDataStructFront
	Basic3Month  []BasicDataStructFront
	Senior       []SeniorDataStructFront
	Senior1Day   []SeniorDataStructFront
	Senior7Day   []SeniorDataStructFront
	Senior1Month []SeniorDataStructFront
	Senior3Month []SeniorDataStructFront
	Panel        PanelStructFront
}

type BasicDataStructFrontReal struct {
	Basic       Tuple
	Basic1Day   Tuple
	Basic7Day   Tuple
	Basic1Month Tuple
	Basic3Month Tuple
}

type SeniorDataStructFrontReal struct {
	Senior       Tuple
	Senior1Day   Tuple
	Senior7Day   Tuple
	Senior1Month Tuple
	Senior3Month Tuple
}

type PanelDataStructFrontReal struct {
	Panel PanelStructFront
}

type Tuple struct {
	B []string
	V []string
}
