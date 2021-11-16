package http

import (
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
	"strconv"
)

func addShare(c *bm.Context) {
	var (
		id     int64
		mid    int64
		params = c.Request.Form
	)
	idStr := params.Get("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	if id <= 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if midInter, ok := c.Get("mid"); ok {
		mid = midInter.(int64)
	}
	c.JSON(nil, artSrv.AddShare(c, id, mid, metadata.String(c, metadata.RemoteIP)))
}
