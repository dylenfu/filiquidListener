package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/utils"
)

func (s *EClient) handleGovernanceEvent(name string, vlog types.Log) (any, error) {
	switch name {
	case eventProposed:
		raw, err := s.govFilter.ParseProposed(vlog)
		if err != nil {
			return nil, err
		}

		proposed := utils.FromProposed(raw, vlog.BlockNumber)
		data := new(dao.Proposal).Up(proposed)
		if err = s.dao.InsertProposal(data); err != nil {
			return nil, err
		}

		return proposed, nil

	case eventBonded:
		raw, err := s.govFilter.ParseBonded(vlog)
		if err != nil {
			return nil, err
		}

		return utils.FromBonded(raw), nil

	case eventUnbonded:
		raw, err := s.govFilter.ParseUnbonded(vlog)
		if err != nil {
			return nil, err
		}

		return utils.FromUnBonded(raw), nil

	case eventExecuted:
		raw, err := s.govFilter.ParseExecuted(vlog)
		if err != nil {
			return nil, err
		}

		executed := utils.FromExcuted(raw)

		if err = s.dao.UpdateProposal(fmt.Sprintf("%d", raw.Result),
			raw.ProposalId.Uint64(),
		); err != nil {
			return nil, err
		}

		return executed, nil

	case eventVoted:
		raw, err := s.govFilter.ParseVoted(vlog)
		if err != nil {
			return nil, err
		}

		voted := utils.FromVoted(raw, vlog.BlockNumber)
		data := new(dao.Vote).Up(voted)

		return voted, s.dao.InsertVote(data)

	default:
		return nil, errEvent
	}
}
