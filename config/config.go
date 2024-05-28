package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

const (
	CALIBRATION_GENESIS_UNIX_EPOCH = 1667326380

	FETCHDATAINTERVAL  = 5
	MAXIMUMRESULTS     = 100
	APYTHRESHOLD       = 50
	BLOCKSPERYEAR      = 1036800
	MAXNOTICELENGTH    = 10000
	MAXBANNERLENGTH    = 10000
	RATEBASE           = 1000000
	U_M                = 0.9
	DAY1               = 24 * 60 * 60
	DAY7               = DAY1 * 7
	MONTH              = DAY1 * 30
	MONTH3             = MONTH * 3
	RPCQUERYREPEATTIME = 6
)

var Conf *Config

type Config struct {
	FetchDuration  int    // default: 15 seconds
	PageNum        int    // default: 6
	MaxCacheSize   int    // default: 10000
	BlockScanRange int    // default: 2880
	Port           string // local rpc server port
	Db             *DbConfig
	Eth            *EClient
}

type DbConfig struct {
	User     string
	Password string // todo(xk): set pwd with stdin
	Url      string
	DbName   string
}

type EClient struct {
	RPCUrl         string
	FetcherAddr    common.Address
	FiliquidAddr   common.Address
	GovernanceAddr common.Address
	FitStakeAddr   common.Address
}

func Setting(path string) {
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Read config file faield, err: %v", err)
	}

	Conf = new(Config)
	if err := json.Unmarshal(raw, Conf); err != nil {
		log.Fatalf("Unmarshal config raw data failed, err: %v", err)
	}
}
