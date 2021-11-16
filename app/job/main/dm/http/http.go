package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/dm/conf"
	"github.com/namelessup/bilibili/app/job/main/dm/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	dmSvc *service.Service
)

// Init new dm job service
func Init(c *conf.Config, s *service.Service) {
	dmSvc = s
	engine := bm.DefaultServer(c.HTTPServer)
	initRouter(engine)
	// run http server
	if err := engine.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// initRouter init router.
func initRouter(e *bm.Engine) {
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := dmSvc.Ping(c); err != nil {
		log.Error("dm-job service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
