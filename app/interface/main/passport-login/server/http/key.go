package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func getKey(c *bm.Context) {
	c.JSON(srv.RSAKey(c), nil)
}
