package utils

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/eth/abi/fitstake"
	"github.com/filiquid/listener/eth/abi/governance"
	"github.com/filiquid/listener/eth/abi/liquid"
	"github.com/filiquid/listener/model"
	"golang.org/x/crypto/sha3"
)

func ToJson(data any) ([]byte, error) {
	return json.Marshal(data)
}

func ToString(data any) string {
	raw, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(raw)
}

func HeightToUnix(height uint64) uint64 {
	return (height * 30) + config.CALIBRATION_GENESIS_UNIX_EPOCH
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Keccak256Hash(s1, s2 string) common.Hash {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(s1))
	hasher.Write([]byte(s2))
	data := hasher.Sum(nil)
	return common.BytesToHash(data)
}

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
