package service

import (
	"sync"

	"github.com/namelessup/bilibili/app/service/main/passport-auth/conf"
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
