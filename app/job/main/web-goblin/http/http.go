package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/web-goblin/conf"
	"github.com/namelessup/bilibili/app/job/main/web-goblin/service/esports"
	"github.com/namelessup/bilibili/app/job/main/web-goblin/service/web"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	srvweb *web.Service
	srvesp *esports.Service
)

// Init init
func Init(c *conf.Config) {
	srvweb = web.New(c)
	srvesp = esports.New(c)
	engine := bm.DefaultServer(c.BM)
	router(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

func router(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
}

func ping(c *bm.Context) {
	if err := srvweb.Ping(c); err != nil {
		log.Error("web-goblin-job ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
