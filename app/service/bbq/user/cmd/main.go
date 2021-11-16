package main

import (
	"context"
	"flag"
	"github.com/namelessup/bilibili/library/conf/paladin"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namelessup/bilibili/app/service/bbq/user/internal/conf"
	"github.com/namelessup/bilibili/app/service/bbq/user/internal/server/grpc"
	"github.com/namelessup/bilibili/app/service/bbq/user/internal/server/http"
	"github.com/namelessup/bilibili/app/service/bbq/user/internal/service"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
)

var (
	_confName      string
	_unameConfName string
)

func init() {
	//线下使用
	flag.StringVar(&_confName, "conf_name", "user.toml", "default config filename")
	flag.StringVar(&_unameConfName, "uname_conf_name", "uname.json", "default config filename")
}

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	if err := paladin.Watch(_confName, conf.Conf); err != nil {
		panic(err)
	}
	if err := paladin.Watch(_unameConfName, conf.UnameConf); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	log.Info("user-service start")
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	ecode.Init(conf.Conf.Ecode)
	svc := service.New(conf.Conf)
	grpcServ := grpc.New(conf.Conf.GRPC, svc)
	http.Init(conf.Conf, svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			grpcServ.Shutdown(context.Background())
			log.Info("user-service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
