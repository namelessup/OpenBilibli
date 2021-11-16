package http

import bm "github.com/namelessup/bilibili/library/net/http/blademaster"

func whiteList(c *bm.Context) {
	c.JSON(whiteSvc.List(), nil)
}
