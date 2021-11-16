package http

import (
	v12 "github.com/namelessup/bilibili/app/service/live/xcaptcha/api/grpc/v1"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
)

// captchaVerify
func captchaVerify(ctx *blademaster.Context) {
	req := new(v12.XVerifyReq)
	if err := ctx.Bind(req); err != nil {
		return
	}
	resp, err := xCaptchaService.Verify(ctx, req)
	ctx.JSON(resp, err)
}
