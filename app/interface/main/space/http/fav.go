package http

import (
	"github.com/namelessup/bilibili/app/interface/main/space/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func favNav(c *bm.Context) {
	var mid int64
	v := new(struct {
		VMid int64 `form:"mid" validate:"min=1"`
	})
	if err := c.Bind(v); err != nil {
		return
	}
	if midInter, ok := c.Get("mid"); ok {
		mid = midInter.(int64)
	}
	c.JSON(spcSvc.FavNav(c, mid, v.VMid))
}

func favArc(c *bm.Context) {
	var mid int64
	v := new(model.FavArcArg)
	if err := c.Bind(v); err != nil {
		return
	}
	if midInter, ok := c.Get("mid"); ok {
		mid = midInter.(int64)
	}
	c.JSON(spcSvc.FavArchive(c, mid, v))
}
