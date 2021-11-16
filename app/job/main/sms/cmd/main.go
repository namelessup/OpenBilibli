package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/job/main/sms/conf"
	"github.com/namelessup/bilibili/app/job/main/sms/http"
	"github.com/namelessup/bilibili/app/job/main/sms/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus/report"
)

var (
	srv *service.Service
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("sms-job start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	srv = service.New(conf.Conf)
	http.Init(conf.Conf, srv)
	report.InitUser(conf.Conf.UserReport)
	signalHandler()
}

func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("get a signal %s, stop the consume process", si.String())
			srv.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
