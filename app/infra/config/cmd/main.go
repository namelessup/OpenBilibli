package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelessup/bilibili/app/infra/config/conf"
	"github.com/namelessup/bilibili/app/infra/config/http"
	"github.com/namelessup/bilibili/app/infra/config/rpc/server"
	"github.com/namelessup/bilibili/app/infra/config/service/v1"
	"github.com/namelessup/bilibili/app/infra/config/service/v2"
	"github.com/namelessup/bilibili/library/conf/env"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/naming"
	"github.com/namelessup/bilibili/library/naming/discovery"
	xip "github.com/namelessup/bilibili/library/net/ip"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.Log)
	defer log.Close()
	// service init
	svr2 := v2.New(conf.Conf)
	svr := v1.New(conf.Conf)
	rpcSvr := rpc.New(conf.Conf, svr, svr2)
	http.Init(conf.Conf, svr, svr2, rpcSvr)
	// start discovery register
	var (
		err    error
		cancel context.CancelFunc
	)
	if env.IP == "" {
		ip := xip.InternalIP()
		hn, _ := os.Hostname()
		dis := discovery.New(nil)
		ins := &naming.Instance{
			Zone:     env.Zone,
			Env:      env.DeployEnv,
			AppID:    "config.service",
			Hostname: hn,
			Addrs: []string{
				"http://" + ip + ":" + env.HTTPPort,
				"gorpc://" + ip + ":" + env.GORPCPort,
			},
		}
		if cancel, err = dis.Register(context.Background(), ins); err != nil {
			panic(err)
		}
	}
	// end discovery register

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("config-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if cancel != nil {
				cancel()
			}
			rpcSvr.Close()
			svr.Close()
			log.Info("config-service exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
