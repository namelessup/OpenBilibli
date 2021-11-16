package http

import (
	"github.com/namelessup/bilibili/app/admin/main/relation/conf"
	"github.com/namelessup/bilibili/app/admin/main/relation/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"
)

var (
	authSvc *permit.Permit
	svc     *service.Service
)

// Init init
func Init(c *conf.Config) {
	initService(c)
	// init router
	engine := bm.DefaultServer(c.BM)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

// initService init services.
func initService(c *conf.Config) {
	authSvc = permit.New(c.Auth)
	svc = service.New(c)
}

// initRouter init outer router api path.
func initRouter(e *bm.Engine) {
	//init api
	e.Ping(ping)
	group := e.Group("/x/admin/relation")
	{
		group.GET("/follower/followers", authSvc.Permit("RELATION_INFO"), followers)
		group.GET("/following/followings", authSvc.Permit("RELATION_INFO"), followings)
		group.GET("/logs", authSvc.Permit("RELATION_INFO"), logs)
		group.GET("/stat", stat)
		group.GET("/stats", stats)
	}
}

// ping check server ok.
func ping(c *bm.Context) {
	svc.Ping(c)
}
