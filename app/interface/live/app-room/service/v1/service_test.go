package v1

import (
	"github.com/namelessup/bilibili/app/interface/live/app-room/conf"
	"github.com/namelessup/bilibili/library/log"
	"os"
	"testing"
)

var (
	testGiftService *GiftService
	s               *GiftService
)

func TestMain(m *testing.M) {
	conf.ConfPath = "../../cmd/test.toml"
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
	testGiftService = NewGiftService(conf.Conf)
	s = testGiftService
	os.Exit(m.Run())
}
