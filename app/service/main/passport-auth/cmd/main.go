package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/main/passport-auth/conf"
	"github.com/namelessup/bilibili/app/service/main/passport-auth/http"
	"github.com/namelessup/bilibili/app/service/main/passport-auth/rpc/grpc"
	rpc "github.com/namelessup/bilibili/app/service/main/passport-auth/rpc/server"
	"github.com/namelessup/bilibili/app/service/main/passport-auth/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Xlog)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	svr := service.New(conf.Conf)
	// rpc server init
	rpcSvr := rpc.New(conf.Conf, svr)
	ws := grpc.New(conf.Conf.WardenServer, svr)
	http.Init(conf.Conf, svr)
	// signal handler
	log.Info("passport-auth-service start")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("passport-auth-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ws.Shutdown(context.Background())
			rpcSvr.Close()
			time.Sleep(time.Second * 2)
			log.Info("passport-auth-service exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
