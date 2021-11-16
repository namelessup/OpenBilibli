package http

import (
	"github.com/namelessup/bilibili/app/admin/main/apm/model/ut"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// utApps
func utApps(c *bm.Context) {
	var (
		err   error
		count int
		res   = []*ut.App{}
	)
	v := new(ut.AppReq)
	if err = c.Bind(v); err != nil {
		return
	}
	if res, count, err = apmSvc.UTApps(c, v); err != nil {
		c.JSON(nil, err)
		return
	}
	data := &Paper{
		Pn:    v.Pn,
		Ps:    v.Ps,
		Items: res,
		Total: count,
	}
	c.JSON(data, nil)
}
