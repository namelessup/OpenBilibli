package http

import (
	"github.com/namelessup/bilibili/app/job/main/ugcpay/conf"
	"github.com/namelessup/bilibili/app/job/main/ugcpay/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	srv *service.Service
)

// Init init
func Init(s *service.Service) {
	srv = s
	engine := bm.DefaultServer(conf.Conf.BM)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm Start error(%+v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
}

func ping(c *bm.Context) {
}
