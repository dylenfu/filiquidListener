package cache

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/model"
	"github.com/filiquid/listener/utils"
)

// todo: setheight
type CacheData struct {
	Basic    []byte
	Senior   []byte
	Panel    []byte
	Families []byte

	MaxCacheSize  int
	CurrentHeight uint64
	LastHeight    uint64

	FamilyAddr   []byte
	FamilyLength int
	StakerNum    int

	HaveSeen           map[string]bool
	Interest           map[string]model.InterestDataTuple
	Staked             map[string]model.StakedDataTuple
	UnStaked           map[string]model.UnstakedDataTuple
	WithdrawFig        map[string]model.WithdrawnFigDataTuple
	ProposedData       map[string]model.ProposedDataTuple
	ProposedByProposer map[string]model.ProposedDataTuple
	Voted              map[string]model.VotedDataTuple
	VotedByVoter       map[string]model.VotedDataTuple

	dao *dao.Dao
}

func NewCacheData(dao *dao.Dao, maxCacheSize int) *CacheData {
	return &CacheData{
		dao:          dao,
		MaxCacheSize: maxCacheSize,
		HaveSeen:     make(map[string]bool),
	}
}

func (c *CacheData) Close() {
}

func (c *CacheData) SetCurrentHeight(height uint64) {
	atomic.StoreUint64(&c.CurrentHeight, height)
}
func (c *CacheData) IncreaseCurrentHeight() {
	atomic.AddUint64(&c.CurrentHeight, 1)
}
func (c *CacheData) GetCurrentHeight() uint64 {
	return c.CurrentHeight
}

func (c *CacheData) SetLastHeight(height uint64) {
	atomic.StoreUint64(&c.LastHeight, height)
}
func (c *CacheData) GetLastHeight() uint64 {
	return c.LastHeight
}

func (c *CacheData) FetchAndSaveBasicCache() error {
	familyAddrs, err := c.dao.GetUserFamilyAddress()
	if err != nil {
		return err
	}
	jsFamilyAddrs, err := utils.ToJson(familyAddrs)
	if err != nil {
		return err
	}

	stakerNum, err := c.dao.GetDistinctStakerNum()
	if err != nil {
		return err
	}

	c.FamilyAddr = jsFamilyAddrs
	c.FamilyLength = len(familyAddrs)
	c.StakerNum = stakerNum
	return nil
}
func (c *CacheData) GetFamilyList() []byte {
	return c.FamilyAddr
}
func (c *CacheData) GetFamilyCount() int {
	return c.FamilyLength
}
func (c *CacheData) GetStakerNum() int {
	return c.StakerNum
}

func (c *CacheData) GetInterest(minter string) (model.InterestDataTuple, bool) {
	data, ok := c.Interest[minter]
	return data, ok
}
func (c *CacheData) SetInterest(now time.Time, minter string, list []model.InterestData) {
	if c.Interest == nil || len(c.Interest) >= c.MaxCacheSize {
		c.Interest = make(map[string]model.InterestDataTuple, 0)
	}
	c.Interest[minter] = model.InterestDataTuple{
		FetchedTime: now,
		Data:        list,
	}
}

func (c *CacheData) GetHaveSeen(data *model.MappingIndex) bool {
	key := mappingIndexKey(data.BlockNumber, data.Index)
	if data, ok := c.HaveSeen[key]; ok {
		return data
	} else {
		return false
	}
}
func (c *CacheData) SetHaveSeen(data *model.MappingIndex) {
	key := mappingIndexKey(data.BlockNumber, data.Index)
	c.HaveSeen[key] = true
}

func (c *CacheData) GetStake(staker string) (model.StakedDataTuple, bool) {
	data, ok := c.Staked[staker]
	return data, ok
}
func (c *CacheData) SetStake(now time.Time, staker string, list []model.StakedData) {
	if c.Staked == nil || len(c.Staked) >= c.MaxCacheSize {
		c.Staked = make(map[string]model.StakedDataTuple)
	}
	c.Staked[staker] = model.StakedDataTuple{
		FetchedTime: now,
		Data:        list,
	}
}

func (c *CacheData) GetUnStake(staker string) (model.UnstakedDataTuple, bool) {
	data, ok := c.UnStaked[staker]
	return data, ok
}
func (c *CacheData) SetUnStake(now time.Time, staker string, list []model.UnstakedData) {
	if c.UnStaked == nil || len(c.UnStaked) >= c.MaxCacheSize {
		c.UnStaked = make(map[string]model.UnstakedDataTuple)
	}
	c.UnStaked[staker] = model.UnstakedDataTuple{
		FetchedTime: now,
		Data:        list,
	}
}

func (c *CacheData) GetWithdrawFig(staker string) (model.WithdrawnFigDataTuple, bool) {
	data, ok := c.WithdrawFig[staker]
	return data, ok
}
func (c *CacheData) SetWithdrawFig(now time.Time, staker string, list []model.WithdrawnFigData) {
	if c.WithdrawFig == nil || len(c.WithdrawFig) >= c.MaxCacheSize {
		c.WithdrawFig = make(map[string]model.WithdrawnFigDataTuple)
	}
	c.WithdrawFig[staker] = model.WithdrawnFigDataTuple{
		FetchedTime: now,
		Data:        list,
	}
}

func (c *CacheData) GetProposedByProposer(proposer, mode string) (model.ProposedDataTuple, bool) {
	key := proposerModeKey(proposer, mode)
	data, ok := c.ProposedByProposer[key]
	return data, ok
}
func (c *CacheData) SetProposedByProposer(now time.Time, proposer, mode string, list []uint64) {
	if c.ProposedByProposer == nil || len(c.ProposedByProposer) >= c.MaxCacheSize {
		c.ProposedByProposer = make(map[string]model.ProposedDataTuple)
	}
	key := proposerModeKey(proposer, mode)
	c.ProposedByProposer[key] = model.ProposedDataTuple{
		FetchedTime: now,
		Data:        list,
	}
}

func (c *CacheData) GetProposed(mode string) (model.ProposedDataTuple, bool) {
	data, ok := c.ProposedData[mode]
	return data, ok
}
func (c *CacheData) SetProposed(now time.Time, mode string, list []uint64) {
	if c.ProposedData == nil || len(c.ProposedData) >= c.MaxCacheSize {
		c.ProposedData = make(map[string]model.ProposedDataTuple)
	}
	c.ProposedData[mode] = model.ProposedDataTuple{
		FetchedTime: now,
		Data:        list,
	}
}

func (c *CacheData) GetVotedByVoter(voter, proposalId string) (model.VotedDataTuple, bool) {
	key := proposerModeKey(voter, proposalId)
	data, ok := c.VotedByVoter[key]
	return data, ok
}
func (c *CacheData) SetVotedByVoter(now time.Time, voter, proposalId string, list []model.VotedData) {
	if c.VotedByVoter == nil || len(c.VotedByVoter) >= c.MaxCacheSize {
		c.VotedByVoter = make(map[string]model.VotedDataTuple)
	}
	key := proposerModeKey(voter, proposalId)
	c.VotedByVoter[key] = model.VotedDataTuple{
		FetchedTime: now,
		Data:        list,
	}
}

func (c *CacheData) GetVoted(proposalId string) (model.VotedDataTuple, bool) {
	data, ok := c.Voted[proposalId]
	return data, ok
}
func (c *CacheData) SetVoted(now time.Time, proposalId string, list []model.VotedData) {
	if c.Voted == nil || len(c.Voted) >= c.MaxCacheSize {
		c.Voted = make(map[string]model.VotedDataTuple)
	}
	c.Voted[proposalId] = model.VotedDataTuple{
		FetchedTime: now,
		Data:        list,
	}
}

func proposerModeKey(proposer, mode string) string {
	hash := utils.Keccak256Hash(proposer, mode)
	return hash.Hex()
}
func mappingIndexKey(block uint64, index uint) string {
	hash := utils.Keccak256Hash(fmt.Sprintf("%d", block), fmt.Sprintf("%d", index))
	return hash.Hex()
}
