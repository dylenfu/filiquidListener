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

func (s *Listener) Run(listenerForceHeight, fetcherForceHeight uint64) {
	go s.rpc.ServerThread()
	go s.eth.Run(listenerForceHeight, fetcherForceHeight)
}

func (s *Listener) Close() {
	s.eth.Close()
	s.rpc.Close()
	s.cache.Close()
	s.dao.Close()
}
