package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/job/main/search/conf"
	"github.com/namelessup/bilibili/app/job/main/search/http"
	_ "github.com/namelessup/bilibili/app/job/main/search/model"
	"github.com/namelessup/bilibili/app/job/main/search/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.XLog)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	log.Info("search-job start")
	defer log.Close()

	// service init
	srv := service.New(conf.Conf)

	http.Init(conf.Conf, srv)

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("search-job get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("search-job exit")
			srv.Close()
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
