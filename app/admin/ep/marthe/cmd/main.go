package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/admin/ep/marthe/conf"
	"github.com/namelessup/bilibili/app/admin/ep/marthe/server/http"
	"github.com/namelessup/bilibili/app/admin/ep/marthe/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("start")
	defer trace.Close()
	ecode.Init(conf.Conf.Ecode)

	s := service.New(conf.Conf)

	http.Init(conf.Conf, s)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		log.Info("get a signal %s", si.String())
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("exit")
			s.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
