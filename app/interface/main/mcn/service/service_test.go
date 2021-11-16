package service

import (
	"flag"
	"os"
	"testing"

	"github.com/namelessup/bilibili/app/interface/main/mcn/conf"
)

var (
	s *Service
)

func TestMain(m *testing.M) {
	flag.Set("conf", "../cmd/mcn-interface.toml")
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s = New(conf.Conf)
	m.Run()
	os.Exit(m.Run())
}
