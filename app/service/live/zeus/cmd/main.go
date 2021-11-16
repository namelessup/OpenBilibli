package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/live/zeus/internal/conf"
	"github.com/namelessup/bilibili/app/service/live/zeus/internal/server/grpc"
	"github.com/namelessup/bilibili/app/service/live/zeus/internal/server/http"
	"github.com/namelessup/bilibili/app/service/live/zeus/internal/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	v1srv "github.com/namelessup/bilibili/app/service/live/zeus/internal/service/v1"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("zeus-service start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	ecode.Init(conf.Conf.Ecode)
	svc := service.New(conf.Conf)
	zeus := v1srv.NewZeusService(conf.Conf)
	grpcServer := grpc.New(conf.Conf, zeus)
	defer grpcServer.Shutdown(context.Background())
	http.Init(conf.Conf, svc, zeus)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			log.Info("zeus-service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
