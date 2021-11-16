package service

import (
	"github.com/namelessup/bilibili/app/service/main/passport/conf"
	"sync"
)

var (
	once sync.Once
	s    *Service
)

func startService() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	// service init
	s = New(conf.Conf)
}
