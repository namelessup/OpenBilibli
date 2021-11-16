package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/main/videoup/conf"
	"github.com/namelessup/bilibili/app/service/main/videoup/http"
	"github.com/namelessup/bilibili/app/service/main/videoup/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus/report"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.Xlog)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	defer log.Close()
	ecode.Init(conf.Conf.Ecode)
	report.InitUser(nil)
	log.Info("github.com/namelessup/bilibili/app/service/videoup start")
	// service init
	svr := service.New(conf.Conf)
	// statsd init
	http.Init(conf.Conf, svr)
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("github.com/namelessup/bilibili/app/service/videoup get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			svr.Close()
			log.Info("github.com/namelessup/bilibili/app/service/videoup exit")
			time.Sleep(1 * time.Second)
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
