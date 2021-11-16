package http

import (
	"github.com/namelessup/bilibili/app/interface/main/creative/model/archive"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"strconv"
)

// videoViewPoints get video highlight viewpoints
func videoViewPoints(c *bm.Context) {
	var (
		aid, cid int64
		err      error
		form     = c.Request.Form
		vp       *archive.ViewPointRow
	)
	if aid, err = strconv.ParseInt(form.Get("aid"), 10, 64); err != nil || aid <= 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if cid, err = strconv.ParseInt(form.Get("cid"), 10, 64); err != nil || cid <= 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if vp, err = arcSvc.VideoPoints(c, aid, cid); err != nil {
		log.Error("arcSvc.VideoPoints(%d,%d) error(%v)", aid, cid, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(vp, nil)
}
