package http

import (
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func paramInsert(c *bm.Context) {
	v := new(struct {
		Name   string `form:"name"`
		Value  int    `form:"value"`
		Remark string `form:"remark"`
	})
	if err := c.Bind(v); err != nil {
		return
	}
	err := svr.InsertParameter(c, v.Name, v.Remark, v.Value)
	if err != nil {
		log.Error("svr.InsertParameter error(%v)", err)
	}
	c.JSON(nil, err)
}
