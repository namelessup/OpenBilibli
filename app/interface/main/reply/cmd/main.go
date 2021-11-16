package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/interface/main/reply/conf"
	"github.com/namelessup/bilibili/app/interface/main/reply/http"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/exp/feature"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus/report"
)

func main() {
	feature.DefaultGate.AddFlag(flag.CommandLine)
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.XLog)
	ecode.Init(conf.Conf.Ecode)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	defer log.Close()
	log.Info("reply start")
	// init report agent
	report.InitUser(conf.Conf.UserReport)
	// init http
	http.Init(conf.Conf)
	// init pprof
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("reply get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("reply exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
