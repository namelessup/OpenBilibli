package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/live/xroom/internal/server/grpc"
	"github.com/namelessup/bilibili/app/service/live/xroom/internal/server/http"
	"github.com/namelessup/bilibili/app/service/live/xroom/internal/service"
	"github.com/namelessup/bilibili/library/conf/paladin"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("xroom-service start")
	ecode.Init(nil)
	svc := service.New()
	grpcSrv := grpc.New(svc)
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, _ := context.WithTimeout(context.Background(), 35*time.Second)
			grpcSrv.Shutdown(ctx)
			httpSrv.Shutdown(ctx)
			log.Info("xroom-service exit")
			svc.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
