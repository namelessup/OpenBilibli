package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func regions(c *bm.Context) {
	c.JSON(srv.Regions, nil)
}
