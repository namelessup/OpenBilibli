package block

import (
	model "github.com/namelessup/bilibili/app/admin/main/member/model/block"
	service "github.com/namelessup/bilibili/app/admin/main/member/service/block"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"

	"github.com/pkg/errors"
)

var (
	svc *service.Service
)

// Setup is.
func Setup(blockSvc *service.Service, e *bm.Engine, authSvc *permit.Permit) {
	svc = blockSvc
	cb := e.Group("/x/admin/block", authSvc.Permit("BLOCK_SEARCH"))
	{
		cb.POST("/search", blockSearch)
		cb.GET("/history", history)
	}
	cb = e.Group("/x/admin/block", authSvc.Permit("BLOCK_BLOCK"))
	{
		cb.POST("", batchBlock)
	}
	cb = e.Group("/x/admin/block", authSvc.Permit("BLOCK_REMOVE"))
	{
		cb.POST("/remove", batchRemove)
	}
}

func bind(c *bm.Context, v model.ParamValidator) (err error) {
	if err = c.Bind(v); err != nil {
		err = errors.WithStack(err)
		return
	}
	if !v.Validate() {
		err = ecode.RequestErr
		c.JSON(nil, ecode.RequestErr)
		return
	}
	return
}
