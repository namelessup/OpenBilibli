package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/saga/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// @params TasksReq
// @router get /ep/admin/saga/v1/tasks/project
// @response TasksResp
func projectTasks(ctx *bm.Context) {
	var (
		req = &model.TasksReq{}
		err error
	)
	if err = ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(srv.MergeTasks(ctx, req))
}
