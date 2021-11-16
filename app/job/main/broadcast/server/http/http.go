package http

import (
	"github.com/namelessup/bilibili/app/job/main/broadcast/conf"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Init init http.
func Init(c *conf.Config) {
	engine := bm.DefaultServer(c.HTTP)
	outerRouter(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

func outerRouter(e *bm.Engine) {
	e.Ping(ping)
}

func ping(c *bm.Context) {

}
