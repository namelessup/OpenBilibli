package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/admin/ep/melloi/conf"
	"github.com/namelessup/bilibili/app/admin/ep/melloi/http"
	"github.com/namelessup/bilibili/app/admin/ep/melloi/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_durationForClosingServer = 2000
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("melloi start")
	// ecode init
	ecode.Init(conf.Conf.Ecode)
	// service init
	s := service.New(conf.Conf)
	http.Init(conf.Conf, s)
	// init pprof conf.Conf.Perf
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		log.Info("melloi get a signal %s", si.String())
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("melloi exit")
			s.Close()
			time.Sleep(_durationForClosingServer)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
