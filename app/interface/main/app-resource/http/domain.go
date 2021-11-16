package http

import bm "github.com/namelessup/bilibili/library/net/http/blademaster"

func domain(c *bm.Context) {
	c.JSON(domainSvc.Domain(), nil)
}
