package http

import bm "github.com/namelessup/bilibili/library/net/http/blademaster"

func sentinel(c *bm.Context) {
	c.JSON(map[string]interface{}{"sentinel": artSrv.Sentinel(c)}, nil)
}
