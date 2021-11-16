package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/job/openplatform/article/conf"
	"github.com/namelessup/bilibili/app/job/openplatform/article/http"
	"github.com/namelessup/bilibili/app/job/openplatform/article/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	initConf()
	initLog()
	defer log.Close()
	log.Info("article-job start")
	srv := service.New(conf.Conf)
	http.Init(conf.Conf, srv)
	initSignal(srv)
}

func initConf() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
}

func initLog() {
	log.Init(conf.Conf.Xlog)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
}

func initSignal(srv *service.Service) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("article-job get a signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if err := srv.Close(); err != nil {
				log.Error("srv close consumer error(%+v)", err)
			}
			return
		case syscall.SIGHUP:
			// TODO: reload
		default:
			return
		}
	}
}
