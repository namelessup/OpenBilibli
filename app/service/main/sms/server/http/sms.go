package http

import (
	pb "github.com/namelessup/bilibili/app/service/main/sms/api"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func send(ctx *bm.Context) {
	req := new(pb.SendReq)
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(smsSvc.Send(ctx, req))
}

func sendBatch(ctx *bm.Context) {
	req := new(pb.SendBatchReq)
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(smsSvc.SendBatch(ctx, req))
}
