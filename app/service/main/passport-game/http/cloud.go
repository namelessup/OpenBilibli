package http

import (
	"github.com/namelessup/bilibili/library/net/http/blademaster"
)

func regions(c *blademaster.Context) {
	c.JSON(srv.Regions(c), nil)
}
