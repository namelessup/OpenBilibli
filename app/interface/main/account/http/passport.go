package http

import (
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func testUserName(c *bm.Context) {
	mid, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	name := c.Request.Form.Get("name")
	if name == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, passSvc.TestUserName(c, name, mid.(int64)))
}
