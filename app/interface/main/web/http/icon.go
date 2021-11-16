package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func indexIcon(c *bm.Context) {
	c.JSON(webSvc.IndexIcon(), nil)
}
