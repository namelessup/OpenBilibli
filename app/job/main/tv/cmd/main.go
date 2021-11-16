package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/job/main/tv/conf"
	"github.com/namelessup/bilibili/app/job/main/tv/http"
	"github.com/namelessup/bilibili/app/job/main/tv/service/pgc"
	"github.com/namelessup/bilibili/app/job/main/tv/service/ugc"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

var (
	pgcsrv *pgc.Service
	ugcsrv *ugc.Service
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.Log)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	defer log.Close()
	log.Info("tv-job start")
	pgcsrv = pgc.New(conf.Conf)
	ugcsrv = ugc.New(conf.Conf)
	http.Init(conf.Conf)
	signalHandler()
}

func signalHandler() {
	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("get a signal %s, stop the consume process", si.String())
			pgcsrv.Close()
			ugcsrv.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
