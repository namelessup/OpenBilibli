package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/interface/main/player/conf"
	"github.com/namelessup/bilibili/app/interface/main/player/http"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.XLog)
	defer log.Close()
	log.Info("play-interface start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	// ecode
	ecode.Init(conf.Conf.Ecode)
	// service init
	http.Init(conf.Conf)
	// monitor
	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("play-interface get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("play-interface exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
