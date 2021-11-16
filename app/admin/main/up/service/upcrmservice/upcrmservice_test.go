package upcrmservice

import (
	"flag"
	"os"
	"testing"

	"github.com/namelessup/bilibili/app/admin/main/up/conf"
)

var (
	s *Service
)

func TestMain(m *testing.M) {
	flag.Set("conf", "../../cmd/up-admin.toml")
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s = New(conf.Conf)
	m.Run()
	os.Exit(m.Run())
}
