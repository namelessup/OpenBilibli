package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/videoup-task/conf"
	"github.com/namelessup/bilibili/app/admin/main/videoup-task/http"
	"github.com/namelessup/bilibili/app/admin/main/videoup-task/service"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/os/signal"
	"github.com/namelessup/bilibili/library/queue/databus/report"
	"github.com/namelessup/bilibili/library/syscall"
)

func main() {
	var err error
	//conf init
	flag.Parse()
	if err = conf.Init(); err != nil {
		panic(err)
	}
	//log init
	log.Init(conf.Conf.Xlog)
	defer log.Close()
	fmt.Printf("conf(%+v)", conf.Conf)
	report.InitManager(conf.Conf.ManagerReport)

	//trace init
	trace.Init(nil)
	defer trace.Close()

	//http init
	srv := service.New(conf.Conf)
	http.Init(conf.Conf, srv)

	//signal notify to change service behavior
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sch
		log.Info("videoup-task-admin got a signal %s", s.String())

		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("videoup-task-admin is closed")
			srv.Close()
			time.Sleep(time.Second * 1)
			return
		case syscall.SIGHUP:
			//reload
		default:
			return
		}
	}
}
