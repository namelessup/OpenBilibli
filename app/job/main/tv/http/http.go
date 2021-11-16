package http

import (
	"github.com/namelessup/bilibili/app/job/main/tv/conf"
	xreport "github.com/namelessup/bilibili/app/job/main/tv/service/report"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var report *xreport.Service

// Init init http service
func Init(c *conf.Config) {
	report = xreport.New(c)
	// init inner router
	engineIn := bm.DefaultServer(c.HTTPServer)
	innerRouter(engineIn)
	// init inner server
	if err := engineIn.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// innerRouter init local router api path.
func innerRouter(e *bm.Engine) {
	// ping monitor
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
}
