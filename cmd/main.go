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
	configPath = flag.String("config", "config.json", "config file path")
	start      = flag.Uint64("forceHeight", 0, "force height for scan history blocks")
	migrate    = flag.Bool("migrate", false, "create database tables")
	limiter    = utils.NewIPRateLimiter(1, 5)
)

func main() {
	flag.Parse()

	config.Setting(*configPath)
	s := core.NewListener()
	if *migrate {
		s.Migrate()
	} else {
		s.Run(*start)
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
