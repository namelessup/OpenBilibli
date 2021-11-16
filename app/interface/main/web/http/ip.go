package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func ipZone(c *bm.Context) {
	c.JSON(webSvc.IPZone(c))
}
