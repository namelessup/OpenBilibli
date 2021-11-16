package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/app/conf"
	"github.com/namelessup/bilibili/app/job/main/app/service"
	"github.com/namelessup/bilibili/app/job/main/app/service/feed"
	"github.com/namelessup/bilibili/app/job/main/app/service/show"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	Svc     *service.Service
	ShowSvc *show.Service
	FeedSvc *feed.Service
)

// Init init http
func Init(c *conf.Config) {
	initService(c)
	// init external router
	engineIn := bm.DefaultServer(c.BM.Inner)
	innerRouter(engineIn)
	// init Inner server
	if err := engineIn.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

func initService(c *conf.Config) {
	Svc = service.New(c)
	ShowSvc = show.New(c)
	FeedSvc = feed.New(c)
}

// innerRouter init inner router api path.
func innerRouter(e *bm.Engine) {
	e.Ping(ping)
}

func ping(c *bm.Context) {
	err := Svc.Ping(c)
	if err == nil {
		err = ShowSvc.Ping(c)
	}
	if err != nil {
		log.Error("app-job service ping error(%+v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
