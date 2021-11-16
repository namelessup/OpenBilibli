package http

import (
	"github.com/namelessup/bilibili/app/infra/discovery/conf"
	"github.com/namelessup/bilibili/app/infra/discovery/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	svr *service.Service
)

// Init init http
func Init(c *conf.Config, s *service.Service) {
	svr = s
	engineInner := bm.DefaultServer(c.BM.Inner)
	innerRouter(engineInner)
	if err := engineInner.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// innerRouter init local router api path.
func innerRouter(e *bm.Engine) {
	group := e.Group("/discovery")
	{
		group.POST("/register", register)
		group.POST("/renew", renew)
		group.POST("/cancel", cancel)
		group.GET("/fetch/all", fetchAll)
		group.GET("/fetch", fetch)
		group.GET("/fetchs", fetchs)
		group.GET("/poll", poll)
		group.GET("/polls", polls)
		//manager
		group.POST("/set", set)
		group.GET("/polling", polling)
		group.GET("/nodes", nodes)
	}
}
