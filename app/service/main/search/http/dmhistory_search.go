package http

import (
	"github.com/namelessup/bilibili/app/service/main/search/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func dmHistorySearch(c *bm.Context) {
	var (
		err    error
		params = c.Request.Form
		sp     = &model.DmHistoryParams{
			Bsp: &model.BasicSearchParams{},
		}
		res *model.SearchResult
	)
	if params.Get("appid") == "" || params.Get("oid") == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if err = c.Bind(sp); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if err = c.Bind(sp.Bsp); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	sp.Bsp.Source = []string{"id"}
	res, err = svr.DmHistory(c, sp)
	if err != nil {
		log.Error("srv.DmHistory(%v) error(%v)", sp, err)
		c.JSON(nil, ecode.ServerErr)
		return
	}
	c.JSON(res, err)
}
