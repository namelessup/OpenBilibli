package http

import (
	artmdl "github.com/namelessup/bilibili/app/interface/openplatform/article/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
)

func categories(c *bm.Context) {
	data, err := artSrv.ListCategories(c, metadata.String(c, metadata.RemoteIP))
	if err != nil {
		c.JSON(nil, err)
		return
	}
	if data == nil {
		data = artmdl.Categories{}
	}
	c.JSON(data, nil)
}
