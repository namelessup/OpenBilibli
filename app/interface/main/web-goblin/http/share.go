package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func encourage(c *bm.Context) {
	var (
		mid int64
	)
	midStr, _ := c.Get("mid")
	mid = midStr.(int64)
	c.JSON(srvShare.Encourage(c, mid))
}
