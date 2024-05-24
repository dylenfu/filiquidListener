package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/utils"
)

func (s *EClient) handleLiquidEvent(name string, vlog types.Log) (any, error) {
	switch name {
	case eventDeposit:
		raw, err := s.liquidFilter.ParseDeposit(vlog)
		if err != nil {
			return nil, err
		}
		return utils.FromDeposit(raw), nil

	case eventRedeem:
		raw, err := s.liquidFilter.ParseRedeem(vlog)
		if err != nil {
			return nil, err
		}
		return utils.FromRedeem(raw), nil

	case eventCollateralizingMiner:
		raw, err := s.liquidFilter.ParseCollateralizingMiner(vlog)
		if err != nil {
			return nil, err
		}

		mCollaMiner := utils.FromCollateralizingMiner(raw)
		sender := mCollaMiner.Sender.Hex()
		if err := s.dao.InsertUser(sender); err != nil {
			return nil, err
		}
		if err = s.dao.IncreaseMinerCount(sender); err != nil {
			return nil, err
		}
		if err = s.dao.InsertFamily(&dao.Family{FamilyAddr: sender, MinerId: raw.MinerId}); err != nil {
			return nil, err
		}
		return mCollaMiner, nil

	case eventUncollateralizingMiner:
		raw, err := s.liquidFilter.ParseUncollateralizingMiner(vlog)
		if err != nil {
			return nil, err
		}

		mUnCollaMiner := utils.FromUnCollateralizingMiner(raw)
		sender := mUnCollaMiner.Sender.Hex()
		minerId := mUnCollaMiner.MinerId.Uint64()
		if err = s.dao.DeleteFamily(minerId); err != nil {
			return nil, err
		}
		if err = s.dao.DecreaseMinerCount(sender); err != nil {
			return nil, err
		}
		if err = s.dao.DeleteUser(sender); err != nil {
			return nil, err
		}
		return mUnCollaMiner, nil

	case eventBorrow:
		raw, err := s.liquidFilter.ParseBorrow(vlog)
		if err != nil {
			return nil, err
		}

		return utils.FromBorrow(raw), nil

	case eventPayback:
		raw, err := s.liquidFilter.ParsePayback(vlog)
		if err != nil {
			return nil, err
		}

		return utils.FromPayback(raw), nil

	case eventLiquidate:
		raw, err := s.liquidFilter.ParseLiquidate(vlog)
		if err != nil {
			return nil, err
		}

		return utils.FromLiquidate(raw), nil

	case eventBorrowUpdated:
		raw, err := s.liquidFilter.ParseBorrowUpdated(vlog)
		if err != nil {
			return nil, err
		}

		return utils.FromBorrowUpdated(raw), nil

	case eventBorrowDropped:
		raw, err := s.liquidFilter.ParseBorrowDropped(vlog)
		if err != nil {
			return nil, err
		}

		return utils.FromBorrowDropped(raw), nil

	default:
		return nil, errEvent
	}
}
