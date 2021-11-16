package service

import (
	"flag"

	"github.com/namelessup/bilibili/app/job/main/passport-game-data/conf"
	"sync"
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
	// service init
	s = New(conf.Conf)
}
