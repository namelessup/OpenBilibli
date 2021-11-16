package http

import (
	"github.com/namelessup/bilibili/app/job/main/appstatic/conf"
	"github.com/namelessup/bilibili/app/job/main/appstatic/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var apsSrv *service.Service

// Init .
func Init(c *conf.Config, srv *service.Service) {
	apsSrv = srv
	engineIn := bm.DefaultServer(c.HTTPServer)
	route(engineIn)
	// init inner server
	if err := engineIn.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
}

func ping(c *bm.Context) {
}
