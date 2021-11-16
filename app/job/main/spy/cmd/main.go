package main

import (
	"flag"
	"os"
	"time"

	"github.com/namelessup/bilibili/app/job/main/spy/conf"
	"github.com/namelessup/bilibili/app/job/main/spy/http"
	"github.com/namelessup/bilibili/app/job/main/spy/service"
	"github.com/namelessup/bilibili/library/log"
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
	log.Init(conf.Conf.Xlog)
	defer log.Close()

	// service init
	s = service.New(conf.Conf)
	http.Init(conf.Conf, s)

	log.Info("spy-job start")
	signalHandler()
}

func signalHandler() {
	var (
		err error
		ch  = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			log.Info("get a signal %s, stop the spy-job process", si.String())
			if err = s.Close(); err != nil {
				log.Error("close spy-job error(%v)", err)
			}
			s.Wait()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
