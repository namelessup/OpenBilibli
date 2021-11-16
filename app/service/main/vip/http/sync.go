package http

import (
	"github.com/namelessup/bilibili/app/service/main/vip/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func syncUser(c *bm.Context) {
	var err error
	user := new(model.VipUserInfo)
	if err = c.Bind(user); err != nil {
		return
	}
	vipSvc.SyncUser(c, user)
	c.JSON(nil, nil)
}
