package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/job/main/videoup/conf"
	"github.com/namelessup/bilibili/app/job/main/videoup/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var svr *service.Service

// Init init http router.
func Init(c *conf.Config, s *service.Service) {
	svr = s
	eng := bm.NewServer(c.Bm)
	route(eng)
	if err := eng.Start(); err != nil {
		log.Error(" eng.Start() error(%v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := svr.Ping(c); err != nil {
		log.Error("svr.Ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
