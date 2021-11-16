package main

import (
	"context"
	"flag"
	"github.com/namelessup/bilibili/app/service/main/identify-game/server/grpc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/main/identify-game/conf"
	rpc "github.com/namelessup/bilibili/app/service/main/identify-game/rpc/server"
	"github.com/namelessup/bilibili/app/service/main/identify-game/server/http"
	"github.com/namelessup/bilibili/app/service/main/identify-game/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	// init conf,log,trace,stat,perf.
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Xlog)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	// service init
	svr := service.New(conf.Conf)
	rpcSvr := rpc.New(conf.Conf, svr)
	http.Init(conf.Conf, svr)

	// init warden server
	ws := grpc.New(conf.Conf.WardenServer, svr)

	// signal handler
	log.Info("identify-game-service start")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("identify-game-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			rpcSvr.Close()
			time.Sleep(time.Second * 2)
			svr.Close()
			ws.Shutdown(context.Background())
			log.Info("identify-game-service exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
