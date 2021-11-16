package http

import (
	kfcmdl "github.com/namelessup/bilibili/app/admin/main/activity/model/kfc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func kfcList(c *bm.Context) {
	arg := new(kfcmdl.ListParams)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(kfcSrv.List(c, arg))
}
