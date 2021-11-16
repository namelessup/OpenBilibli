package main

import (
	"flag"
	"os"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/apm/conf"
	"github.com/namelessup/bilibili/app/admin/main/apm/http"
	"github.com/namelessup/bilibili/app/admin/main/apm/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/os/signal"
	"github.com/namelessup/bilibili/library/queue/databus/report"
	"github.com/namelessup/bilibili/library/syscall"
)

var (
	s *service.Service
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	ecode.Init(conf.Conf.Ecode)
	log.Init(conf.Conf.Log)
	defer log.Close()
	report.InitManager(conf.Conf.ManagerReport)
	// service init
	s = service.New(conf.Conf)
	http.Init(conf.Conf, s)
	log.Info("apm-admin start")
	signalHandler()
}

func signalHandler() {
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		log.Info("get a signal %s, stop the apm-admin process", si.String())
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
