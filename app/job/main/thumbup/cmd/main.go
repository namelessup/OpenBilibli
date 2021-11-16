package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/job/main/thumbup/conf"
	"github.com/namelessup/bilibili/app/job/main/thumbup/server/http"
	"github.com/namelessup/bilibili/app/job/main/thumbup/service"
	"github.com/namelessup/bilibili/library/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	srv := service.New(conf.Conf)
	http.Init(conf.Conf, srv)
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
