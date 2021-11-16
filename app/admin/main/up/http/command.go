package http

import (
	"context"
	"github.com/namelessup/bilibili/app/admin/main/up/model"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
)

func commandRefreshUpRank(c *blademaster.Context) {
	httpQueryFunc(new(model.CommandCommonArg),
		func(context context.Context, arg interface{}) (res interface{}, err error) {
			return Svc.Crmservice.CommandRefreshUpRank(context, arg.(*model.CommandCommonArg))
		},
		"SignCheckTask")(c)
}
