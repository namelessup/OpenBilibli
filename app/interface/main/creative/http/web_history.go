package http

import (
	"strconv"

	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
)

func webHistoryList(c *bm.Context) {
	params := c.Request.Form
	aidStr := params.Get("aid")
	ip := metadata.String(c, metadata.RemoteIP)
	// check
	midI, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	mid, _ := midI.(int64)
	aid, err := strconv.ParseInt(aidStr, 10, 64)
	if err != nil || aid < 1 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	history, err := arcSvc.HistoryList(c, mid, aid, ip)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	c.JSON(history, nil)
}

func webHistoryView(c *bm.Context) {
	params := c.Request.Form
	hidStr := params.Get("history")
	ip := metadata.String(c, metadata.RemoteIP)
	// check
	midI, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	mid, _ := midI.(int64)
	hid, err := strconv.ParseInt(hidStr, 10, 64)
	if err != nil || hid < 1 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	history, err := arcSvc.HistoryView(c, mid, hid, ip)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(history, nil)
}
