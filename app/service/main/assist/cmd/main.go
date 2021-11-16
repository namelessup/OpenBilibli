package main

import (
	"flag"
	"github.com/namelessup/bilibili/app/service/main/assist/conf"
	"github.com/namelessup/bilibili/app/service/main/assist/http"
	"github.com/namelessup/bilibili/app/service/main/assist/rpc/server"
	"github.com/namelessup/bilibili/app/service/main/assist/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.Xlog)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	defer log.Close()
	log.Info("assist-service start")
	// service init
	svr := service.New(conf.Conf)
	rpcSvr := server.New(conf.Conf, svr)
	http.Init(conf.Conf, svr)
	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("assist-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			rpcSvr.Close()
			time.Sleep(time.Second * 2)
			log.Info("assist-service exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
