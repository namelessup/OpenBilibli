package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func queryCluster(c *bm.Context) {
	c.JSON(svc.QueryCluster(c))
}
