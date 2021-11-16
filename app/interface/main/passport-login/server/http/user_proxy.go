package http

import (
	"github.com/namelessup/bilibili/app/interface/main/passport-login/model"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func proxyCheckUserData(c *bm.Context) {
	param := new(model.ParamLogin)
	c.Bind(param)
	if param.UserName == "" || param.Pwd == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(srv.ProxyCheckUser(c, param))
}
