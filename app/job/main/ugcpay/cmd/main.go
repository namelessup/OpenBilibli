package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/job/main/ugcpay/conf"
	"github.com/namelessup/bilibili/app/job/main/ugcpay/server/http"
	"github.com/namelessup/bilibili/app/job/main/ugcpay/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("ugcpay-job start")
	ecode.Init(conf.Conf.Ecode)
	svc := service.New(conf.Conf)
	http.Init(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			log.Info("ugcpay-job exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
