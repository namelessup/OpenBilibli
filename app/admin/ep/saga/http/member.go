package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/saga/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func queryProjectMembers(c *bm.Context) {
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
	c.JSON(srv.QueryProjectMembers(c, req))
}
