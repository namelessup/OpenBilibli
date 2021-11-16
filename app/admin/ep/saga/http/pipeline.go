package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/saga/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// @params TeamDataRequest
// @router get /ep/admin/saga/v1/data/pipeline/report
// @response TeamDataResp
func queryTeamPipeline(c *bm.Context) {
	var (
		req = &model.TeamDataRequest{}
		err error
	)
	if err = c.Bind(req); err != nil {
		return
	}

	if req.Username, err = getUsername(c); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(srv.QueryTeamPipeline(c, req))
}

func queryProjectPipelineLists(c *bm.Context) {
	var (
		req = &model.PipelineDataReq{}
		err error
	)
	if err = c.Bind(req); err != nil {
		return
	}

	if req.Username, err = getUsername(c); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(srv.QueryProjectPipelineNew(c, req))
}
