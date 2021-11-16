package http

import (
	"strconv"

	"github.com/namelessup/bilibili/app/interface/main/web/conf"
	artmdl "github.com/namelessup/bilibili/app/interface/openplatform/article/model"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/xstr"
)

func articleList(c *bm.Context) {
	var (
		rid, mid     int64
		pn, ps, sort int
		aids         []int64
		err          error
	)
	param := c.Request.Form
	pnStr := param.Get("pn")
	psStr := param.Get("ps")
	ridStr := param.Get("rid")
	aidsStr := param.Get("aids")
	sortStr := param.Get("sort")
	if midInter, ok := c.Get("mid"); ok {
		mid = midInter.(int64)
	}
	if pn, err = strconv.Atoi(pnStr); err != nil || pn < 1 {
		pn = 1
	}
	if ps, err = strconv.Atoi(psStr); err != nil || ps < 1 || ps > conf.Conf.Rule.MaxArtPageSize {
		ps = conf.Conf.Rule.MaxArtPageSize
	}
	if ridStr != "" {
		if rid, err = strconv.ParseInt(ridStr, 10, 64); err != nil || rid < 0 {
			c.JSON(nil, ecode.RequestErr)
			return
		}
	}
	if aidsStr != "" {
		if aids, err = xstr.SplitInts(aidsStr); err != nil {
			c.JSON(nil, ecode.RequestErr)
			return
		}
	}
	if sortStr != "" {
		if sort, err = strconv.Atoi(sortStr); err != nil || sort < 0 {
			c.JSON(nil, ecode.RequestErr)
			return
		}
		sortCheck := false
		for _, v := range artmdl.SortFields {
			if sort == v {
				sortCheck = true
				break
			}
		}
		if !sortCheck && sort != 0 {
			c.JSON(nil, ecode.RequestErr)
			return
		}
	}
	c.JSON(webSvc.ArticleList(c, rid, mid, sort, pn, ps, aids))
}

func articleUpList(c *bm.Context) {
	var mid int64
	if midInter, ok := c.Get("mid"); ok {
		mid = midInter.(int64)
	}
	c.JSON(webSvc.ArticleUpList(c, mid))
}

func categories(c *bm.Context) {
	c.JSON(webSvc.Categories(c))
}

func newCount(c *bm.Context) {
	var (
		count, pubTime int64
		err            error
	)
	pubTimeStr := c.Request.Form.Get("pubtime")
	if pubTime, err = strconv.ParseInt(pubTimeStr, 10, 64); err != nil || pubTime < 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if count, err = webSvc.NewCount(c, pubTime); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(struct {
		NewCount int64 `json:"new_count"`
	}{NewCount: count}, nil)
}

func upMoreArts(c *bm.Context) {
	var (
		aid int64
		err error
	)
	aidStr := c.Request.Form.Get("aid")
	if aid, err = strconv.ParseInt(aidStr, 10, 64); err != nil || aid <= 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(webSvc.UpMoreArts(c, aid))
}
