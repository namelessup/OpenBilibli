package main

import (
	"flag"
	"os"
	"time"

	"github.com/namelessup/bilibili/app/service/main/spy/conf"
	rpc "github.com/namelessup/bilibili/app/service/main/spy/rpc/server"
	grpc "github.com/namelessup/bilibili/app/service/main/spy/server/grpc"
	"github.com/namelessup/bilibili/app/service/main/spy/server/http"
	"github.com/namelessup/bilibili/app/service/main/spy/service"
	"github.com/namelessup/bilibili/library/log"
	xrpc "github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/os/signal"
	"github.com/namelessup/bilibili/library/syscall"
)

var (
	svr     *service.Service
	rpcSvr  *xrpc.Server
	grpcSvr *warden.Server
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	// service init
	svr = service.New(conf.Conf)
	http.Init(conf.Conf, svr)
	rpcSvr = rpc.New(conf.Conf, svr)
	grpcSvr = grpc.New(conf.Conf.GRPC, svr)
	log.Info("spy-service start")
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		log.Info("spy-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			rpcSvr.Close()
			time.Sleep(time.Second * 2)
			log.Info("spy-service exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
