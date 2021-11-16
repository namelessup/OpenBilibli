package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/service/main/share/conf"
	"github.com/namelessup/bilibili/app/service/main/share/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	svr       *service.Service
	verifySvc *verify.Verify
)

// Init init http router.
func Init(c *conf.Config, s *service.Service) {
	svr = s
	verifySvc = verify.New(c.Verify)
	engine := bm.DefaultServer(c.BM)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	group := e.Group("/x/internal/share")
	{
		group.POST("/add", verifySvc.Verify, add)
		group.GET("/stat", verifySvc.Verify, stat)
		group.GET("/stats", verifySvc.Verify, stats)
	}
}

func ping(c *bm.Context) {
	if err := svr.Ping(); err != nil {
		log.Error("share-service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func register(c *bm.Context) {
	c.JSON(nil, nil)
}
