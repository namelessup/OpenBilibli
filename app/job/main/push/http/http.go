package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/push/conf"
	"github.com/namelessup/bilibili/app/job/main/push/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var pushSrv *service.Service

// Init .
func Init(c *conf.Config, srv *service.Service) {
	pushSrv = srv
	engine := bm.DefaultServer(c.HTTPServer)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	e.GET("/refresh_token_cache", refreshTokenCache)
}

func ping(ctx *bm.Context) {
	if err := pushSrv.Ping(ctx); err != nil {
		log.Error("push-job ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func register(ctx *bm.Context) {
	ctx.JSON(map[string]interface{}{}, nil)
}

func refreshTokenCache(ctx *bm.Context) {
	go pushSrv.RefreshTokenCache()
	ctx.JSON(nil, nil)
}
