package dao

import (
	"github.com/namelessup/bilibili/app/interface/live/app-room/conf"
	"github.com/namelessup/bilibili/library/log"
	"os"
	"testing"
)

var testDao *Dao

func TestMain(m *testing.M) {
	conf.ConfPath = "../cmd/test.toml"
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	testDao = New(conf.Conf)
	os.Exit(m.Run())
}
