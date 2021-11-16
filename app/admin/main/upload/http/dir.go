package http

import (
	"github.com/namelessup/bilibili/app/admin/main/upload/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func addDir(c *bm.Context) {
	var err error
	adp := &model.AddDirParam{}
	if err = c.BindWith(adp, binding.FormPost); err != nil {
		return
	}

	c.JSON(nil, uaSvc.AddDir(c, adp))
}
