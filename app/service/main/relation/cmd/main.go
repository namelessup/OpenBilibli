package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"
	"github.com/namelessup/bilibili/app/service/main/relation/conf"
	"github.com/namelessup/bilibili/app/service/main/relation/http"
	"github.com/namelessup/bilibili/app/service/main/relation/rpc/server"
	"github.com/namelessup/bilibili/app/service/main/relation/server/grpc"
	"github.com/namelessup/bilibili/app/service/main/relation/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus/report"
)

func main() {
	flag.Parse()
	// init conf,log,trace,stat,perf.
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	// report
	report.InitUser(conf.Conf.Report)
	// service init	WardenServer  *warden.ServerConfig
	svr := service.New(conf.Conf)
	rpcSvr := rpc.New(conf.Conf, svr)
	//start grpc
	ws := grpc.New(conf.Conf, svr)
	http.Init(conf.Conf, svr)
	// signal handler
	log.Info("relation-service start")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("relation-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ws.Shutdown(context.Background())
			rpcSvr.Close()
			time.Sleep(time.Second * 2)
			log.Info("relation-service exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
