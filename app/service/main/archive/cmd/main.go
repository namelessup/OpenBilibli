package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/main/archive/conf"
	rpc "github.com/namelessup/bilibili/app/service/main/archive/server/gorpc"
	"github.com/namelessup/bilibili/app/service/main/archive/server/grpc"
	"github.com/namelessup/bilibili/app/service/main/archive/server/http"
	"github.com/namelessup/bilibili/app/service/main/archive/service"
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
	// init ecode
	ecode.Init(nil)
	// init log
	log.Init(conf.Conf.Xlog)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	defer log.Close()
	log.Info("archive-service start")
	// service init
	svr := service.New(conf.Conf)
	// statsd init
	rpcSvr := rpc.New(conf.Conf, svr)
	grpcSvr, err := grpc.New(nil, svr)
	if err != nil {
		panic(err)
	}
	http.Init(conf.Conf, svr)
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("archive-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			rpcSvr.Close()
			grpcSvr.Shutdown(context.TODO())
			time.Sleep(time.Second * 2)
			log.Info("archive-service exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
