package model

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Deposit struct {
	Account        common.Address
	AmountFIL      *big.Int
	AmountFILTrust *big.Int
}

type Redeem struct {
	Account        common.Address
	AmountFILTrust *big.Int
	AmountFIL      *big.Int
	Fee            *big.Int
}

type CollateralizingMiner struct {
	MinerId     *big.Int
	Sender      common.Address
	Beneficiary []byte
	Quota       *big.Int
	Expiration  int64
}

type UncollateralizingMiner struct {
	MinerId     *big.Int
	Sender      common.Address
	Beneficiary []byte
	Quota       *big.Int
	Expiration  int64
}

type Borrow struct {
	BorrowId     *big.Int
	Account      common.Address
	MinerId      *big.Int
	AmountFIL    *big.Int
	Fee          *big.Int
	InterestRate *big.Int
	InitialTime  *big.Int
}

type Payback struct {
	Account      common.Address
	MinerIdPayee *big.Int
	MinerIdPayer *big.Int
	Principal    *big.Int
	Interest     *big.Int
	Withdrawn    *big.Int
	MintedFIG    *big.Int
}

type Liquidate struct {
	Account      common.Address
	MinerIdPayee *big.Int
	MinerIdPayer *big.Int
	Principal    *big.Int
	Interest     *big.Int
	Reward       *big.Int
	Fee          *big.Int
	MintedFIG    *big.Int
}

type BorrowUpdated struct {
	Id                      *big.Int
	BorrowAmount            *big.Int
	RemainingOriginalAmount *big.Int
	DatedTime               *big.Int
}

type BorrowDropped struct {
	Id *big.Int
}

type Interest struct {
	Minter    common.Address
	Principal *big.Int
	Interest  *big.Int
	Minted    *big.Int
	Timestamp uint64
}

type InterestData struct {
	Amount    string
	Minted    string
	Timestamp string
}

type InterestDataRaw struct {
	Principal string
	Interest  string
	Minted    string
	Timestamp string
}

type Staked struct {
	Staker    common.Address
	Id        *big.Int
	Amount    *big.Int
	Start     *big.Int
	End       *big.Int
	Minted    *big.Int
	Timestamp uint64
}

type StakedData struct {
	Id        string
	Amount    string
	Start     string
	End       string
	Minted    string
	Timestamp string
}

type Unstaked struct {
	Staker    common.Address
	Id        *big.Int
	Amount    *big.Int
	Start     *big.Int
	End       *big.Int
	RealEnd   *big.Int
	Minted    *big.Int
	Timestamp uint64
}

type UnstakedData struct {
	Id        string
	Amount    string
	Start     string
	End       string
	RealEnd   string
	Minted    string
	Timestamp string
}

type WithdrawnFig struct {
	Staker    common.Address
	Id        *big.Int
	Amount    *big.Int
	Timestamp uint64
}

type WithdrawnFigData struct {
	Id        string
	Amount    string
	Timestamp string
}

type Proposed struct {
	Proposer        common.Address
	ProposalId      *big.Int
	DiscussionIndex *big.Int
	Category        uint8
	Deadline        *big.Int
	Deposited       *big.Int
	Text            string
	Values          []*big.Int
	Timestamp       uint64
	Executed        bool
	Result          uint8
}

func (p *Proposed) ValueString() string {
	if p.Values == nil || len(p.Values) == 0 {
		return "[]"
	}
	var vstr string = "["
	for i, v := range p.Values {
		vstr += v.String()
		if i != len(p.Values)-1 {
			vstr += ","
		}
	}
	vstr += "]"
	return vstr
}

type Bonded struct {
	Bonder common.Address
	Amount *big.Int
}

type Voted struct {
	Voter      common.Address
	ProposalId *big.Int
	Category   uint8
	Amount     *big.Int
	Timestamp  uint64
}

type VotedData struct {
	Voter    string
	Category string
	Amount   string
}

type Executed struct {
	Executor   common.Address
	ProposalId *big.Int
	Result     uint8
}

type MappingIndex struct {
	BlockNumber uint64
	Index       uint
}

func NewMappingIndex(block uint64, index uint) MappingIndex {
	return MappingIndex{
		BlockNumber: block,
		Index:       index,
	}
}

type InterestDataTuple struct {
	FetchedTime time.Time
	Data        []InterestData
}

type StakedDataTuple struct {
	FetchedTime time.Time
	Data        []StakedData
}

type UnstakedDataTuple struct {
	FetchedTime time.Time
	Data        []UnstakedData
}

type WithdrawnFigDataTuple struct {
	FetchedTime time.Time
	Data        []WithdrawnFigData
}

type ProposedDataTuple struct {
	FetchedTime time.Time
	Data        []uint64
}

type ProposedDataList struct {
	Ids         []uint64
	TotalLength int
}

type VotedDataTuple struct {
	FetchedTime time.Time
	Data        []VotedData
}

type VotedDataList struct {
	Votes       []VotedData
	TotalLength int
}

type ProposedDataKey struct {
	Proposer string
	Mode     string
}

type VotedDataKey struct {
	Voter      string
	ProposalId string
}
