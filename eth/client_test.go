package eth

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
)

var testcli *EClient

func TestMain(m *testing.M) {
	db, err := dao.NewDao(&config.DbConfig{
		User:     "tlistener",
		Password: "123456",
		Url:      "127.0.0.1:3306",
		DbName:   "filevents",
	})
	if err != nil {
		panic(err)
	}
	cache := cache.NewCacheData(db, 10000)
	testcli, err = NewETHClient(cache, db, &config.EClient{
		RPCUrl:         "https://calibration.filfox.info/rpc/v1",
		FiliquidAddr:   common.HexToAddress("0x8D6C016F996d441068F8085F42269d99d428883B"),
		GovernanceAddr: common.HexToAddress("0x36c16e5E89D239451C78bacbFDE67E03469C8253"),
		FitStakeAddr:   common.HexToAddress("0x53BE8BAf564a43Db16178156206F1383eF300aD5"),
	})
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

// go test -count=1 -v github.com/filiquid/listener/eth -run TestFilterLiquidEvent
func TestFilterLiquidEvent(t *testing.T) {
	num := uint64(1616692)
	start, end := num, num
	testcli.handleSingleContrctEvents("liquid", start, end)
}

// go test -count=1 -v github.com/filiquid/listener/eth -run TestFilterProposal
func TestFilterProposal(t *testing.T) {
	num := uint64(1556703)
	start, end := num, num
	testcli.handleSingleContrctEvents("governance", start, end)
}
