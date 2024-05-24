package eth

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/model"
)

func (c *EClient) getABSCurrentHeight() uint64 {
	for {
		height, err := c.client.BlockNumber(context.Background())
		if err == nil {
			return height
		}
		time.Sleep(2 * time.Second)
	}
}

func (s *EClient) GetCurrentHeight() (uint64, error) {
	return s.client.BlockNumber(context.Background())
}

var errRange = errors.New("lower greater than higher")

func (s *EClient) getScanRange() (current uint64, lower uint64, higher uint64, err error) {
	current = s.getABSCurrentHeight()
	last := s.cache.GetLastHeight()
	if last < current {
		lower = last + 1
	} else {
		lower = last
	}

	if s.forceHeight > current {
		err = errRange
		return
	}
	if s.forceHeight > 0 && s.forceHeight <= lower {
		err = errRange
		return
	}
	if s.forceHeight > lower {
		lower = s.forceHeight
		s.forceHeight = 0
	}

	higher = current
	if lower > higher {
		err = errRange
		return
	}
	if lower < higher-uint64(config.Conf.BlockScanRange) {
		//lower = higher - config.RANGE
		higher = lower + uint64(config.Conf.BlockScanRange)
	}
	return
}

func (s *EClient) getLogs(start, end uint64, addr common.Address) []types.Log {
	query := setQuery(start, end, addr)
	for {
		logs, err := s.client.FilterLogs(context.Background(), query)
		if err == nil {
			return logs
		}
		log.Printf("fetch contract event log failed, err: %v\r\n", err)
	}
}

func (s *EClient) checkMappingIndex(vlog *types.Log) bool {
	index := model.NewMappingIndex(vlog.BlockNumber, vlog.Index)
	if s.cache.GetHaveSeen(&index) {
		return false
	} else {
		s.cache.SetHaveSeen(&index)
		return true
	}
}

func setQuery(start, end uint64, contract common.Address) ethereum.FilterQuery {
	return ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
	}
}
