package http

import (
	"strconv"

	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func userInfo(c *bm.Context) {
	var (
		mid int
		err error
	)
	midStr := c.Request.Form.Get("mid")
	if mid, err = strconv.Atoi(midStr); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	res, err := svc.UserInfo(c, int64(mid))
	c.JSON(res, err)
}
