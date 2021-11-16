package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/namelessup/bilibili/app/service/live/xuserex/server/http"
	"github.com/namelessup/bilibili/app/service/live/xuserex/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/service/live/resource/sdk"
	"github.com/namelessup/bilibili/app/service/live/xuserex/conf"
	"github.com/namelessup/bilibili/app/service/live/xuserex/server/grpc"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("[xuserex] start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	ecode.Init(conf.Conf.Ecode)
	http.Init(conf.Conf)
	svc := service.New(conf.Conf)
	svr, err := grpc.New(svc)
	if err != nil {
		panic(fmt.Sprintf("start xuser grpc server fail! %s", err))
	}

	titansSdk.Init(conf.Conf.Titan)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			if svr != nil {
				svr.Shutdown(context.Background())
			}
			log.Info("[xuserex] exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
