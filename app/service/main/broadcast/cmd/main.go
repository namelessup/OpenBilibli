package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/namelessup/bilibili/app/service/main/broadcast/server/grpc"
	"github.com/namelessup/bilibili/app/service/main/broadcast/server/http"
	"github.com/namelessup/bilibili/app/service/main/broadcast/service"
	"github.com/namelessup/bilibili/library/conf/env"
	"github.com/namelessup/bilibili/library/conf/paladin"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/naming"
	"github.com/namelessup/bilibili/library/naming/discovery"
	"github.com/namelessup/bilibili/library/net/ip"
)

const (
	ver = "v1.4.4"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	var (
		ac struct {
			Discovery *discovery.Config
		}
	)
	if err := paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	log.Init(nil)
	defer log.Close()
	log.Info("broadcast-service %s start", ver)
	// use internal discovery
	dis := discovery.New(ac.Discovery)
	// new a service
	srv := service.New(dis)
	ecode.Init(nil)
	http.Init(srv)
	// grpc server
	rpcSrv, rpcPort := server.New(srv)
	rpcSrv.Start()
	// register discovery
	var (
		err    error
		cancel context.CancelFunc
	)
	if env.IP == "" {
		ipAddr := ip.InternalIP()
		// broadcast discovery
		ins := &naming.Instance{
			Zone:     env.Zone,
			Env:      env.DeployEnv,
			Hostname: env.Hostname,
			AppID:    "push.service.broadcast",
			Addrs: []string{
				"grpc://" + ipAddr + ":" + rpcPort,
			},
		}
		cancel, err = dis.Register(context.Background(), ins)
		if err != nil {
			panic(err)
		}
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("broadcast-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("broadcast-service %s exit", ver)
			if cancel != nil {
				cancel()
			}
			rpcSrv.Shutdown(context.Background())
			time.Sleep(time.Second * 2)
			srv.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
