package eth

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/utils"
)

// todo
func (c *EClient) getFamilies() ([]byte, error) {
	raw := c.cache.GetFamilies()
	var list []common.Address
	if err := json.Unmarshal(raw, &list); err != nil {
		return nil, err
	}

	info, err := c.fetcher.GetBatchedUserBorrows(nil, list)
	if err != nil {
		return nil, err
	}

	return utils.ToJson(info)
}

func (c *EClient) FetchAndSaveFamilies() error {
	raw, err := c.getFamilies()
	if err != nil {
		return err
	}
	c.cache.SetFamilies(raw)
	return nil
}

// func (c *EClient) FetchBackendDataLoop(lastHeight uint64) {
// 	for {
// 		c.FetchBackendData(lastHeight)
// 		<-time.After(time.Second * config.SECONDSBETWEENFETCH)
// 		lastHeight += 1
// 	}
// }

// // fetch contiuniously once failed, todo(xk): process error
// func (c *EClient) FetchandSaveFamilesLoop() {
// 	for {
// 		if err := c.FetchAndSaveFamilies(); err != nil {
// 			log.Printf("FetchandSaveFamilesLoop fetch failed, err: %v\r\n", err)
// 		}
// 		<-time.After(time.Second * config.SECONDSBETWEENFETCH)
// 	}
// }

func (c *EClient) FetchBackendData(height uint64) error {
	rBasic, err := c.fetcher.FetchData(&bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(height),
	})
	if err != nil {
		log.Printf("Cannot get data, height %d, err: %v", height, err)
		return err
	}
	log.Printf("ETH Client Get height: %d", rBasic.BlockHeight.Uint64())
	dBasic := new(dao.BasicData).Up(&rBasic)
	if err := c.dao.InsertBasic(dBasic); err != nil {
		log.Printf("Insert basic data failed, err: %v", err)
		return err
	}

	if c.ticker == 0 {
		response, err := c.fetcher.GetTotalPendingInterest(nil)
		if err != nil {
			log.Printf("Get total pending intersest failed, err: %v", err)
			return err
		}

		result := new(dao.SeniorData).Up(response)
		if err := c.dao.InsertSenior(result); err != nil {
			log.Printf("Insert senior data failed, err: %v", err)
			return err
		}
	}
	c.ticker = (c.ticker + 1) % config.FETCHDATAINTERVAL
	if err := c.cache.FetchAndSaveBasicSeniorData(); err != nil {
		log.Printf("insert senior data failed, err: %v", err)
		return err
	}

	return nil
}
