package eth

import (
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
	"github.com/filiquid/listener/eth/abi/fetcher"
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

type EClient struct {
	cfg         *config.EClient
	forceHeight uint64

	client *ethclient.Client
	cache  *cache.CacheData
	dao    *dao.Dao
	quit   chan bool

	liquidABI *abi.ABI
	fitABI    *abi.ABI
	govABI    *abi.ABI

	liquidFilter *liquid.LiquidFilterer
	fitFilter    *fitstake.FitstakeFilterer
	govFilter    *governance.GovernanceFilterer

	liquidHandler     eventHandler
	governanceHandler eventHandler
	fitstakeHandler   eventHandler

	fetcher *fetcher.DataCaller
	ticker  int
}

func NewETHClient(cache *cache.CacheData, dao *dao.Dao, cfg *config.EClient) (*EClient, error) {
	client := &EClient{}
	c, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		return nil, err
	}

	client.cfg = cfg
	client.client = c
	client.dao = dao
	client.cache = cache
	client.quit = make(chan bool, 1)

	client.setFilter()
	client.setABI()
	client.setHandler()
	return client, nil
}

func (s *EClient) setFilter() (err error) {
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

func (s *EClient) setABI() (err error) {
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

func (s *EClient) setHandler() {
	s.liquidHandler = s.handleLiquidEvent
	s.governanceHandler = s.handleGovernanceEvent
	s.fitstakeHandler = s.handleFitStakeEvent
}

func (s *EClient) setFetcher() {
	fetcher, err := fetcher.NewDataCaller(s.cfg.FetcherAddr, s.client)
	if err != nil {
		log.Fatal(err)
	}
	s.fetcher = fetcher
}

func (s *EClient) SetForceHeight(height uint64) {
	s.forceHeight = height
}

func (c *EClient) Close() {
	close(c.quit)
	c.client.Close()
}

// todo: set last height before running
// todo: running then sleep
func (s *EClient) FetchandSaveDataLoop() {
	ticker := time.NewTicker(time.Second * time.Duration(config.Conf.FetchDuration))
	for {
		select {
		case <-ticker.C:
			s.fetchAndSaveData()
		case <-s.quit:
			return
		}
	}
}

func (s *EClient) fetchAndSaveData() {
	current, start, end, err := s.getScanRange()
	if err != nil {
		log.Println(err)
	}
	log.Printf("start: %d, end: %d\n", start, end)
	if end-start < uint64(1) {
		return
	}

	s.cache.SetCurrentHeight(current)

	s.handleSingleContrctEvents(contractLiquid, start, end)
	s.handleSingleContrctEvents(contractGovernance, start, end)
	s.handleSingleContrctEvents(contractFitstake, start, end)

	if err := s.dao.UpdateLastHeight(end); err != nil {
		log.Printf("UpdateLastHeight failed, err: %v", err)
	} else {
		s.cache.SetLastHeight(end)
		log.Printf("UpdateLastHeight %d", end)
	}

	if err := s.cache.FetchAndSaveBasicCache(); err != nil {
		log.Printf("FetchAndSaveBasicCache failed, err: %v", err)
	}
}

func (s *EClient) handleSingleContrctEvents(name string, start, end uint64) {
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

func (s *EClient) getHandler(name string) (contract common.Address, abi *abi.ABI, handler eventHandler) {
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
