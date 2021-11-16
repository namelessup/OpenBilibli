package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/bbq/push/conf"
	"github.com/namelessup/bilibili/app/service/bbq/push/server/grpc"
	"github.com/namelessup/bilibili/app/service/bbq/push/server/http"
	"github.com/namelessup/bilibili/app/service/bbq/push/service"
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
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	ecode.Init(conf.Conf.Ecode)
	serv := service.New(conf.Conf)
	http.Init(conf.Conf, serv)
	gsrv, err := grpc.New(conf.Conf.RPCConf, serv)
	if err != nil {
		log.Error("grpc server start error: %s", err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			gsrv.Shutdown(ctx)
			log.Info("exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
