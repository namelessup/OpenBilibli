package http

import (
	"github.com/namelessup/bilibili/app/service/main/card/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func byMids(c *bm.Context) {
	var err error
	arg := new(model.ArgMids)
	if err = c.Bind(arg); err != nil {
		return
	}
	c.JSON(srv.UserCards(c, arg.Mids))
}
