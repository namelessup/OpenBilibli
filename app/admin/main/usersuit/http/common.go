package http

import (
	"github.com/namelessup/bilibili/app/admin/main/usersuit/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func httpData(c *bm.Context, data interface{}, pager *model.Pager) {
	res := make(map[string]interface{})
	if data == nil {
		data = struct{}{}
	}
	if pager == nil {
		pager = &model.Pager{}
	}
	res["data"] = data
	res["pager"] = &model.Pager{
		Total: pager.Total,
		PN:    pager.PN,
		PS:    pager.PS,
		Order: pager.Order,
		Sort:  pager.Sort,
	}
	c.JSONMap(res, nil)
}

func httpCode(c *bm.Context, err error) {
	c.JSON(nil, err)
}
