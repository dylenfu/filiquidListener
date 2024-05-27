package rpc

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/utils"
)

const (
	FnGetFamilyCount           = "getFamilyCount"
	FnGetFamilyAmount          = "getFamilyAmount"
	FnGetDistinctStakersAmount = "getDistinctStakersAmount"
	FnGetInterests             = "getInterests"
	FnGetStakes                = "getStakes"
	FnGetUnstakes              = "getUnstakes"
	FnGetWithdrawFigs          = "getWithdrawnFigs"
	FnGetProposals             = "getProposals"
	FnGetVotes                 = "getVotes"
	FnGetBasicData             = "getBasic"
	FnGetSeniorData            = "getSenior"
	FnGetPanel                 = "getPanel"
	FnGetFamilies              = "getFamilies"
)

const (
	ALL      = "all"
	VOTING   = "voting"
	FINISHED = "finished"
	EXECUTED = "executed"
)

type RPCServer struct {
	port  string
	dao   *dao.Dao
	cache *cache.CacheData
}

func NewRPCServer(port string, dao *dao.Dao, cache *cache.CacheData) *RPCServer {
	return &RPCServer{
		port:  port,
		dao:   dao,
		cache: cache,
	}
}

func (s *RPCServer) ServerThread() {
	mux := http.NewServeMux()
	mux.HandleFunc("/"+FnGetFamilyCount, s.getFamilyList)
	mux.HandleFunc("/"+FnGetFamilyAmount, s.getFamilyAmount)
	mux.HandleFunc("/"+FnGetDistinctStakersAmount, s.getDistinctStakersAmount)
	mux.HandleFunc("/"+FnGetInterests, s.getInterest)
	mux.HandleFunc("/"+FnGetStakes, s.getStaked)
	mux.HandleFunc("/"+FnGetUnstakes, s.getUnstaked)
	mux.HandleFunc("/"+FnGetWithdrawFigs, s.getWithdrawnFig)
	mux.HandleFunc("/"+FnGetProposals, s.getProposals)
	mux.HandleFunc("/"+FnGetVotes, s.getVotes)
	mux.HandleFunc("/"+FnGetBasicData, s.getCachedDataBasic)
	mux.HandleFunc("/"+FnGetSeniorData, s.getCachedDataSenior)
	mux.HandleFunc("/"+FnGetPanel, s.getCachedDataPanel)
	mux.HandleFunc("/"+FnGetFamilies, s.getCachedFamilies)

	if err := http.ListenAndServe(":"+s.port, mux); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

func (s *RPCServer) Close() {
}

func (s *RPCServer) getFamilyList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeHeader(w)
		list := s.cache.GetFamilyList()
		if list != nil {
			w.Write(list)
		}
	}
}

func (s *RPCServer) getFamilyAmount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeHeader(w)
		num := s.cache.GetFamilyCount()
		if raw, err := utils.ToJson(num); err == nil {
			w.Write(raw)
		} else {
			w.Write(toJson(err.Error()))
			log.Printf("getFamilyAmount failed, err: %v", err)
		}
	}
}

func (s *RPCServer) getDistinctStakersAmount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeHeader(w)
		num := s.cache.GetStakerNum()
		if raw, err := utils.ToJson(num); err == nil {
			w.Write(raw)
		} else {
			w.Write(toJson(err.Error()))
			log.Printf("getDistinctStakersAmount failed, err: %v", err)
		}
	}
}

func (s *RPCServer) getInterest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		minter := trimQueryField(r, "minter")
		result, err := s.getInterestData(minter)
		if err != nil {
			log.Printf("getInterest failed, err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(toJson(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(toJson(result))
		}
	}
}

func (s *RPCServer) getStaked(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		staker := trimQueryField(r, "staker")
		result, err := s.getStakedData(staker)
		if err != nil {
			log.Printf("getStaked failed, err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(toJson(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(toJson(result))
		}
	}
}

func (s *RPCServer) getUnstaked(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		staker := trimQueryField(r, "staker")
		result, err := s.getUnstakedData(staker)
		if err != nil {
			log.Printf("getUnstaked failed, err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(toJson(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(toJson(result))
		}
	}
}

func (s *RPCServer) getWithdrawnFig(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		staker := trimQueryField(r, "staker")
		result, err := s.getWithdrawnFigData(staker)
		if err != nil {
			log.Printf("getWithdrawnFig failed, err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(toJson(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(toJson(result))
		}
	}
}

func (s *RPCServer) getProposals(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		mode := trimQueryField(r, "mode")
		proposer := trimQueryField(r, "proposer")
		page := trimQueryField(r, "page")
		result, err := s.getProposalsData(mode, proposer, page)
		if err != nil {
			log.Printf("getProposals failed, err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(toJson(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(toJson(result))
		}
	}
}

func (s *RPCServer) getVotes(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		proposalId := trimQueryField(r, "id")
		voter := trimQueryField(r, "voter")
		page := trimQueryField(r, "page")
		result, err := s.getVotesData(proposalId, voter, page)
		if err != nil {
			log.Printf("getVotes failed, err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(toJson(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(toJson(result))
		}
	}
}

func (s *RPCServer) getCachedDataBasic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeHeader(w)

		basicData, _, _ := s.cache.GetBasicSeniorPanel()
		if basicData != nil {
			w.Write(basicData)
		}
	}
}

func (s *RPCServer) getCachedDataSenior(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeHeader(w)

		_, seniorData, _ := s.cache.GetBasicSeniorPanel()
		if seniorData != nil {
			w.Write(seniorData)
		}
	}
}

func (s *RPCServer) getCachedDataPanel(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeHeader(w)

		_, _, panelData := s.cache.GetBasicSeniorPanel()
		if panelData != nil {
			w.Write(panelData)
		}
	}
}

func (s *RPCServer) getCachedFamilies(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeHeader(w)

		if data := s.cache.GetFamilies(); data != nil {
			w.Write(data)
		}
	}
}

func writeHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func toJson(data any) []byte {
	raw, _ := json.Marshal(data)
	return raw
}

func trimQueryField(r *http.Request, name string) string {
	raw := r.URL.Query().Get(name)
	raw = strings.TrimSpace(raw)
	raw = strings.Trim(raw, "\"")
	raw = strings.Trim(raw, "'")
	raw = strings.ToLower(raw)
	return raw
}
