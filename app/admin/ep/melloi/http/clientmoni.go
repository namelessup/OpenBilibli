package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/melloi/model"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/binding"
)

func queryClientMoni(c *bm.Context) {

	var (
		cli    model.ClientMoni
		resMap = make(map[string]interface{})
	)
	if err := c.BindWith(&cli, binding.Form); err != nil {
		c.JSON(nil, err)
		return
	}
	clientMonis, err := srv.QueryClientMoni(&cli)
	if err != nil {
		log.Error("srv.QueryClientMoni err (%v)", err)
		return
	}
	resMap["clientMonis"] = clientMonis
	c.JSON(resMap, err)
}
