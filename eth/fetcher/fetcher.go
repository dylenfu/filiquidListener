package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
	caller "github.com/filiquid/listener/eth/abi/data"
	"github.com/filiquid/listener/utils"
)

type Fetcher struct {
	cfg        *config.EClient
	lastHeight uint64
	ticker     int

	client *ethclient.Client
	Caller *caller.DataCaller
	cache  *cache.CacheData
	dao    *dao.Dao

	quit chan bool
}

func NewFetcher(dao *dao.Dao, cache *cache.CacheData, cfg *config.EClient) *Fetcher {
	fetcher := &Fetcher{
		cfg:   cfg,
		cache: cache,
		dao:   dao,
	}
	client, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		log.Fatal(err)
	}
	caller, err := caller.NewDataCaller(cfg.FetcherAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	fetcher.client = client
	fetcher.Caller = caller
	fetcher.quit = make(chan bool, 1)
	return fetcher
}

func (s *Fetcher) IterateDataCallerQuerys(forceHeight uint64) {
	lastHeight, err := s.dao.GetLatestBlockHeight()
	if err != nil {
		log.Printf("fetch latest blcok heigh failed, err:%v", err)
	}

	// 初始运行时，以及后续需要回退一定距离时可以强行设置lastHeight
	if forceHeight > 0 && (lastHeight == 0 || (lastHeight > 0 && forceHeight < lastHeight)) {
		lastHeight = forceHeight
	}
	// cache global state info
	s.cache.FetchAndSaveBasicSeniorData()
	s.lastHeight = lastHeight

	ticker := time.NewTicker(time.Second * time.Duration(config.Conf.FetchDuration))
	for {
		select {
		case <-ticker.C:
			current, err := s.client.BlockNumber(context.Background())
			if err != nil {
				log.Printf("get current blockNumber failed,err: %v", err)
			}
			if s.lastHeight < current {
				s.FetchAndSaveFamilies(s.lastHeight)
				s.FetchAndSaveData(s.lastHeight)
				s.lastHeight += 1
			}
			log.Printf("Fetcher height, last %v, onchain %v", s.lastHeight, current)
		case <-s.quit:
			return
		}
	}
}

func (s *Fetcher) Close() {
	close(s.quit)
}

func (c *Fetcher) getFamilies(users []common.Address) error {
	info, err := c.Caller.GetBatchedUserBorrows(nil, users)
	if err != nil {
		return fmt.Errorf("getBatchedUserBorrows failed, err: %v", err)
	}
	bs, err := utils.ToJson(info)
	if err != nil {
		return fmt.Errorf("unmarshal families info failed, err: %v", err)
	}
	c.cache.SetFamilies(bs)
	return nil
}

func (c *Fetcher) FetchAndSaveFamilies(height uint64) error {
	raw := c.cache.GetFamilyList()
	if raw == nil {
		return fmt.Errorf("cache.GetFamilyList is nil")
	}

	var users []common.Address
	if err := json.Unmarshal(raw, &users); err != nil {
		return fmt.Errorf("unmarshl familyList failed, err:%v", err)
	}
	return c.getFamilies(users)
}

func (c *Fetcher) FetchAndSaveData(height uint64) {
	query := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(height)}
	rBasic, err := c.Caller.FetchData(query)
	if err != nil {
		log.Printf("fetchData failed, height %v, err: %v", height, err)
		return
	}

	basic := new(dao.BasicData).Up(&rBasic)
	if err := c.dao.InsertBasic(basic); err != nil {
		log.Printf("Insert basic data failed, err: %v", err)
		return
	}

	if c.ticker == 0 {
		response, err := c.Caller.GetTotalPendingInterest(query)
		if err != nil {
			log.Printf("Get total pending intersest failed, err: %v", err)
			return
		}

		result := new(dao.SeniorData).Up(response)
		if err := c.dao.InsertSenior(result); err != nil {
			log.Printf("Insert senior data failed, err: %v", err)
			return
		}
	}
	c.ticker = (c.ticker + 1) % config.FETCHDATAINTERVAL
	if err := c.cache.FetchAndSaveBasicSeniorData(); err != nil {
		log.Printf("insert senior data failed, err: %v", err)
		return
	}
}
