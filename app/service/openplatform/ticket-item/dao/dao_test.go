package dao

import (
	"time"

	"github.com/namelessup/bilibili/app/service/openplatform/ticket-item/conf"
	"github.com/namelessup/bilibili/library/log"
)

func initConf() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
}

func startService() {
	initConf()
	d = New(conf.Conf)
	time.Sleep(time.Second * 2)
}
