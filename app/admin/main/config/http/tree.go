package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func syncTree(c *bm.Context) {
	svr.SyncTree(c, user(c), c.Request.Header.Get("Cookie"))
	c.JSON(nil, nil)
}
