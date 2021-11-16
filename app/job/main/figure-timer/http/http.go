package http

import (
	"github.com/namelessup/bilibili/app/job/main/figure-timer/conf"
	"github.com/namelessup/bilibili/app/job/main/figure-timer/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	svc *service.Service
)

// Init init a http server
func Init(s *service.Service) {
	svc = s

	e := bm.DefaultServer(conf.Conf.BM)
	innerRouter(e)
	if err := e.Start(); err != nil {
		log.Error("%+v", err)
		panic(err)
	}
}

// innerRouter init inner router.
func innerRouter(e *bm.Engine) {
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	// if err := svc.Ping(c); err != nil {
	// 	log.Error("figure-timer-job ping err (%+v)", err)
	// 	c.AbortWithStatus(http.StatusServiceUnavailable)
	// }
}
