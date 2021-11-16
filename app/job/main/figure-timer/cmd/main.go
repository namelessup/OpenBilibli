package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/namelessup/bilibili/app/job/main/figure-timer/conf"
	"github.com/namelessup/bilibili/app/job/main/figure-timer/http"
	"github.com/namelessup/bilibili/app/job/main/figure-timer/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/syscall"
)

var (
	srv *service.Service
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	srv = service.New(conf.Conf)
	http.Init(srv)
	log.Info("figure-timer-job start")
	signalHandler()
}

func signalHandler() {
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		si := <-ch
		log.Info("figure-timer-job got a signal (%d)", si)
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			srv.Close()
			srv.Wait()
			log.Info("figure-timer-job exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
