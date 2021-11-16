package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/admin/main/bfs/conf"
	"github.com/namelessup/bilibili/app/admin/main/bfs/http"
	"github.com/namelessup/bilibili/app/admin/main/bfs/service"
	"github.com/namelessup/bilibili/library/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("bfs admin start")
	c := make(chan os.Signal, 1)
	srv := service.New(conf.Conf)
	http.Init(conf.Conf, srv)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	log.Info("bfs admin singal notify complete")
	for {
		s := <-c
		log.Info("bfs admin get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			srv.Close()
			log.Info("bfs admin exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
