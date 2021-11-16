package http

import (
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func joinColumn(c *bm.Context) {
	v := new(struct {
		AccountType int `form:"account_type"`
		SignType    int `form:"sign_type"`
	})
	if err := c.Bind(v); err != nil {
		return
	}

	midI, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	mid, _ := midI.(int64)
	err := svc.JoinColumn(c, mid, v.AccountType, v.SignType)
	if err != nil {
		log.Error("svc.JoinColumn mid(%d) accountType(%d) signType(%d) error(%v)", mid, v.AccountType, v.SignType, err)
	}
	c.JSON(nil, err)
}
