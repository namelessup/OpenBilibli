package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/admin/main/feed/conf"
	"github.com/namelessup/bilibili/app/admin/main/feed/http"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus/report"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	report.InitManager(conf.Conf.ManagerReport)
	// service init
	http.Init(conf.Conf)
	log.Info("feed-admin start")
	signalHandler()
}

func signalHandler() {
	var (
		c = make(chan os.Signal, 1)
	)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("feed-admin get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			log.Info("feed-admin exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
