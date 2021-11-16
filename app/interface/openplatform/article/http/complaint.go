package http

import (
	"strconv"

	"github.com/namelessup/bilibili/app/interface/openplatform/article/conf"
	"github.com/namelessup/bilibili/app/interface/openplatform/article/dao"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
)

func addComplaint(c *bm.Context) {
	var (
		err               error
		mid, aid, cid     int64
		params            = c.Request.Form
		ip                = metadata.String(c, metadata.RemoteIP)
		reason, imageUrls string
	)
	midInter, _ := c.Get("mid")
	mid = midInter.(int64)
	aidStr := params.Get("aid")
	if aid, err = strconv.ParseInt(aidStr, 10, 64); err != nil || aid <= 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	cidStr := params.Get("cid")
	if cid, err = strconv.ParseInt(cidStr, 10, 64); err != nil || cid <= 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	reason = params.Get("reason")
	if int64(len([]rune(reason))) > conf.Conf.Article.MaxComplaintReasonLimit {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	imageUrls = params.Get("images")
	if err = artSrv.AddComplaint(c, aid, mid, cid, reason, imageUrls, ip); err != nil {
		dao.PromError("新增投诉")
		log.Error("artSrv.AddComplaint(%d,%d,%d, %s, %s) error(%+v)", mid, aid, cid, reason, imageUrls, err)
	}
	c.JSON(nil, err)
}
