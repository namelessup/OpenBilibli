package server

import (
	"flag"
	"path/filepath"
	"testing"
	"time"
	_ "time"

	"github.com/namelessup/bilibili/app/service/main/upcredit/conf"
	"github.com/namelessup/bilibili/app/service/main/upcredit/service"
	"github.com/namelessup/bilibili/library/net/rpc"
	xtime "github.com/namelessup/bilibili/library/time"
)

func init() {
	dir, _ := filepath.Abs("../../cmd/upcredit-service.toml")
	flag.Set("conf", dir)
}

func initSvrAndClient(t *testing.T) (client *rpc.Client, err error) {
	if err = conf.Init(); err != nil {
		t.Errorf("conf.Init() error(%v)", err)
		t.FailNow()
	}
	svr := service.New(conf.Conf)
	New(conf.Conf, svr)

	client = rpc.Dial("127.0.0.1:6079", xtime.Duration(time.Second), nil)
	return
}

func TestInfo(t *testing.T) {
	client, err := initSvrAndClient(t)
	defer client.Close()
	if err != nil {
		t.Errorf("rpc.Dial error(%v)", err)
		t.FailNow()
	}
	//time.Sleep(1 * time.Second)
}
