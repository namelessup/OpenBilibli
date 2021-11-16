package main

import (
	"flag"
	"os"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/appstatic/conf"
	"github.com/namelessup/bilibili/app/admin/main/appstatic/http"
	"github.com/namelessup/bilibili/app/admin/main/appstatic/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/os/signal"
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
	log.Init(conf.Conf.XLog)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	// service init
	s = service.New(conf.Conf)
	http.Init(conf.Conf, s)
	log.Info("appstatic-admin start")
	signalHandler()
}

func signalHandler() {
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			log.Info("get a signal %s, stop the appstatic-admin process", si.String())
			s.Wait()
			s.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
