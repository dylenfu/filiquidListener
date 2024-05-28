package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/core"
	"github.com/filiquid/listener/utils"
)

var (
	configPath     = flag.String("config", "config.json", "config file path")
	listenerHeight = flag.Uint64("listenerHeight", 0, "force height for iterate block events")
	fetcherHeight  = flag.Uint64("fetcherHeight", 0, "force height for data caller query starting")
	migrate        = flag.Bool("migrate", false, "create database tables")
	limiter        = utils.NewIPRateLimiter(1, 5)
)

func main() {
	flag.Parse()

	config.Setting(*configPath)
	s := core.NewListener()
	if *migrate {
		s.Migrate()
	} else {
		s.Run(*listenerHeight, *fetcherHeight)
		waitToExit(s)
	}
}

func waitToExit(s *core.Listener) {
	exit := make(chan bool)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for sig := range sc {
			log.Printf("listener received exit signal: %v.", sig.String())
			log.Println("closing server...")
			s.Close()
			close(exit)
			break
		}
	}()
	<-exit
}
