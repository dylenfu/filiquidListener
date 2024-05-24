package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/utils"
)

func (s *EClient) handleFitStakeEvent(name string, vlog types.Log) (any, error) {
	switch name {
	case eventInterest:
		raw, err := s.fitFilter.ParseInterest(vlog)
		if err != nil {
			return nil, err
		}

		interest := utils.FromInterest(raw, vlog.BlockNumber)
		data := new(dao.Interest).Up(interest)

		return interest, s.dao.InsertInterest(data)

	case eventStaked:
		raw, err := s.fitFilter.ParseStaked(vlog)
		if err != nil {
			return nil, err
		}
		staked := utils.FromStaked(raw, vlog.BlockNumber)
		data := new(dao.Stake).Up(staked)

		return staked, s.dao.InsertStake(data)

	case eventUnstaked:
		raw, err := s.fitFilter.ParseUnstaked(vlog)
		if err != nil {
			return nil, err
		}

		unstaked := utils.FromUnstaked(raw, vlog.BlockNumber)
		data := new(dao.UnStake).Up(unstaked)

		return unstaked, s.dao.InsertUnstake(data)

	case eventWithdrawnFig:
		raw, err := s.fitFilter.ParseWithdrawnFIG(vlog)
		if err != nil {
			return nil, err
		}

		withdraw := utils.FromWithdrawFig(raw, vlog.BlockNumber)
		data := new(dao.WithdrawFig).Up(withdraw)
		return withdraw, s.dao.InsertWithdrawFig(data)

	default:
		return nil, errEvent
	}
}
