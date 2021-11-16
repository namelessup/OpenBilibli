package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/dm2/conf"
	"github.com/namelessup/bilibili/app/job/main/dm2/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var svc *service.Service

// Init new dm2 job service
func Init(c *conf.Config, s *service.Service) {
	svc = s
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
	e.Ping(ping)
}

// ping check whether server is ok
func ping(c *bm.Context) {
	if err := svc.Ping(c); err != nil {
		log.Error("dm2-job service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
