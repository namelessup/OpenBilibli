package main

import (
	"flag"
	"os"
	"time"

	"github.com/namelessup/bilibili/app/job/main/answer/conf"
	"github.com/namelessup/bilibili/app/job/main/answer/http"
	"github.com/namelessup/bilibili/app/job/main/answer/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/os/signal"
	"github.com/namelessup/bilibili/library/syscall"
	"github.com/namelessup/bilibili/library/text/translate/chinese"
)

var (
	svr *service.Service
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	chinese.Init()
	svr = service.New(conf.Conf)
	http.Init(conf.Conf, svr)
	log.Info("answer-job start")
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		log.Info("answer-job  get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			log.Info("answer-job  exit")
			if err := svr.Close(); err != nil {
				log.Error("srv close consumer error(%v)", err)
			}
			time.Sleep(2 * time.Second)
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
