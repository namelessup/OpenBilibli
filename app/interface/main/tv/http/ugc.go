package http

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

//ugcPlayurl is used for getting ugc play url
func ugcPlayurl(c *bm.Context) {
	var (
		err error
		mid int64
	)
	param := new(model.PlayURLReq)
	if err = c.Bind(param); err != nil {
		return
	}
	if param.AccessKey != "" {
		if cmid, ok := c.Get("mid"); ok {
			mid = cmid.(int64)
		}
	}
	c.JSONMap(gobSvc.UgcPlayurl(c, param, mid))
}
