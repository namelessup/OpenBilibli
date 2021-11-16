package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/service/main/open/conf"
	"github.com/namelessup/bilibili/app/service/main/open/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	openSvc *service.Service
)

// Init init http sever instance.
func Init(c *conf.Config, s *service.Service) {
	openSvc = s
	// init inner router
	engineIn := bm.DefaultServer(nil)
	innerRouter(engineIn)
	// init inner server
	if err := engineIn.Start(); err != nil {
		log.Error("engineInner.Start error (%v)", err)
		panic(err)
	}
}

// innerRouter .
func innerRouter(e *bm.Engine) {
	e.Ping(ping)
	e.GET("/api/getsecret", openSvc.Verify, secret)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := openSvc.Ping(c); err != nil {
		log.Error("service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
