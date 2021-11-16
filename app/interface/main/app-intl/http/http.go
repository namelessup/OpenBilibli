package http

import (
	"github.com/namelessup/bilibili/app/interface/main/app-intl/conf"
	"github.com/namelessup/bilibili/app/interface/main/app-intl/service/feed"
	"github.com/namelessup/bilibili/app/interface/main/app-intl/service/player"
	"github.com/namelessup/bilibili/app/interface/main/app-intl/service/search"
	"github.com/namelessup/bilibili/app/interface/main/app-intl/service/view"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	// depend service
	authSvc   *auth.Auth
	verifySvc *verify.Verify
	// self service
	feedSvc   *feed.Service
	viewSvc   *view.Service
	playerSvc *player.Service
	searchSvc *search.Service
)

// Init is
func Init(c *conf.Config) {
	initService(c)
	// init external router
	engineOut := bm.DefaultServer(c.BM.Outer)
	outerRouter(engineOut)
	// init outer server
	if err := engineOut.Start(); err != nil {
		log.Error("engineOut.Start() error(%v)", err)
		panic(err)
	}
}

// initService init services.
func initService(c *conf.Config) {
	authSvc = auth.New(nil)
	verifySvc = verify.New(nil)
	// init self service
	feedSvc = feed.New(c)
	viewSvc = view.New(c)
	playerSvc = player.New(c)
	searchSvc = search.New(c)
}

// outerRouter init outer router api path.
func outerRouter(e *bm.Engine) {
	e.Ping(ping)

	feed := e.Group("/x/intl/feed")
	feed.GET("/index", authSvc.GuestMobile, feedIndex)

	view := e.Group("/x/intl/view")
	view.GET("", verifySvc.Verify, authSvc.GuestMobile, viewIndex)
	view.GET("/page", verifySvc.Verify, authSvc.GuestMobile, viewPage)

	e.GET("/x/intl/playurl", verifySvc.Verify, authSvc.GuestMobile, playurl)

	search := e.Group("/x/intl/search")
	search.GET("", authSvc.GuestMobile, searchAll)
	search.GET("/type", authSvc.GuestMobile, searchByType)
	search.GET("/suggest3", authSvc.GuestMobile, suggest3)
}
