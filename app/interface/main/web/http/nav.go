package http

import (
	"github.com/namelessup/bilibili/app/interface/main/web/model"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func nav(c *bm.Context) {
	var (
		rawMid interface{}
		ok     bool
	)

	if rawMid, ok = c.Get("mid"); !ok {
		// NOTE NoLogin here only for web
		c.JSON(model.FailedNavResp{}, ecode.NoLogin)
		return
	}
	mid := rawMid.(int64)
	c.JSON(webSvc.Nav(c, mid, c.Request.Header.Get("Cookie")))
}
