package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/conf"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/server/grpc"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/server/http"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/service"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/service/mis"
	"github.com/namelessup/bilibili/library/conf/paladin"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	if err := paladin.Watch("ticket-sales.toml", conf.Conf); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	log.Info("ticket-sales start")
	ecode.Init(nil)
	// service init
	srv := service.New(conf.Conf)
	misSrv := mis.New(srv.Get())
	grpc.New(srv, misSrv)
	http.Init(conf.Conf, srv)
	log.Info("ready to serv")
	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("ticket-sales get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			// rpcSvr.Close()
			time.Sleep(time.Second * 2)
			log.Info("ticket-sales exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
