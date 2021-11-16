package service

import (
	"flag"
	"os"
	"testing"

	"github.com/namelessup/bilibili/app/service/main/point/conf"
)

var (
	s *Service
)

func TestMain(m *testing.M) {
	flag.Set("conf", "../cmd/point-service.toml")
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s = New(conf.Conf)
	os.Exit(m.Run())
}
