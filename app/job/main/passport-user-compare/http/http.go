package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/passport-user-compare/conf"
	"github.com/namelessup/bilibili/app/job/main/passport-user-compare/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	srv *service.Service
)

// Init init http sever instance.
func Init(c *conf.Config) {
	initService(c)
	// init inner router
	// engine
	engIn := bm.DefaultServer(c.BM)
	innerRouter(engIn)
	// init inner server
	if err := engIn.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// initService init services.
func initService(c *conf.Config) {
	srv = service.New(c)
}

// innerRouter init inner router.
func innerRouter(e *bm.Engine) {
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := srv.Ping(c); err != nil {
		log.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
