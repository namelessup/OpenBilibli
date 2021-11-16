package service

import (
	"flag"
	"path/filepath"

	"github.com/namelessup/bilibili/app/service/main/usersuit/conf"
)

var (
	s *Service
)

func init() {
	dir, _ := filepath.Abs("../cmd/convey-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	s = New(conf.Conf)
}
