package http

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/mcn/model"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
)

func arcTopDataStatistics(c *blademaster.Context) {
	httpGetWriterByExport(
		new(model.McnGetRankReq),
		func(cont context.Context, arg interface{}) (res interface{}, err error) {
			return srv.ArcTopDataStatistics(cont, arg.(*model.McnGetRankReq))
		},
		"ArcTopDataStatistics")(c)
}

func mcnsTotalDatas(c *blademaster.Context) {
	httpGetFunCheckCookie(
		new(model.TotalMcnDataReq),
		func(cont context.Context, arg interface{}) (res interface{}, err error) {
			return srv.McnsTotalDatas(cont, arg.(*model.TotalMcnDataReq))
		},
		"McnsTotalDatas")(c)
}

func mcnFansAnalyze(c *blademaster.Context) {
	httpGetFunCheckCookie(
		new(model.McnCommonReq),
		func(cont context.Context, arg interface{}) (res interface{}, err error) {
			return srv.McnFansAnalyze(cont, arg.(*model.McnCommonReq))
		},
		"McnFansAnalyze")(c)
}
