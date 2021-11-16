package http

import (
	"github.com/namelessup/bilibili/app/job/main/account-summary/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func syncOne(ctx *bm.Context) {
	arg := &model.ArgMid{}
	if err := ctx.Bind(arg); err != nil {
		return
	}
	if err := srv.SyncOne(ctx, arg.Mid); err != nil {
		ctx.JSON(nil, err)
		return
	}

	ctx.JSON(srv.GetOne(ctx, arg.Mid))
}

func getOne(ctx *bm.Context) {
	arg := &model.ArgMid{}
	if err := ctx.Bind(arg); err != nil {
		return
	}
	ctx.JSON(srv.GetOne(ctx, arg.Mid))
}
