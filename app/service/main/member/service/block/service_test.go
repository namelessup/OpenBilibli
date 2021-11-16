package block

import (
	"flag"
	"os"
	"testing"

	"github.com/namelessup/bilibili/app/service/main/member/conf"
	mbDao "github.com/namelessup/bilibili/app/service/main/member/dao"
)

var (
	s *Service
)

func TestMain(m *testing.M) {
	flag.Set("conf", "../../cmd/member-service-example.toml")
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s = New(conf.Conf, mbDao.New(conf.Conf).BlockImpl())
	os.Exit(m.Run())
}
