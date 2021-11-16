package service

import (
	"flag"

	"github.com/namelessup/bilibili/app/service/main/account/conf"
)

var s *Service

func init() {
	flag.Parse()

	if err := conf.Init(); err != nil {
		panic(err)
	}

	s = New(conf.Conf)
}
