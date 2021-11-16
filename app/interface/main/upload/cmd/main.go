package main

import (
	"flag"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/upload/conf"
	"github.com/namelessup/bilibili/app/interface/main/upload/http"
	"github.com/namelessup/bilibili/app/interface/main/upload/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/queue/databus/report"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	ecode.Init(conf.Conf.Ecode)
	// init log
	log.Init(conf.Conf.XLog)
	defer log.Close()

	// service init
	s := service.New(conf.Conf)
	http.Init(conf.Conf, s)
	report.InitUser(nil)
	log.Info("bfs-upload-interface start!")
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("upload get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(1 * time.Second)
			log.Info("upload exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
