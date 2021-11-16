package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/job/main/reply/conf"
	"github.com/namelessup/bilibili/app/job/main/reply/http"
	"github.com/namelessup/bilibili/app/job/main/reply/service"
	"github.com/namelessup/bilibili/library/exp/feature"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

var (
	s *service.Service
)

func main() {
	feature.DefaultGate.AddFlag(flag.CommandLine)
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.XLog)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	defer log.Close()
	log.Info("reply_consumer start")
	s = service.New(conf.Conf)
	http.Init(conf.Conf, s)
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
			if err = s.Close(); err != nil {
				log.Error("close consumer error(%v)", err)
			}
			s.Wait()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
