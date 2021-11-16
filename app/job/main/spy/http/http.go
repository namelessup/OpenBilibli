package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/spy/conf"
	"github.com/namelessup/bilibili/app/job/main/spy/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	svr *service.Service
)

// Init .
func Init(c *conf.Config, s *service.Service) {
	svr = s
	// init inner router
	engine := bm.DefaultServer(c.HTTPServer)
	innerRouter(engine)
	// init local server
	if err := engine.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// innerRouter init inner router.
func innerRouter(e *bm.Engine) {
	// health check
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := svr.Ping(c); err != nil {
		log.Error("spy-job service ping error")
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
