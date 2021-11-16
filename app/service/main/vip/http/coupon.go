package http

import (
	"github.com/namelessup/bilibili/app/service/main/vip/model"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
)

func cancelUseCoupon(c *bm.Context) {
	var (
		err error
		arg = new(model.ArgCancelUseCoupon)
	)
	if err = c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, vipSvc.CancelUseCoupon(c, arg.Mid, arg.CouponToken, metadata.String(c, metadata.RemoteIP)))
}

func allowanceInfo(c *bm.Context) {
	var err error
	arg := new(model.ArgCancelUseCoupon)
	if err = c.Bind(arg); err != nil {
		log.Error("use allowance coupon bind %+v", err)
		return
	}
	c.JSON(vipSvc.AllowanceInfo(c, arg.Mid, arg.CouponToken))
}
