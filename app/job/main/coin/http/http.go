package http

import (
	"net/http"
	"strconv"

	"github.com/namelessup/bilibili/app/job/main/coin/conf"
	"github.com/namelessup/bilibili/app/job/main/coin/service"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var coinSvr *service.Service

// Init init http router.
func Init(c *conf.Config, s *service.Service) {
	// init internal router
	coinSvr = s
	// init outer router
	engineOuter := bm.DefaultServer(c.BM)
	outerRouter(engineOuter)
	if err := engineOuter.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

func outerRouter(r *bm.Engine) {
	r.Ping(ping)
	r.Register(register)
	r.POST("/redo", redo)
	r.POST("/settle", settle)
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := coinSvr.Ping(c); err != nil {
		log.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func redo(c *bm.Context) {
	idStr := c.Request.Form.Get("table_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, coinSvr.Redo(id))
}

func settle(c *bm.Context) {
	idStr := c.Request.Form.Get("table_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, coinSvr.Settle(id))
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
