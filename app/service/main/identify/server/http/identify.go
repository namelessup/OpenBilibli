package http

import (
	"github.com/namelessup/bilibili/app/service/main/identify/api/grpc"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// const (
// 	_actionChangePWD = "changePwd"
// 	_actionLoginOut  = "loginOut"
// )

func accessCookie(c *bm.Context) {
	cookie := c.Request.Header.Get("Cookie")
	if cookie == "" {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	res, err := srv.GetCookieInfo(c, cookie)
	if err == nil {
		c.Set("mid", res.Mid)
	}
	c.JSON(res, err)
}

func accessToken(c *bm.Context) {
	token := new(v1.GetTokenInfoReq)
	if err := c.Bind(token); err != nil {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	res, err := srv.GetTokenInfo(c, token)
	if err == nil {
		c.Set("mid", res.Mid)
	}
	c.JSON(res, err)
}

func delCache(c *bm.Context) {
	// query := c.Request.Form
	// action := query.Get("modifiedAttr")
	// if action != _actionChangePWD && action != _actionLoginOut {
	// 	return
	// }
	// key := query.Get("access_token")
	// if key == "" {
	// 	key = query.Get("session")
	// }
	// if key == "" {
	// 	return
	// }
	c.JSON(nil, nil)
}
