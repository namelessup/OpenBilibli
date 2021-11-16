package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/interface/main/favorite/conf"
	"github.com/namelessup/bilibili/app/interface/main/favorite/http"
	"github.com/namelessup/bilibili/app/interface/main/favorite/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

var svc *service.Service

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
	log.Info("favorite start")
	ecode.Init(conf.Conf.Ecode)
	// service init
	svc = service.New(conf.Conf)
	http.Init(conf.Conf, svc)
	signalHandler()
}

func signalHandler() {
	var (
		err error
		ch  = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("get a signal %s, stop the consume process", si.String())
			svc.Close()
			if err != nil {
				log.Error("svc.Close() error(%v)", err)
			}
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
