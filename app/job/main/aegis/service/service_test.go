package service

import (
	"context"
	"flag"
	"testing"

	"github.com/namelessup/bilibili/app/job/main/aegis/conf"
	"github.com/namelessup/bilibili/library/log"
)

var (
	s *Service
)

func initConf() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
}

func init() {
	flag.Set("conf", "../cmd/aegis-job.toml")
	initConf()
	s = New(conf.Conf)
}

func Test_syncReport(t *testing.T) {
	s.syncReport(context.Background())
}
