package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"
	"github.com/namelessup/bilibili/app/service/live/xlottery/conf"
	"github.com/namelessup/bilibili/app/service/live/xlottery/server/grpc"
	"github.com/namelessup/bilibili/app/service/live/xlottery/server/http"
	"github.com/namelessup/bilibili/app/service/live/xlottery/service"
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
	log.Info("lottery-service start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	ecode.Init(conf.Conf.Ecode)
	svc := service.New(conf.Conf)
	gsvr := grpc.Init(conf.Conf)
	http.Init(conf.Conf, svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			gsvr.Shutdown(context.Background())
			svc.Close()
			log.Info("lottery-service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
