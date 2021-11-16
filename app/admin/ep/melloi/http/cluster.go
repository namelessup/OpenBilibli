package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

//ClusterInfo get cluster infomation
func ClusterInfo(c *bm.Context) {
	c.JSON(srv.ClusterInfo(c))
}
