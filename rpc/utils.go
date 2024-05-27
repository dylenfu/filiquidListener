package rpc

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/model"
	"github.com/filiquid/listener/utils"
)

var (
	errPage    = errors.New("invalid page")
	errConvert = errors.New("convert data failed")

	emptyProposal = model.ProposedDataList{}
	emptyVote     = model.VotedDataList{}
)

func (s *RPCServer) getInterestData(minter string) ([]model.InterestData, error) {
	result, err := s.dao.GetInterest(minter)
	if err != nil {
		return nil, err
	}

	list := make([]model.InterestData, 0)
	for _, v := range result {
		if data := v.Down(); data == nil {
			return nil, errConvert
		} else {
			list = append(list, *data)
		}
	}

	return list, nil
}

func (s *RPCServer) getStakedData(staker string) ([]model.StakedData, error) {
	result, err := s.dao.GetStakes(staker)
	if err != nil {
		return nil, err
	}

	list := make([]model.StakedData, 0)
	for _, v := range result {
		if data := v.Down(); data != nil {
			list = append(list, *data)
		}
	}

	return list, nil
}

func (s *RPCServer) getUnstakedData(staker string) ([]model.UnstakedData, error) {
	result, err := s.dao.GetUnstake(staker)
	if err != nil {
		return nil, err
	}

	list := make([]model.UnstakedData, 0)
	for _, v := range result {
		data := v.Down()
		list = append(list, *data)
	}

	return list, nil
}

func (s *RPCServer) getWithdrawnFigData(staker string) ([]model.WithdrawnFigData, error) {
	result, err := s.dao.GetWithdrawFig(staker)
	if err != nil {
		return nil, err
	}

	list := make([]model.WithdrawnFigData, 0)
	for _, v := range result {
		data := v.Down()
		list = append(list, *data)
	}

	return list, nil
}

func (s *RPCServer) getProposalsData(mode string, proposer string, page string) (model.ProposedDataList, error) {
	start, end, err := processPage(page)
	if err != nil {
		return emptyProposal, err
	}
	if !checkMode(mode) {
		return emptyProposal, errors.New("invalid mode")
	}

	var (
		result        []string
		currentHeight = s.cache.GetCurrentHeight()
		list          []uint64
	)

	if proposer != "" {
		switch mode {
		case ALL:
			result, err = s.dao.GetProposalIdByProposer(proposer)
		case VOTING:
			result, err = s.dao.GetProposalIdByProposerAndDeadline(proposer, currentHeight)
		case FINISHED:
			result, err = s.dao.GetFinishedProposalByProposer(proposer, currentHeight)
		case EXECUTED:
			result, err = s.dao.GetExcutedProposalByProposer(proposer)
		}
	} else {
		switch mode {
		case ALL:
			result, err = s.dao.GetAllProposalIds()
		case VOTING:
			result, err = s.dao.GetProposalIdByDeadline(currentHeight)
		case FINISHED:
			result, err = s.dao.GetFinishedProposal(currentHeight)
		case EXECUTED:
			result, err = s.dao.GetExcutedProposal()
		}
	}
	if err != nil {
		return emptyProposal, err
	}
	for _, v := range result {
		proposalId, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return emptyProposal, err
		}
		list = append(list, proposalId)
	}

	if len(list) == 0 {
		return model.ProposedDataList{Ids: list, TotalLength: 0}, nil
	} else {
		start = utils.Min(start, len(list))
		end = utils.Min(end, len(list))
		return model.ProposedDataList{Ids: list[start:end], TotalLength: len(list)}, nil
	}
}

func (s *RPCServer) getVotesData(proposalId string, voter string, page string) (model.VotedDataList, error) {
	start, end, err := processPage(page)
	if err != nil {
		return emptyVote, err
	}

	var (
		result []dao.Vote
	)
	if voter != "" {
		result, err = s.dao.GetVoteByProposalIdAndVote(proposalId, voter)
	} else {
		result, err = s.dao.GetVoteByProposalId(proposalId)
	}
	if err != nil {
		return emptyVote, err
	}
	list := make([]model.VotedData, 0)
	for _, v := range result {
		data := v.Down()
		list = append(list, *data)
	}
	if len(list) == 0 {
		return model.VotedDataList{Votes: list, TotalLength: 0}, nil
	} else {
		start = utils.Min(start, len(list))
		end = utils.Min(end, len(list))
		return model.VotedDataList{Votes: list[start:end], TotalLength: len(list)}, nil
	}
}

func processPage(page string) (start, end int, err error) {
	p, err := strconv.Atoi(page)
	if err != nil || p < 0 {
		return 0, 0, errPage
	}
	start, end = p*config.Conf.PageNum, (p+1)*config.Conf.PageNum
	if start < 0 || end < 0 || start > end {
		return 0, 0, errPage
	}
	return
}

func checkMode(mode string) bool {
	fmt.Println(mode, ALL, mode == ALL)
	if mode != ALL && mode != VOTING && mode != FINISHED && mode != EXECUTED {
		return false
	}
	return true
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//enableCors(&w)
		target := r.RemoteAddr
		p := strings.LastIndex(target, ":")
		if p >= 0 {
			target = target[:p]
		}
		//log.Println("Visitor: ", target)
		/*if target != "127.0.0.1" {
			l := limiter.GetLimiter(target)
			if !l.Allow() {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
		}*/
		next.ServeHTTP(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
