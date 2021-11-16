package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"
	"github.com/namelessup/bilibili/app/service/video/stream-mng/conf"
	"github.com/namelessup/bilibili/app/service/video/stream-mng/server/grpc"
	"github.com/namelessup/bilibili/app/service/video/stream-mng/server/http"
	"github.com/namelessup/bilibili/app/service/video/stream-mng/service"
	"github.com/namelessup/bilibili/library/conf/env"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	if env.DeployEnv == env.DeployEnvProd {
		log.Init(nil)
	} else {
		log.Init(conf.Conf.Log)
	}

	defer log.Close()
	log.Info("stream-mng-service start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	ecode.Init(conf.Conf.Ecode)
	svc := service.New(conf.Conf)
	http.Init(conf.Conf, svc)
	ws := grpc.New(nil, svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			ws.Shutdown(context.Background())
			log.Info("stream-mng-service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
