package http

import (
	"github.com/namelessup/bilibili/app/job/main/passport-game-cloud/conf"
	"github.com/namelessup/bilibili/app/job/main/passport-game-cloud/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	srv *service.Service
)

// Init init http sever instance.
func Init(c *conf.Config, s *service.Service) {
	srv = s
	// init inner router
	eng := bm.DefaultServer(c.BM)
	initRouter(eng)
	// init inner server
	if err := eng.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// initRouter init inner router.
func initRouter(e *bm.Engine) {
	// health check
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	c.JSON(nil, srv.Ping(c))
}
