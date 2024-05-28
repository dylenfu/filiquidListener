package listener

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filiquid/listener/model"
)

func (s *Listener) GetCurrentHeight() (uint64, error) {
	return s.client.BlockNumber(context.Background())
}

func (s *Listener) getLogs(start, end uint64, addr common.Address) []types.Log {
	query := setQuery(start, end, addr)
	for {
		logs, err := s.client.FilterLogs(context.Background(), query)
		if err == nil {
			return logs
		}
		log.Printf("fetch contract event log failed, err: %v\r\n", err)
	}
}

func setQuery(start, end uint64, contract common.Address) ethereum.FilterQuery {
	return ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
	}
}

func (s *Listener) checkMappingIndex(vlog *types.Log) bool {
	index := model.NewMappingIndex(vlog.BlockNumber, vlog.Index)
	if s.cache.GetHaveSeen(&index) {
		return false
	} else {
		s.cache.SetHaveSeen(&index)
		return true
	}
}
