package utils

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/eth/abi/fitstake"
	"github.com/filiquid/listener/eth/abi/governance"
	"github.com/filiquid/listener/eth/abi/liquid"
	"github.com/filiquid/listener/model"
)

func FromDeposit(raw *liquid.LiquidDeposit) *model.Deposit {
	data := new(model.Deposit)
	data.Account = raw.Account
	data.AmountFIL = raw.AmountFIL
	data.AmountFILTrust = raw.AmountFILTrust
	return data
}

func FromRedeem(raw *liquid.LiquidRedeem) *model.Redeem {
	data := new(model.Redeem)
	data.Account = raw.Account
	data.AmountFILTrust = raw.AmountFILTrust
	data.AmountFIL = raw.AmountFIL
	data.Fee = raw.Fee
	return data
}

func FromCollateralizingMiner(raw *liquid.LiquidCollateralizingMiner) *model.CollateralizingMiner {
	data := new(model.CollateralizingMiner)
	data.MinerId = new(big.Int).SetUint64(raw.MinerId)
	data.Sender = raw.Sender
	data.Beneficiary = raw.Beneficiary
	data.Quota = raw.Quota
	data.Expiration = raw.Expiration
	return data
}

func FromUnCollateralizingMiner(raw *liquid.LiquidUncollateralizingMiner) *model.UncollateralizingMiner {
	data := new(model.UncollateralizingMiner)
	data.MinerId = new(big.Int).SetUint64(raw.MinerId)
	data.Sender = raw.Sender
	data.Beneficiary = raw.Beneficiary
	data.Quota = raw.Quota
	data.Expiration = raw.Expiration
	return data
}

func FromBorrow(raw *liquid.LiquidBorrow) *model.Borrow {
	data := new(model.Borrow)
	data.BorrowId = raw.BorrowId
	data.Account = raw.Account
	data.MinerId = new(big.Int).SetUint64(raw.MinerId)
	data.AmountFIL = raw.AmountFIL
	data.Fee = raw.Fee
	data.InterestRate = raw.InterestRate
	data.InitialTime = raw.InitialTime
	return data
}

func FromBorrowUpdated(raw *liquid.LiquidBorrowUpdated) *model.BorrowUpdated {
	data := new(model.BorrowUpdated)
	data.Id = raw.BorrowId
	data.BorrowAmount = raw.BorrowAmount
	data.RemainingOriginalAmount = raw.RemainingOriginalAmount
	data.DatedTime = raw.DatedTime
	return data
}

func FromBorrowDropped(raw *liquid.LiquidBorrowDropped) *model.BorrowDropped {
	data := new(model.BorrowDropped)
	data.Id = raw.Id
	return data
}

func FromPayback(raw *liquid.LiquidPayback) *model.Payback {
	data := new(model.Payback)
	data.Account = raw.Account
	data.MinerIdPayee = new(big.Int).SetUint64(raw.MinerIdPayee)
	data.MinerIdPayer = new(big.Int).SetUint64(raw.MinerIdPayer)
	data.Principal = raw.Principal
	data.Interest = raw.Interest
	data.Withdrawn = raw.Withdrawn
	data.MintedFIG = raw.MintedFIG
	return data
}

func FromLiquidate(raw *liquid.LiquidLiquidate) *model.Liquidate {
	data := new(model.Liquidate)
	data.Account = raw.Account
	data.MinerIdPayee = new(big.Int).SetUint64(raw.MinerIdPayee)
	data.MinerIdPayer = new(big.Int).SetUint64(raw.MinerIdPayer)
	data.Principal = raw.Principal
	data.Interest = raw.Interest
	data.Reward = raw.Reward
	data.Fee = raw.Fee
	data.MintedFIG = raw.MintedFIG
	return data
}

func FromProposed(raw *governance.GovernanceProposed, blockNum uint64) *model.Proposed {
	data := new(model.Proposed)
	data.Proposer = raw.Proposer
	data.ProposalId = raw.ProposalId
	data.DiscussionIndex = raw.DiscussionIndex
	data.Category = raw.Category
	data.Deadline = raw.Deadline
	data.Deposited = raw.Deposited
	data.Text = raw.Text
	data.Values = raw.Values
	data.Timestamp = HeightToUnix(blockNum)
	data.Executed = false
	data.Result = 0
	return data
}

func FromBonded(raw *governance.GovernanceBonded) *model.Bonded {
	data := new(model.Bonded)
	data.Amount = raw.Amount
	data.Bonder = raw.Bonder
	return data
}

func FromUnBonded(raw *governance.GovernanceUnbonded) *model.Bonded {
	data := new(model.Bonded)
	data.Amount = raw.Amount
	data.Bonder = raw.Bonder
	return data
}

func FromExcuted(raw *governance.GovernanceExecuted) *model.Executed {
	data := new(model.Executed)
	data.Executor = raw.Executor
	data.ProposalId = raw.ProposalId
	data.Result = raw.Result
	return data
}

func FromVoted(raw *governance.GovernanceVoted, blockNum uint64) *model.Voted {
	data := new(model.Voted)
	data.Voter = raw.Voter
	data.ProposalId = raw.ProposalId
	data.Category = raw.Category
	data.Amount = raw.Amount
	data.Timestamp = HeightToUnix(blockNum)
	return data
}

func FromInterest(raw *fitstake.FitstakeInterest, blockNum uint64) *model.Interest {
	data := new(model.Interest)
	data.Minter = raw.Minter
	data.Principal = raw.Principal
	data.Interest = raw.Interest
	data.Minted = raw.Minted
	data.Timestamp = HeightToUnix(blockNum)
	return data
}

func FromStaked(raw *fitstake.FitstakeStaked, blockNum uint64) *model.Staked {
	data := new(model.Staked)
	data.Staker = raw.Staker
	data.Id = raw.Id
	data.Amount = raw.Amount
	data.Start = raw.Start
	data.End = raw.End
	data.Minted = raw.Minted
	data.Timestamp = HeightToUnix(blockNum)
	return data
}

func FromUnstaked(raw *fitstake.FitstakeUnstaked, blockNum uint64) *model.Unstaked {
	data := new(model.Unstaked)
	data.Staker = raw.Staker
	data.Id = raw.Id
	data.Amount = raw.Amount
	data.Start = raw.Start
	data.End = raw.End
	data.RealEnd = raw.RealEnd
	data.Minted = raw.Minted
	data.Timestamp = HeightToUnix(blockNum)
	return data
}

func FromWithdrawFig(raw *fitstake.FitstakeWithdrawnFIG, blockNum uint64) *model.WithdrawnFig {
	data := new(model.WithdrawnFig)
	data.Staker = raw.Staker
	data.Id = raw.Id
	data.Amount = raw.Amount
	data.Timestamp = HeightToUnix(blockNum)
	return data
}

func DataStructFrontConvert(front *model.DataStructFront) (basic *model.BasicDataStructFrontReal, senior *model.SeniorDataStructFrontReal, panel *model.PanelDataStructFrontReal) {
	basic = new(model.BasicDataStructFrontReal)
	senior = new(model.SeniorDataStructFrontReal)
	panel = new(model.PanelDataStructFrontReal)
	basic.Basic = BasicDataStructFrontArray2Tuple(front.Basic)
	basic.Basic1Day = BasicDataStructFrontArray2Tuple(front.Basic1Day)
	basic.Basic7Day = BasicDataStructFrontArray2Tuple(front.Basic7Day)
	basic.Basic1Month = BasicDataStructFrontArray2Tuple(front.Basic1Month)
	basic.Basic3Month = BasicDataStructFrontArray2Tuple(front.Basic3Month)
	senior.Senior = SeniorDataStructFrontArray2Tuple(front.Senior)
	senior.Senior1Day = SeniorDataStructFrontArray2Tuple(front.Senior1Day)
	senior.Senior7Day = SeniorDataStructFrontArray2Tuple(front.Senior7Day)
	senior.Senior1Month = SeniorDataStructFrontArray2Tuple(front.Senior1Month)
	senior.Senior3Month = SeniorDataStructFrontArray2Tuple(front.Senior3Month)
	panel.Panel = front.Panel
	return
}

func BasicDataStructFrontArray2Tuple(a []model.BasicDataStructFront) model.Tuple {
	var v model.Tuple
	v.B = make([]string, len(a))
	v.V = make([]string, len(a))
	for i := 0; i < len(a); i++ {
		v.B[i] = strconv.Itoa(int(a[i].BlockTimeStamp))
		v.V[i] = strconv.FormatFloat(a[i].InterestRate, 'f', 6, 64)
	}
	return v
}

func SeniorDataStructFrontArray2Tuple(a []model.SeniorDataStructFront) model.Tuple {
	var v model.Tuple
	v.B = make([]string, len(a))
	v.V = make([]string, len(a))
	for i := 0; i < len(a); i++ {
		v.B[i] = strconv.Itoa(int(a[i].BlockTimeStamp))
		v.V[i] = strconv.FormatFloat(a[i].APY, 'f', 6, 64)
	}
	return v
}

func GetBasicDataStructFront(d *model.BasicDataStruct) (*model.BasicDataStructFront, error) {
	var (
		f = new(model.BasicDataStructFront)
	)

	rBlockTimeStamp, err := Str2Uint64(d.BlockTimeStamp)
	if err != nil {
		return nil, err
	}
	f.BlockTimeStamp = rBlockTimeStamp

	uRateBase, err := Str2Uint64(d.RateBase)
	if err != nil {
		return nil, err
	}
	rateBase := float64(uRateBase)

	rUtilizationRation, err := Str2Uint64(d.UtilizationRate)
	if err != nil {
		return nil, err
	}
	f.UtilizationRate = float64(rUtilizationRation) / rateBase

	rFILFIT, err := Str2Uint64(d.ExchangeRate)
	if err != nil {
		return nil, err
	}
	f.FIL_FIT = float64(rFILFIT) / rateBase

	if f.UtilizationRate >= config.U_M {
		uRM, err := Str2Uint64(d.RM)
		if err != nil {
			return nil, err
		}
		f.InterestRate = float64(uRM) / rateBase
	} else {
		uInterestRate, err := Str2Uint64(d.InterestRate)
		if err != nil {
			return nil, err
		}
		f.InterestRate = float64(uInterestRate) / rateBase
	}
	f.AccumulatedInterestMint = d.AccumulatedInterestMint
	f.AccumulatedStakeMint = d.AccumulatedStakeMint
	return f, nil
}

func GetSeniorDataStructFront3(d *model.SeniorDataStruct) (*model.SeniorDataStructFront, error) {
	f := new(model.SeniorDataStructFront)
	rBlockTimeStamp, err := Str2Uint64(d.BlockTimeStamp)
	if err != nil {
		return nil, err
	}
	f.BlockTimeStamp = rBlockTimeStamp
	r := big.NewInt(0)
	r, ok := r.SetString(d.InterestExp, 10)
	if !ok {
		return nil, fmt.Errorf("InterestExp value not legal")
	}
	r.Mul(r, big.NewInt(config.RATEBASE))
	s := big.NewInt(0)
	s, ok = s.SetString(d.TotalFIL, 10)
	if !ok {
		return nil, fmt.Errorf("TotalFIL value not legal")
	}
	r.Div(r, s)
	f.APY = float64(r.Uint64()) / config.RATEBASE
	return f, nil
}

func CalculateAPY(interestExp, totalFil string) (float64, error) {
	r := big.NewInt(0)
	r, ok := r.SetString(interestExp, 10)
	if !ok {
		return 0, fmt.Errorf("InterestExp value not legal")
	}
	r.Mul(r, big.NewInt(config.RATEBASE))
	s := big.NewInt(0)
	s, ok = s.SetString(totalFil, 10)
	if !ok {
		return 0, fmt.Errorf("TotalFIL value not legal")
	}
	r.Div(r, s)
	return float64(r.Uint64()) / config.RATEBASE, nil
}

func Str2Uint64(input string) (uint64, error) {
	r, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0, err
	}
	return r, nil
}
