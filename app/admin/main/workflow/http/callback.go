package http

import (
	"github.com/namelessup/bilibili/app/admin/main/workflow/model/param"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func listCallback(ctx *bm.Context) {
	ctx.JSON(wkfSvc.ListCallback(ctx))
}

func addOrUpCallback(ctx *bm.Context) {
	cbp := &param.AddCallbackParam{}
	if err := ctx.BindWith(cbp, binding.JSON); err != nil {
		return
	}

	if cbp.State > 0 {
		cbp.State = 1
	}

	cbID, err := wkfSvc.AddOrUpCallback(ctx, cbp)
	if err != nil {
		log.Error("wkfSvc.AddUpCallback(%+v) error(%v)", cbp, err)
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ctx.JSON(map[string]int32{
		"callbackNo": cbID,
	}, nil)
}
