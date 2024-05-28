package eth

import (
	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/eth/fetcher"
	"github.com/filiquid/listener/eth/listener"
)

type EClient struct {
	listener *listener.Listener
	fetcher  *fetcher.Fetcher
}

func NewETHClient(cache *cache.CacheData, dao *dao.Dao, cfg *config.EClient) (*EClient, error) {
	client := &EClient{}
	client.listener = listener.NewListener(dao, cache, cfg)
	client.fetcher = fetcher.NewFetcher(dao, cache, cfg)
	return client, nil
}

func (c *EClient) Close() {
	c.listener.Close()
	c.fetcher.Close()
}

func (c *EClient) Run(listenerForceHeight, fetcherForceHeight uint64) {
	go c.listener.IterateOnChainEvents(listenerForceHeight)
	go c.fetcher.IterateDataCallerQuerys(fetcherForceHeight)
}
