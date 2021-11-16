package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/thumbup/conf"
	"github.com/namelessup/bilibili/app/job/main/thumbup/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var srv *service.Service

// Init init http server.
func Init(c *conf.Config, s *service.Service) {
	srv = s
	engine := bm.DefaultServer(nil)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

func route(engine *bm.Engine) {
	engine.Ping(ping)
}

func ping(ctx *bm.Context) {
	if err := srv.Ping(ctx); err != nil {
		log.Error("thumbup-job ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
