package fetcher

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	//testcli.client.CallContract()
	//height := uint64(1632915)
	// rBasic, err := testcli.caller.FetchData(nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// abi, err := data.DataMetaData.GetAbi()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// for id, v := range abi.Methods {
	// 	t.Log(id, common.Bytes2Hex(v.ID), v)
	// }

	//var out [6]common.Address
	// list, err := testcli.Caller.GetAddresses(nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(list)
	//testcli.client.CallContract()
	//caller.CallerABI
	//abi.UnpackRevert()
	// raw, err := json.Marshal(rBasic)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(string(raw))

	// testcli.Caller.MetaData.GetAbi()

	// data, err := instanceABI.Pack("getValue")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// msg := ethereum.CallMsg{
	// 	To:   &contractAddress,
	// 	Data: data,
	// }

	// result, err := client.CallContract(context.Background(), msg, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client, err := ethclient.Dial("https://calibration.filfox.info/rpc/v1")
	if err != nil {
		t.Fatal(err)
	}
	contract := common.HexToAddress("0x8a1c17Fe673aaed1A85cA9EbB0fb494e9D43Dc6A") // governance
	//contract := common.HexToAddress("0x86eAa93250538905ab4c9d81A3B28fa59Ae3c8B2") // dataFetcher
	t.Log(client.BlockNumber(context.Background()))

	//raw := common.Hex2Bytes("0xa39fac12") // dataFetcher.getAddresses
	raw := common.Hex2Bytes("0x4e69d560") // governance.getStatus
	msg := ethereum.CallMsg{
		//From: common.HexToAddress("0x655CAA41be50c97FD75193a4A173BaBBf9A5A36C"),
		To:   &contract,
		Data: raw,
		//Gas:      100000,
		//GasPrice: big.NewInt(1000000000000000),
	}

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		t.Fatal(err)
	}
	// raw 0xa39fac12
	t.Log(result)
}
