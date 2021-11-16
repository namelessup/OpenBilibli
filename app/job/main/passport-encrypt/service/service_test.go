package service

import (
	"flag"
	"sync"

	"github.com/namelessup/bilibili/app/job/main/passport-encrypt/conf"
)

var (
	once sync.Once
	s    *Service
)

func startService() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s = New(conf.Conf)
}
