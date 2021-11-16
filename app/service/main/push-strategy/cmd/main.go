package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/main/push-strategy/conf"
	"github.com/namelessup/bilibili/app/service/main/push-strategy/http"
	"github.com/namelessup/bilibili/app/service/main/push-strategy/service"
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
	log.Init(conf.Conf.Log)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	log.Info("push-strategystart")
	ecode.Init(conf.Conf.Ecode)
	srv := service.New(conf.Conf)
	http.Init(conf.Conf, srv)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("push-strategy get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			srv.Close()
			time.Sleep(1 * time.Second)
			log.Info("push-strategyexit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
