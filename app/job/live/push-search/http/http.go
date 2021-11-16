package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/live/push-search/conf"
	"github.com/namelessup/bilibili/app/job/live/push-search/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	Srv *service.Service
	vfy *verify.Verify
)

// Init init
func Init(c *conf.Config) {
	Srv = service.New(c)
	vfy = verify.New(c.Verify)
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
	g := e.Group("/x/push-search")
	{
		g.GET("/start", vfy.Verify, howToStart)
	}
}

func ping(c *bm.Context) {
	if err := Srv.Ping(c); err != nil {
		log.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}

// example for http request handler
func howToStart(c *bm.Context) {
	c.String(0, "Golang 大法好 !!!")
}
