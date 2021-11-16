package http

import (
	"github.com/namelessup/bilibili/app/job/main/block/conf"
	"github.com/namelessup/bilibili/app/job/main/block/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	svc *service.Service
)

// Init init
func Init() {
	initService()
	engine := bm.DefaultServer(conf.Conf.BM)
	router(engine)
	if err := engine.Start(); err != nil {
		log.Error("%+v", err)
		panic(err)
	}
}

// initService init services.
func initService() {
	svc = service.New()
}

// router init inner router api path.
func router(e *bm.Engine) {
	//init api
	e.GET("/monitor/ping", ping)
}

// ping check server ok.
func ping(c *bm.Context) {
}
