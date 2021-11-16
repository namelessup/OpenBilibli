package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/merlin/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func genMachinesV2(c *bm.Context) {
	var (
		gmr      = &model.GenMachinesRequest{}
		err      error
		username string
	)

	if username, err = getUsername(c); err != nil {
		return
	}

	if err = c.BindWith(gmr, binding.JSON); err != nil {
		return
	}

	if err = gmr.Verify(); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(nil, svc.GenMachinesV2(c, gmr, username))
}
