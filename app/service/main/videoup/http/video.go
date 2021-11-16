package http

import (
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func obtainCid(c *bm.Context) {
	params := c.Request.Form
	fn := params.Get("filename")
	if fn == "" {
		log.Error("filename not exist")
		c.JSON(nil, ecode.NothingFound)
		return
	}
	cid, err := vdpSvc.ObtainCid(c, fn)
	if err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(map[string]int64{
		"cid": cid,
	}, nil)
}

func queryCid(c *bm.Context) {
	params := c.Request.Form
	fn := params.Get("filename")
	if fn == "" {
		log.Error("filename not exist")
		c.JSON(nil, ecode.NothingFound)
		return
	}
	cid, err := vdpSvc.FindCidByFn(c, fn)
	if err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(map[string]int64{
		"cid": cid,
	}, nil)
}
