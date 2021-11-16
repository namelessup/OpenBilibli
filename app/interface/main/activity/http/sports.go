package http

import (
	"github.com/namelessup/bilibili/app/interface/main/activity/model/sports"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func qq(c *bm.Context) {
	params := c.Request.Form
	p := new(sports.ParamQq)
	if err := c.Bind(p); err != nil {
		return
	}
	if p.Tp <= 0 && p.Route == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(sportsSvc.Qq(c, params, p))
}

func news(c *bm.Context) {
	params := c.Request.Form
	p := new(sports.ParamNews)
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(sportsSvc.News(c, params, p))
}
