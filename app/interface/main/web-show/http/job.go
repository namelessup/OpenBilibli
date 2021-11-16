package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// join
func join(c *bm.Context) {
	c.JSON(jobSvc.Jobs(c), nil)
}
