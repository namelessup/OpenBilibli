package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/admin/main/mcn/conf"
	"github.com/namelessup/bilibili/app/admin/main/mcn/server/http"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	manager "github.com/namelessup/bilibili/library/queue/databus/report"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.SetFormat("[%D %T] [%L] [%S] %M")
	log.Info("start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	// manager log init
	manager.InitManager(conf.Conf.ManagerLog)
	ecode.Init(conf.Conf.Ecode)
	http.Init(conf.Conf)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
