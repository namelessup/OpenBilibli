package http

import (
	"github.com/namelessup/bilibili/app/service/main/passport/model"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func historyPwdCheck(c *bm.Context) {
	param := new(model.HistoryPwdCheckParam)
	c.Bind(param)
	if param.Mid <= 0 || param.Pwd == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(passportSvc.HistoryPwdCheck(c, param))
}
