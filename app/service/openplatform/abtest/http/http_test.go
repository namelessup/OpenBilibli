package http

import (
	"flag"
	"fmt"

	"github.com/namelessup/bilibili/app/service/openplatform/abtest/conf"
	"github.com/namelessup/bilibili/app/service/openplatform/abtest/service"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"

	_ "github.com/smartystreets/goconvey/convey"
)

var client *httpx.Client

func init() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(fmt.Errorf("conf.Init() error(%v)", err))
	}
	svr := service.New(conf.Conf)
	client = httpx.NewClient(conf.Conf.HTTPClient.Read)
	Init(conf.Conf, svr)
}
