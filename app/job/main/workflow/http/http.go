package http

import (
	"github.com/namelessup/bilibili/app/job/main/workflow/conf"
	"github.com/namelessup/bilibili/app/job/main/workflow/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	svr *service.Service
)

func Init(c *conf.Config, s *service.Service) {
	svr = s
	// init local router
	engine := bm.DefaultServer(c.HTTPServer)
	route(engine)
	// init local server
	if err := engine.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// route .
func route(e *bm.Engine) {
	e.Ping(ping)
}

// ping check server ok.
func ping(ctx *bm.Context) {
	if err := svr.Ping(ctx); err != nil {
		log.Error("workflow job ping error(%v)", err)
		ctx.AbortWithStatus(503)
	}
}
