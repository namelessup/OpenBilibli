package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func cache(c *bm.Context) {
	c.JSONMap(srv.Cache(), nil)
}
