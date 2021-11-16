package service

import (
	"github.com/namelessup/bilibili/app/job/main/identify/conf"
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
	s = New(conf.Conf)
}
