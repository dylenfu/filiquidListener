package listener

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
)

var testcli *Listener

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
	testcli = NewListener(db, cache, &config.EClient{
		RPCUrl:         "https://calibration.filfox.info/rpc/v1",
		FiliquidAddr:   common.HexToAddress("0x46f7719da0F79Ac64a455526055dB42703c780d3"),
		GovernanceAddr: common.HexToAddress("0x8a1c17Fe673aaed1A85cA9EbB0fb494e9D43Dc6A"),
		FitStakeAddr:   common.HexToAddress("0x7bf370666854A69ad103BdA0F059746572ea3134"),
	})
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

// go test -count=1 -v github.com/filiquid/listener/eth/listener -run TestFilterLiquidEvent
func TestFilterLiquidEvent(t *testing.T) {
	num := uint64(1513000)
	start, end := num, num
	testcli.handleSingleContrctEvents("liquid", start, end)
}

// go test -count=1 -v github.com/filiquid/listener/eth/listener -run TestFilterProposal
func TestFilterProposal(t *testing.T) {
	num := uint64(1556703)
	start, end := num, num
	testcli.handleSingleContrctEvents("governance", start, end)
}
