package core

import (
	"log"

	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/eth"
	"github.com/filiquid/listener/rpc"
)

type Listener struct {
	cache *cache.CacheData
	dao   *dao.Dao
	eth   *eth.EClient
	rpc   *rpc.RPCServer
}

func NewListener() *Listener {
	dao, err := dao.NewDao(config.Conf.Db)
	if err != nil {
		log.Fatal(err)
	}

	cache := cache.NewCacheData(dao, config.Conf.MaxCacheSize)
	eclient, err := eth.NewETHClient(cache, dao, config.Conf.Eth)
	if err != nil {
		log.Fatal(err)
	}
	rpclient := rpc.NewRPCServer(config.Conf.Port, dao, cache)

	listener := &Listener{
		dao:   dao,
		cache: cache,
		eth:   eclient,
		rpc:   rpclient,
	}

	return listener
}

func (s *Listener) Migrate() {
	s.dao.Migrate()
}

func (s *Listener) Run(forceHeight uint64) {

	lastHeight, err := s.dao.GetLastHeight()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Last height: %v\r\n", lastHeight)

	currentHeight, err := s.eth.GetCurrentHeight()
	if err != nil {
		log.Fatalf("Get current height failed, err: %v", err)
	}
	log.Printf("current height %v", currentHeight)

	s.cache.SetLastHeight(lastHeight)
	s.cache.SetCurrentHeight(currentHeight)

	if err := s.cache.FetchAndSaveBasicCache(); err != nil {
		log.Fatal(err)
	}

	if forceHeight > 0 {
		s.eth.SetForceHeight(forceHeight)
	}

	go s.rpc.ServerThread()
	go s.eth.FetchandSaveDataLoop()
}

/*
	func (s *BOServer) Run(forceHeight uint64) {
		lastHeight, err := s.dao.GetLatestBlockHeight()
		if err != nil {
			log.Fatalf("fetch latest blcok heigh failed, err:%v", err)
		}

		// 初始运行时，以及后续需要回退一定距离时可以强行设置lastHeight
		if forceHeight > 0 && (lastHeight == 0 || (lastHeight > 0 && forceHeight < lastHeight)) {
			lastHeight = forceHeight
		}
		log.Printf("LastHeight %v", lastHeight)
		// cache global state info
		if err := s.ethcli.FetchAndSaveFamilies(); err != nil {
			log.Fatalf("fetch and save families failed, err: %v", err)
		}

		s.cache.FetchAndSaveBasicSeniorData()
		fmt.Println("fetched basic senior data")
		go s.rpc.Run()
		go s.ethcli.FetchandSaveFamilesLoop()
		go s.ethcli.FetchandSaveDataLoop(lastHeight)
	}
*/
func (s *Listener) Close() {
	s.eth.Close()
	s.rpc.Close()
	s.cache.Close()
	s.dao.Close()
}
