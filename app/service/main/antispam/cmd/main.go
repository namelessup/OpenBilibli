package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/main/antispam/conf"
	"github.com/namelessup/bilibili/app/service/main/antispam/http"
	rpc "github.com/namelessup/bilibili/app/service/main/antispam/rpc/server"
	"github.com/namelessup/bilibili/app/service/main/antispam/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"

	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(conf.ConfPath); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	ecode.Init(conf.Conf.Ecode)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	defer log.Close()
	log.Info("antispam start")
	svr := service.New(conf.Conf)
	rpcSvr := rpc.New(conf.Conf, svr)
	http.Init(conf.Conf, svr)
	// init signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("get a signal %s, stop the consume process", si.String())
			rpcSvr.Close()
			time.Sleep(time.Second * 2)
			svr.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
