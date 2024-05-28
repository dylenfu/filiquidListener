package fetcher

import (
	"encoding/json"
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
	LastHeight uint64
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
	s.FetchAndSaveFamilies()
	s.cache.FetchAndSaveBasicSeniorData()

	ticker := time.NewTicker(time.Second * config.SECONDSBETWEENFETCH)
	for {
		select {
		case <-ticker.C:
			s.FetchAndSaveFamilies()
			s.FetchAndSaveData(lastHeight)
			lastHeight += 1
		case <-s.quit:
			return
		}
	}
}

func (s *Fetcher) Close() {
	close(s.quit)
}

func (c *Fetcher) FetchAndSaveFamilies() {
	raw := c.cache.GetFamilies()
	if raw == nil {
		return
	}
	var list []common.Address
	if err := json.Unmarshal(raw, &list); err != nil {
		log.Printf("getFamilies unmarshal data failed, raw: %s, err: %v", string(raw), err)
		return
	}

	info, err := c.Caller.GetBatchedUserBorrows(nil, list)
	if err != nil {
		log.Printf("getBatchedUserBorrows failed, err: %v", err)
		return
	}
	if info != nil {
		bs, err := utils.ToJson(info)
		if err != nil {
			log.Printf("marshal borrows info to json failed, err:%v", err)
			return
		}
		c.cache.SetFamilies(bs)
	}
}

func (c *Fetcher) FetchAndSaveData(height uint64) {
	rBasic, err := c.Caller.FetchData(&bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(height),
	})
	if err != nil {
		log.Printf("fetchData failed, height %v, err: %v", height, err)
		return
	}

	log.Printf("ETH Client Get height: %d", rBasic.BlockHeight.Uint64())
	basic := new(dao.BasicData).Up(&rBasic)
	if err := c.dao.InsertBasic(basic); err != nil {
		log.Printf("Insert basic data failed, err: %v", err)
		return
	}

	if c.ticker == 0 {
		response, err := c.Caller.GetTotalPendingInterest(nil)
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
