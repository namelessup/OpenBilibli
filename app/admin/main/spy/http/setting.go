package http

import (
	"github.com/namelessup/bilibili/app/admin/main/spy/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func settingList(c *bm.Context) {
	var (
		err  error
		data []*model.Setting
	)
	if data, err = spySrv.SettingList(c); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(data, err)
}

func updateSetting(c *bm.Context) {
	var (
		err    error
		params = c.Request.Form
		name   = params.Get("name")
		prop   = params.Get("prop")
		val    = params.Get("val")
	)
	if name == "" || prop == "" || val == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if err = spySrv.UpdateSetting(c, name, prop, val); err != nil {
		log.Error("s.UpdateSetting(%s,%s,%s) error(%v)", name, prop, val, err)
		c.JSON(nil, err)
		return
	}
	c.JSON(nil, nil)
}
