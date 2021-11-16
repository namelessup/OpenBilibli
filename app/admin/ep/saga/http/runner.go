package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/saga/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func queryProjectRunners(c *bm.Context) {
	var (
		req = &model.ProjectDataReq{}
		err error
	)
	if err = c.Bind(req); err != nil {
		return
	}
	if req.Username, err = getUsername(c); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(srv.QueryProjectRunners(c, req))
}
