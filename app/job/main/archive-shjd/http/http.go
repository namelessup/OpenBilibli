package http

import (
	"github.com/namelessup/bilibili/app/job/main/archive-shjd/conf"
	"github.com/namelessup/bilibili/app/job/main/archive-shjd/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"net/http"
)

var (
	arcSvr *service.Service
)

// Init init http router.
func Init(c *conf.Config, s *service.Service) {
	arcSvr = s
	e := bm.DefaultServer(c.BM)
	innerRouter(e)
	// init internal server
	if err := e.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// innerRouter init inner router.
func innerRouter(e *bm.Engine) {
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := arcSvr.Ping(); err != nil {
		log.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
