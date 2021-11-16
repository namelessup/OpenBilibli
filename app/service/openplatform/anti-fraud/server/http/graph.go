package http

import (
	"github.com/namelessup/bilibili/app/service/openplatform/anti-fraud/api/grpc/v1"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

//graphPrepare 拉起图形验证
func graphPrepare(c *bm.Context) {
	params := new(v1.GraphPrepareRequest)
	if err := c.Bind(params); err != nil {
		return
	}
	c.JSON(svc.GraphPrepare(c, params))
}

//graphCheck 图形验证
func graphCheck(c *bm.Context) {
	params := new(v1.GraphCheckRequest)
	if err := c.Bind(params); err != nil {
		return
	}
	c.JSON(svc.GraphCheck(c, params))
}
