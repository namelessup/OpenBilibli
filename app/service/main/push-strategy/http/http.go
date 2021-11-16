package http

import (
	"github.com/namelessup/bilibili/app/service/main/push-strategy/conf"
	"github.com/namelessup/bilibili/app/service/main/push-strategy/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	idfSrv *verify.Verify
	srv    *service.Service
)

// Init .
func Init(c *conf.Config, svc *service.Service) {
	srv = svc
	idfSrv = verify.New(c.Verify)
	engine := bm.DefaultServer(c.HTTPServer)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start() error(%v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/x/internal/push-strategy", idfSrv.Verify)
	{
		g.POST("/task/add", addTask)
	}
}

func ping(ctx *bm.Context) {
	if err := srv.Ping(ctx); err != nil {
		ctx.Error = err
		ctx.AbortWithStatus(503)
	}
}
