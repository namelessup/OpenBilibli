package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/main/account/conf"
	rpc "github.com/namelessup/bilibili/app/service/main/account/rpc/server"
	"github.com/namelessup/bilibili/app/service/main/account/server/grpc"
	"github.com/namelessup/bilibili/app/service/main/account/server/http"
	"github.com/namelessup/bilibili/app/service/main/account/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/rpc/warden/resolver/livezk"
	"github.com/namelessup/bilibili/library/net/trace"
)

const (
	discoveryID = "account.service"
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
	log.Info("account-service start")
	// service init
	svr := service.New(conf.Conf)
	rpcSvr := rpc.New(conf.Conf, svr)

	// warden init
	var wardensvr *warden.Server
	if conf.Conf.WardenServer != nil {
		var err error
		if wardensvr, err = grpc.Start(conf.Conf, svr); err != nil {
			panic(fmt.Sprintf("start warden server fail! %s", err))
		}
		cancel, err := livezk.Register(conf.Conf.LiveZK, conf.Conf.WardenServer.Addr, discoveryID)
		if err != nil {
			panic(fmt.Sprintf("register grpc service into live zookeeper error: %s", err))
		}
		defer cancel()
	}

	http.Init(conf.Conf, svr)
	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("account-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("account-service exit")
			rpcSvr.Close()
			if wardensvr != nil {
				wardensvr.Shutdown(context.Background())
			}
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
