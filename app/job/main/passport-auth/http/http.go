package http

import (
	"github.com/namelessup/bilibili/app/job/main/passport-auth/conf"
	"github.com/namelessup/bilibili/app/job/main/passport-auth/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	svc *service.Service
)

// Init init
func Init(c *conf.Config) {
	initService(c)
	// init router
	engineInner := bm.DefaultServer(c.BM)
	outerRouter(engineInner)
	if err := engineInner.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// initService init services.
func initService(c *conf.Config) {
	svc = service.New(c)
}

// outerRouter init outer router api path.
func outerRouter(e *bm.Engine) {
	//init api
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
}

// this can delete
func howToStart(c *bm.Context) {
	out := "[\n {\n\ttitle: 如有问题请联系(企业微信) |&&|,\n\tname: 刘玄(小鱼生)\n },\n {\n\ttitle: 一键初始化项目文档,\n\turl: http://info.bilibili.co/pages/viewpage.action?pageId=7548250\n }\n]"
	c.String(0, out)
}
