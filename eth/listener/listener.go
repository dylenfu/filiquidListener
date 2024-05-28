package listener

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/eth/abi/fitstake"
	"github.com/filiquid/listener/eth/abi/governance"
	"github.com/filiquid/listener/eth/abi/liquid"
	"github.com/filiquid/listener/utils"
)

const (
	eventDeposit                = "Deposit"
	eventRedeem                 = "Redeem"
	eventCollateralizingMiner   = "CollateralizingMiner"
	eventUncollateralizingMiner = "UncollateralizingMiner"
	eventBorrow                 = "Borrow"
	eventPayback                = "Payback"
	eventLiquidate              = "Liquidate"
	eventBorrowUpdated          = "BorrowUpdated"
	eventBorrowDropped          = "BorrowDropped"
	eventInterest               = "Interest"
	eventStaked                 = "Staked"
	eventUnstaked               = "Unstaked"
	eventWithdrawnFig           = "WithdrawnFig"
	eventWithdrawnFig2          = "WithdrawnFIG"
	eventProposed               = "Proposed"
	eventBonded                 = "Bonded"
	eventUnbonded               = "Unbonded"
	eventVoted                  = "Voted"
	eventExecuted               = "Executed"
)

const (
	contractLiquid     = "liquid"
	contractGovernance = "governance"
	contractFitstake   = "fitstake"
)

var errEvent = errors.New("event type error")

type eventHandler func(string, types.Log) (any, error)

type Listener struct {
	cfg *config.EClient

	client *ethclient.Client
	cache  *cache.CacheData
	dao    *dao.Dao

	liquidABI *abi.ABI
	fitABI    *abi.ABI
	govABI    *abi.ABI

	liquidFilter *liquid.LiquidFilterer
	fitFilter    *fitstake.FitstakeFilterer
	govFilter    *governance.GovernanceFilterer

	liquidHandler     eventHandler
	governanceHandler eventHandler
	fitstakeHandler   eventHandler

	forceHeight uint64
	quit        chan bool
}

func NewListener(dao *dao.Dao, cache *cache.CacheData, cfg *config.EClient) *Listener {

	client, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		log.Fatal(err)
	}

	listener := &Listener{}
	listener.cfg = cfg
	listener.client = client
	listener.cache = cache
	listener.dao = dao

	listener.setFilter()
	listener.setABI()
	listener.setHandler()

	listener.quit = make(chan bool, 1)
	return listener
}

func (s *Listener) setFilter() (err error) {
	if s.liquidFilter, err = liquid.NewLiquidFilterer(s.cfg.FiliquidAddr, s.client); err != nil {
		return err
	}
	if s.fitFilter, err = fitstake.NewFitstakeFilterer(s.cfg.FitStakeAddr, s.client); err != nil {
		return err
	}
	if s.govFilter, err = governance.NewGovernanceFilterer(s.cfg.GovernanceAddr, s.client); err != nil {
		return err
	}
	return nil
}

func (s *Listener) setABI() (err error) {
	if s.liquidABI, err = liquid.LiquidMetaData.GetAbi(); err != nil {
		return
	}
	if s.fitABI, err = fitstake.FitstakeMetaData.GetAbi(); err != nil {
		return
	}
	if s.govABI, err = governance.GovernanceMetaData.GetAbi(); err != nil {
		return
	}
	return
}

func (s *Listener) setHandler() {
	s.liquidHandler = s.handleLiquidEvent
	s.governanceHandler = s.handleGovernanceEvent
	s.fitstakeHandler = s.handleFitStakeEvent
}

func (s *Listener) IterateOnChainEvents(forceHeight uint64) {
	if forceHeight > 0 {
		s.forceHeight = forceHeight
	}

	lastHeight, err := s.dao.GetLastHeight()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Last height: %v\r\n", lastHeight)

	currentHeight, err := s.GetCurrentHeight()
	if err != nil {
		log.Fatalf("Get current height failed, err: %v", err)
	}
	log.Printf("current height %v", currentHeight)

	s.cache.SetLastHeight(lastHeight)
	s.cache.SetCurrentHeight(currentHeight)

	if err := s.cache.FetchAndSaveBasicCache(); err != nil {
		log.Fatalf("fetchAndSaveBasicCache failed, err:%v", err)
	}

	ticker := time.NewTicker(time.Second * time.Duration(config.Conf.FetchDuration))
	for {
		select {
		case <-ticker.C:
			s.iterateBlockEvents()
		case <-s.quit:
			return
		}
	}
}

func (s *Listener) Close() {
	close(s.quit)
}

func (s *Listener) iterateBlockEvents() {
	var start, end uint64
	current, err := s.client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("client.BlockNumber failed, err: %v", err)
		return
	}
	last := s.cache.GetLastHeight()

	if last < current {
		start = last + 1
	} else {
		return
	}

	if s.forceHeight == 0 {
	} else if s.forceHeight <= start {
		log.Fatalf("forceHeight %d  < lower %d", s.forceHeight, start)
		return
	} else if s.forceHeight > current {
		log.Fatalf("forceHeight %d > current %d", s.forceHeight, current)
		return
	} else {
		start = s.forceHeight
		s.forceHeight = 0
	}

	end = current
	if start > end {
		log.Fatalf("lower %d > higher %d", start, end)
		return
	}
	if start < end-uint64(config.Conf.BlockScanRange) {
		end = start + uint64(config.Conf.BlockScanRange)
	}

	log.Printf("Listener height, current: %v, start: %d, end: %d\n", current, start, end)

	s.handleSingleContrctEvents(contractLiquid, start, end)
	s.handleSingleContrctEvents(contractGovernance, start, end)
	s.handleSingleContrctEvents(contractFitstake, start, end)

	s.cache.SetCurrentHeight(current)
	if err := s.dao.UpdateLastHeight(end); err != nil {
		log.Printf("UpdateLastHeight failed, err: %v", err)
	}
	s.cache.SetLastHeight(end)
	log.Printf("Listener height %d", end)

	if err := s.cache.FetchAndSaveBasicCache(); err != nil {
		log.Printf("FetchAndSaveBasicCache failed, err: %v", err)
	}
}

func (s *Listener) handleSingleContrctEvents(name string, start, end uint64) {
	contract, abi, handler := s.getHandler(name)
	eventlogs := s.getLogs(start, end, contract)
	for _, vlog := range eventlogs {
		if !s.checkMappingIndex(&vlog) {
			continue
		}

		evt, err := abi.EventByID(vlog.Topics[0])
		if err != nil {
			log.Fatalf("Invalid event: %v", err)
		} else {
			log.Printf("Block: %d, Index: %d, Event: %s\r\n", vlog.BlockNumber, vlog.Index, evt.Name)
		}

		if data, err := handler(evt.Name, vlog); err != nil {
			log.Printf("Failed to handle %s event, err: %v", name, err)
			continue
		} else {
			log.Printf("Height %v, %s, event %s\r\n", vlog.BlockNumber, name, utils.ToString(data))
		}
	}
}

func (s *Listener) getHandler(name string) (contract common.Address, abi *abi.ABI, handler eventHandler) {
	switch name {
	case contractLiquid:
		contract = s.cfg.FiliquidAddr
		abi = s.liquidABI
		handler = s.liquidHandler
	case contractGovernance:
		contract = s.cfg.GovernanceAddr
		abi = s.govABI
		handler = s.governanceHandler
	case contractFitstake:
		contract = s.cfg.FitStakeAddr
		abi = s.fitABI
		handler = s.fitstakeHandler
	default:
		panic("getHandler error: type invalid")
	}
	return
}
