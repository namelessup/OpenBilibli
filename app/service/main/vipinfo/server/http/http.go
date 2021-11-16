package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/service/main/vipinfo/conf"
	"github.com/namelessup/bilibili/app/service/main/vipinfo/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	srv       *service.Service
	verifySvc *verify.Verify
)

// Init init
func Init(c *conf.Config, s *service.Service) {
	srv = s
	verifySvc = verify.New(nil)
	engine := bm.DefaultServer(c.BM)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm Start error(%v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	g := e.Group("/x/internal/vipinfo", verifySvc.Verify)
	{
		g.GET("/info", info)
		g.GET("/infos", infos)
	}
}

func ping(c *bm.Context) {
	if err := srv.Ping(c); err != nil {
		log.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
