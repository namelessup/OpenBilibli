package newcomer

import (
	"flag"
	"os"
	"testing"

	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	"github.com/namelessup/bilibili/app/interface/main/creative/service"
)

var (
	s *Service
)

func TestMain(m *testing.M) {
	flag.Set("conf", "../../cmd/creative.toml")
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	rpcdaos := service.NewRPCDaos(conf.Conf)
	s = New(conf.Conf, rpcdaos)
	m.Run()
	os.Exit(m.Run())
}
