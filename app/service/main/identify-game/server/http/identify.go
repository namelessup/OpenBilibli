package http

import (
	"github.com/namelessup/bilibili/app/service/main/identify-game/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/main/identify-game/model"
	"github.com/namelessup/bilibili/app/service/main/identify-game/service"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func oauth(c *bm.Context) {
	var (
		data *model.AccessInfo
		err  error
	)
	req := c.Request
	accesskey := req.Form.Get("access_key")
	if accesskey == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	from := req.Form.Get("from")
	if data, err = srv.Oauth(c, accesskey, from); err != nil {
		if err == service.ErrDispatcherError {
			c.JSONMap(map[string]interface{}{"message": err.Error()}, err)
			return
		}
		c.JSON(nil, err)
		return
	}
	c.JSON(data, nil)
}

func renewToken(c *bm.Context) {
	var (
		data *model.RenewInfo
		err  error
	)
	req := c.Request
	accesskey := req.Form.Get("access_key")
	if accesskey == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	from := req.Form.Get("from")
	if data, err = srv.RenewToken(c, accesskey, from); err != nil {
		if err == service.ErrDispatcherError {
			c.JSONMap(map[string]interface{}{"message": err.Error()}, err)
			return
		}
		c.JSON(nil, err)
		return
	}
	c.JSON(data, nil)
}

func getCookieByToken(c *bm.Context) {
	var (
		data *v1.CreateCookieReply
		err error
	)
	p := new(v1.CreateCookieReq)
	if err = c.Bind(p); err != nil {
		log.Error("c.Bind err(%+v)", err)
		return
	}
	data, err = srv.GetCookieByToken(c, p.Token, p.From)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(data, nil)
}
