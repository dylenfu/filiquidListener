package rpc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/filiquid/listener/cache"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/utils"
)

const (
	FUNC_GETFAMILYCOUNT           = "getFamilyCount"
	FUNC_GETFAMILYAMOUNT          = "getFamilyAmount"
	FUNC_GETDISTINCTSTAKERSAMOUNT = "getDistinctStakersAmount"
	FUNC_GETINTERESTS             = "getInterests"
	FUNC_GETSTAKES                = "getStakes"
	FUNC_GETUNSTAKES              = "getUnstakes"
	FUNC_GETWITHDRAWNFIGS         = "getWithdrawnFigs"
	FUNC_GETPROPOSALS             = "getProposals"
	FUNC_GETVOTES                 = "getVotes"

	//Panel data API
	FUNCGETDATABASIC  = "getbasic"
	FUNCGETDATASENIOR = "getsenior"
	FUNCGETDATAPANEL  = "getpanel"
	FUNCGETFAMILIES   = "getfamilies"
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
	mux.HandleFunc("/"+FUNC_GETFAMILYCOUNT, s.getFamilyList)
	mux.HandleFunc("/"+FUNC_GETFAMILYAMOUNT, s.getFamilyAmount)
	mux.HandleFunc("/"+FUNC_GETDISTINCTSTAKERSAMOUNT, s.getDistinctStakersAmount)
	mux.HandleFunc("/"+FUNC_GETINTERESTS, s.getInterest)
	mux.HandleFunc("/"+FUNC_GETSTAKES, s.getStaked)
	mux.HandleFunc("/"+FUNC_GETUNSTAKES, s.getUnstaked)
	mux.HandleFunc("/"+FUNC_GETWITHDRAWNFIGS, s.getWithdrawnFig)
	mux.HandleFunc("/"+FUNC_GETPROPOSALS, s.getProposals)
	mux.HandleFunc("/"+FUNC_GETVOTES, s.getVotes)
	mux.HandleFunc("/"+FUNCGETDATABASIC, s.getCachedDataBasic)
	mux.HandleFunc("/"+FUNCGETDATASENIOR, s.getCachedDataSenior)
	mux.HandleFunc("/"+FUNCGETDATAPANEL, s.getCachedDataPanel)
	mux.HandleFunc("/"+FUNCGETFAMILIES, s.getCachedFamilies)

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
		fmt.Println("------xxx  getintp getbasic")
		writeHeader(w)

		basicData, _, _ := s.cache.GetBasicSeniorPanel()
		if basicData != nil {
			w.Write(basicData)
		} else {
			fmt.Println("------xxxxxx777 no data")
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
