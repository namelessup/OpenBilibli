package http

import (
	"strconv"

	"github.com/namelessup/bilibili/app/admin/main/spy/model"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func report(c *bm.Context) {
	var (
		params = c.Request.Form
		pn, ps int
		err    error
		data   *model.ReportPage
	)
	if ps, err = strconv.Atoi(params.Get("ps")); err != nil {
		ps = model.DefPs
	}
	if pn, err = strconv.Atoi(params.Get("pn")); err != nil {
		pn = model.DefPn
	}
	data, err = spySrv.ReportList(c, ps, pn)
	if err != nil {
		log.Error("spySrv.ReportList error(%v)", err)
		c.JSON(nil, err)
		return
	}
	c.JSON(data, err)
}
