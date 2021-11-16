package http

import (
	"github.com/namelessup/bilibili/app/interface/main/creative/model/order"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
	"strconv"
)

func webCmOasisStat(c *bm.Context) {
	ip := metadata.String(c, metadata.RemoteIP)
	midI, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	mid, _ := midI.(int64)
	var (
		oasisEarnings *order.OasisEarnings
	)
	oa, _ := arcSvc.Oasis(c, mid, ip)
	oasisEarnings = &order.OasisEarnings{}
	if oa != nil {
		oasisEarnings.State = oa.State
		oasisEarnings.Realese = oa.RealeseOrder
		oasisEarnings.Total = oa.TotalOrder
	}
	c.JSON(oasisEarnings, nil)
}

func arcCommercial(c *bm.Context) {
	ip := metadata.String(c, metadata.RemoteIP)
	params := c.Request.Form
	aidStr := params.Get("aid")
	aid, err := strconv.ParseInt(aidStr, 10, 64)
	if err != nil || aid <= 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	pd, err := arcSvc.ArcCommercial(c, aid, ip)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	if pd == nil || pd.GameID == 0 {
		c.JSON(nil, ecode.NothingFound)
		return
	}
	c.JSON(pd, nil)
}
