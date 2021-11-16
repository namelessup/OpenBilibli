package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/reply/conf"
	"github.com/namelessup/bilibili/app/job/main/reply/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var rpSvc *service.Service

// Init init http
func Init(c *conf.Config, svc *service.Service) {
	rpSvc = svc
	engine := bm.DefaultServer(c.BM)
	outerRouter(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// outerRouter init outer router api path.
func outerRouter(e *bm.Engine) {
	e.GET("/monitor/ping", ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := rpSvc.Ping(c); err != nil {
		log.Error("reply job ping error")
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
