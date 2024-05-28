package fetcher

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
)

var testcli *Fetcher

func TestMain(m *testing.M) {
	db, err := dao.NewDao(&config.DbConfig{
		User:     "tliquid",
		Password: "123456",
		Url:      "127.0.0.1:3306",
		DbName:   "liquid",
	})
	if err != nil {
		panic(err)
	}
	cache := cache.NewCacheData(db, 10000)
	testcli = NewFetcher(db, cache, &config.EClient{
		RPCUrl:      "https://calibration.filfox.info/rpc/v1",
		FetcherAddr: common.HexToAddress("0x86eAa93250538905ab4c9d81A3B28fa59Ae3c8B2"),
	})
	os.Exit(m.Run())
}

// go test -count=1 -v github.com/filiquid/listener/eth/fetcher -run TestFetchData
func TestFetchData(t *testing.T) {
	raw, err := testcli.Caller.FetchData(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(raw)
}

// go test -count=1 -v github.com/filiquid/listener/eth/fetcher -run TestGetFamilies
func TestGetFamilies(t *testing.T) {
	height := uint64(1651682)
	testcli.cache.FetchAndSaveBasicCache()
	testcli.FetchAndSaveFamilies(height)
	data := testcli.cache.GetFamilies()
	t.Log(string(data))
}
