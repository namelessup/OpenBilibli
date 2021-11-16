package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/main/tag/conf"
	rpc "github.com/namelessup/bilibili/app/service/main/tag/rpc/server"
	"github.com/namelessup/bilibili/app/service/main/tag/server/grpc"
	"github.com/namelessup/bilibili/app/service/main/tag/server/http"
	"github.com/namelessup/bilibili/app/service/main/tag/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("tag-service start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	svr := service.New(conf.Conf)
	ecode.Init(conf.Conf.Ecode)
	http.Init(conf.Conf, svr)
	rpcSvr := rpc.Init(conf.Conf, svr)
	grpcSvr := grpc.New(conf.Conf.GRPC, svr)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("tag-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			rpcSvr.Close()
			grpcSvr.Shutdown(context.TODO())
			time.Sleep(time.Second * 2)
			svr.Close()
			log.Info("tag-service exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
