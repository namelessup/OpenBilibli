package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/service/main/share/conf"
	"github.com/namelessup/bilibili/app/service/main/share/server/grpc"
	"github.com/namelessup/bilibili/app/service/main/share/server/http"
	"github.com/namelessup/bilibili/app/service/main/share/service"
	"github.com/namelessup/bilibili/library/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.Xlog)
	defer log.Close()
	log.Info("share-service start")
	svr := service.New(conf.Conf)
	grpcSvr := grpc.New(conf.Conf.WardenServer, svr)
	http.Init(conf.Conf, svr)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("share-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			grpcSvr.Shutdown(context.Background())
			log.Info("share-service exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
