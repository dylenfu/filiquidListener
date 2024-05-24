package rpc

// func (s *RPCServer) getInterestData(minter string) ([]model.InterestData, error) {
// 	raw, ok := s.cache.GetInterest(minter)
// 	if !ok {
// 		return nil, errors.New("no interest data cached")
// 	}
// 	now := time.Now()
// 	if ok && timeBefore(now, raw.FetchedTime) {
// 		return raw.Data, nil
// 	}
// 	result, err := s.dao.GetInterest(minter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	list := make([]model.InterestData, 0)
// 	for _, v := range result {
// 		if data := v.Down(); data == nil {
// 			return nil, errConvert
// 		} else {
// 			list = append(list, *data)
// 		}
// 	}

// 	s.cache.SetInterest(now, minter, list)
// 	return list, nil
// }

// func (s *RPCServer) getStakedData(staker string) ([]model.StakedData, error) {
// 	raw, ok := s.cache.GetStake(staker)
// 	now := time.Now()
// 	if ok && timeBefore(now, raw.FetchedTime) {
// 		return raw.Data, nil
// 	}
// 	result, err := s.dao.GetStakes(staker)
// 	if err != nil {
// 		return nil, err
// 	}

// 	list := make([]model.StakedData, 0)
// 	for _, v := range result {
// 		if data := v.Down(); data != nil {
// 			list = append(list, *data)
// 		}
// 	}

// 	s.cache.SetStake(now, staker, list)
// 	return list, nil
// }

// func (s *RPCServer) getUnstakedData(staker string) ([]model.UnstakedData, error) {
// 	raw, ok := s.cache.GetUnStake(staker)
// 	now := time.Now()
// 	if ok && timeBefore(now, raw.FetchedTime) {
// 		return raw.Data, nil
// 	}

// 	result, err := s.dao.GetUnstake(staker)
// 	if err != nil {
// 		return nil, err
// 	}

// 	list := make([]model.UnstakedData, 0)
// 	for _, v := range result {
// 		data := v.Down()
// 		list = append(list, *data)
// 	}

// 	s.cache.SetUnStake(now, staker, list)
// 	return list, nil
// }

// func (s *RPCServer) getWithdrawnFigData(staker string) ([]model.WithdrawnFigData, error) {
// 	raw, ok := s.cache.GetWithdrawFig(staker)
// 	now := time.Now()
// 	if ok && timeBefore(now, raw.FetchedTime) {
// 		return raw.Data, nil
// 	}

// 	result, err := s.dao.GetWithdrawFig(staker)
// 	if err != nil {
// 		return nil, err
// 	}

// 	list := make([]model.WithdrawnFigData, 0)
// 	for _, v := range result {
// 		data := v.Down()
// 		list = append(list, *data)
// 	}

// 	s.cache.SetWithdrawFig(now, staker, list)
// 	return list, nil
// }

// func (s *RPCServer) getProposalsData(mode string, proposer string, page string) (model.ProposedDataList, error) {
// 	start, end, err := processPage(page)
// 	if err != nil {
// 		return emptyProposal, err
// 	}
// 	if !checkMOde(mode) {
// 		return emptyProposal, errors.New("invalid mode")
// 	}

// 	var (
// 		raw           model.ProposedDataTuple
// 		ok            bool
// 		result        []string
// 		currentHeight = s.cache.GetCurrentHeight()
// 	)

// 	if proposer != "" {
// 		raw, ok = s.cache.GetProposedByProposer(proposer, mode)
// 		now := time.Now()
// 		if !ok {
// 			fmt.Println("xxxxxxxxis not ok")
// 		} else {
// 			fmt.Println(now.UnixNano(), raw.FetchedTime.Add(time.Duration(config.Conf.FetchDuration)*time.Second).UnixNano())
// 			fmt.Println("isBefore ", !timeBefore(now, raw.FetchedTime))
// 		}

// 		if !ok || !timeBefore(now, raw.FetchedTime) {
// 			switch mode {
// 			case ALL:
// 				result, err = s.dao.GetProposalIdByProposer(proposer)
// 			case VOTING:
// 				result, err = s.dao.GetProposalIdByProposerAndDeadline(proposer, currentHeight)
// 			case FINISHED:
// 				result, err = s.dao.GetFinishedProposalByProposer(proposer, currentHeight)
// 			case EXECUTED:
// 				result, err = s.dao.GetExcutedProposalByProposer(proposer)
// 			}
// 			if err != nil {
// 				return emptyProposal, err
// 			}
// 			raw.Data = make([]uint64, 0)
// 			for _, v := range result {
// 				proposalId, err := strconv.ParseUint(v, 10, 64)
// 				if err != nil {
// 					return emptyProposal, err
// 				}
// 				raw.Data = append(raw.Data, proposalId)
// 			}
// 			s.cache.SetProposedByProposer(now, proposer, mode, raw.Data)
// 		}
// 	} else {
// 		raw, ok = s.cache.GetProposed(mode)
// 		now := time.Now()
// 		if !ok || !timeBefore(now, raw.FetchedTime) {
// 			switch mode {
// 			case ALL:
// 				result, err = s.dao.GetAllProposalIds()
// 			case VOTING:
// 				result, err = s.dao.GetProposalIdByDeadline(currentHeight)
// 			case FINISHED:
// 				result, err = s.dao.GetFinishedProposal(currentHeight)
// 			case EXECUTED:
// 				result, err = s.dao.GetExcutedProposal()
// 			}
// 			if err != nil {
// 				return emptyProposal, err
// 			}
// 			raw.Data = make([]uint64, 0, 10)
// 			for _, v := range result {
// 				u, err := strconv.ParseUint(v, 10, 64)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				raw.Data = append(raw.Data, u)
// 			}
// 			s.cache.SetProposed(now, mode, raw.Data)
// 		}
// 	}
// 	if len(raw.Data) == 0 {
// 		return model.ProposedDataList{Ids: raw.Data, TotalLength: 0}, nil
// 	} else {
// 		start = utils.Min(start, len(raw.Data))
// 		end = utils.Min(end, len(raw.Data))
// 		return model.ProposedDataList{Ids: raw.Data[start:end], TotalLength: len(raw.Data)}, nil
// 	}
// }

// func (s *RPCServer) getVotesData(proposalId string, voter string, page string) (model.VotedDataList, error) {
// 	start, end, err := processPage(page)
// 	if err != nil {
// 		return emptyVote, err
// 	}

// 	var (
// 		raw model.VotedDataTuple
// 		ok  bool
// 	)
// 	if voter != "" {
// 		raw, ok = s.cache.GetVotedByVoter(voter, proposalId)
// 		now := time.Now()
// 		if !ok || !timeBefore(now, raw.FetchedTime) {
// 			result, err := s.dao.GetVoteByProposalIdAndVote(
// 				proposalId,
// 				voter,
// 			)
// 			if err != nil {
// 				return emptyVote, err
// 			}
// 			raw.Data = make([]model.VotedData, 0)
// 			for _, v := range result {
// 				data := v.Down()
// 				raw.Data = append(raw.Data, *data)
// 			}
// 			s.cache.SetVotedByVoter(now, voter, proposalId, raw.Data)
// 		}
// 	} else {
// 		raw, ok = s.cache.GetVoted(proposalId)
// 		now := time.Now()
// 		if !ok || !timeBefore(now, raw.FetchedTime) {
// 			result, err := s.dao.GetVoteByProposalId(proposalId)
// 			if err != nil {
// 				return emptyVote, err
// 			}
// 			raw.Data = make([]model.VotedData, 0)
// 			for _, v := range result {
// 				data := v.Down()
// 				raw.Data = append(raw.Data, *data)
// 			}
// 			s.cache.SetVoted(now, proposalId, raw.Data)
// 		}
// 	}
// 	if len(raw.Data) == 0 {
// 		return model.VotedDataList{Votes: raw.Data, TotalLength: 0}, nil
// 	} else {
// 		start = utils.Min(start, len(raw.Data))
// 		end = utils.Min(end, len(raw.Data))
// 		return model.VotedDataList{Votes: raw.Data[start:end], TotalLength: len(raw.Data)}, nil
// 	}
// }

// func processPage(page string) (start, end int, err error) {
// 	p, err := strconv.Atoi(page)
// 	if err != nil || p < 0 {
// 		return 0, 0, errPage
// 	}
// 	start, end = p*config.Conf.PageNum, (p+1)*config.Conf.PageNum
// 	if start < 0 || end < 0 || start > end {
// 		return 0, 0, errPage
// 	}
// 	return
// }

// func timeBefore(now, other time.Time) bool {
// 	return now.Before(other.Add(time.Duration(config.Conf.FetchDuration) * time.Second))
// }

// func checkMOde(mode string) bool {
// 	if mode != ALL && mode != VOTING && mode != FINISHED && mode != EXECUTED {
// 		return false
// 	}
// 	return true
// }
