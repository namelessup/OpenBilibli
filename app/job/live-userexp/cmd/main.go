package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/namelessup/bilibili/app/job/live-userexp/conf"
	_ "github.com/namelessup/bilibili/app/job/live-userexp/model"
	"github.com/namelessup/bilibili/app/job/live-userexp/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/syscall"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.Log)
	log.Info("search-job start")
	defer log.Close()

	// service init
	srv := service.New(conf.Conf)

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		log.Info("live-userexp-job get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			log.Info("live-userexp-job exit")
			srv.Close()
			return
		case syscall.SIGHUP:
			log.Info("TODO: reload for syscall.SIGHUP")
			return
		default:
			return
		}
	}
}
