package member

import (
	"flag"

	"github.com/namelessup/bilibili/app/interface/main/account/conf"
)

var (
	s *Service
)

func init() {
	flag.Set("conf", "../../cmd/account-interface-example.toml")
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s = New(conf.Conf)
}
