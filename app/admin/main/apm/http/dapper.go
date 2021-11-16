package http

import (
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func dapperProxy(c *bm.Context) {
	apmSvc.DapperProxy(c.Writer, c.Request)
}
