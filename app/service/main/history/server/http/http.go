package http

import (
	"github.com/namelessup/bilibili/app/service/main/history/conf"
	"github.com/namelessup/bilibili/app/service/main/history/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	srv *service.Service
	vfy *verify.Verify
)

// Init init
func Init(c *conf.Config, svc *service.Service) {
	srv = svc
	vfy = verify.New(c.Verify)
	engine := bm.DefaultServer(c.BM)
	interRouter(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

func interRouter(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	g := e.Group("/x/internal/history", vfy.Verify)
	{
		g.POST("/add", add)
		g.POST("/add/multi", addHistories)
		g.POST("/del", del)
		g.POST("/clear", clear)
		g.GET("/user", userHistories)
		g.GET("/aids", histories)
		g.GET("/hide", userHide)
		g.POST("/hide/update", updateHide)
	}
}

func ping(c *bm.Context) {
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
