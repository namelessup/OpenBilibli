package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/answer/conf"
	"github.com/namelessup/bilibili/app/job/main/answer/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var svr *service.Service

// Init init a http server
func Init(c *conf.Config, s *service.Service) {
	svr = s
	engine := bm.DefaultServer(c.BM)
	innerRouter(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm.Start() error(%v)", err)
		panic(err)
	}
}

func innerRouter(e *bm.Engine) {
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := svr.Ping(c); err != nil {
		log.Error("answer-job service ping error")
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
