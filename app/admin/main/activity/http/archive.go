package http

import (
	"github.com/namelessup/bilibili/app/admin/main/activity/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func archives(c *bm.Context) {
	p := &model.ArchiveParam{}
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(actSrv.Archives(c, p))
}
