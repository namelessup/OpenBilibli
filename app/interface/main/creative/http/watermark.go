package http

import (
	"strconv"

	"github.com/namelessup/bilibili/app/interface/main/creative/model/watermark"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
)

func waterMark(c *bm.Context) {
	midI, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.CreativeNotLogin)
		return
	}
	mid, _ := midI.(int64)
	wm, err := wmSvc.WaterMark(c, mid)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(wm, nil)
}

func waterMarkSet(c *bm.Context) {
	params := c.Request.Form
	stStr := params.Get("state")
	tyStr := params.Get("type")
	posStr := params.Get("position")
	ip := metadata.String(c, metadata.RemoteIP)
	var (
		err         error
		wm          *watermark.Watermark
		ty, pos, st int64
	)
	midI, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.CreativeNotLogin)
		return
	}
	mid, _ := midI.(int64)
	if ty, err = strconv.ParseInt(tyStr, 10, 8); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", tyStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if pos, err = strconv.ParseInt(posStr, 10, 8); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", posStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if st, err = strconv.ParseInt(stStr, 10, 8); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", stStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	wm, err = wmSvc.WaterMarkSet(c, &watermark.WatermarkParam{
		MID:   mid,
		State: int8(st),
		Ty:    int8(ty),
		Pos:   int8(pos),
		IP:    ip,
	})
	if err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(wm, nil)
}

func waterMarkSetInternal(c *bm.Context) {
	params := c.Request.Form
	stStr := params.Get("state")
	tyStr := params.Get("type")
	posStr := params.Get("position")
	midStr := params.Get("mid")
	syncStr := params.Get("sync")
	ip := metadata.String(c, metadata.RemoteIP)
	var (
		err              error
		mid, ty, pos, st int64
		sync             int
	)
	mid, err = strconv.ParseInt(midStr, 10, 64)
	if err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", midStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	ty, err = strconv.ParseInt(tyStr, 10, 8)
	if err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", tyStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	pos, err = strconv.ParseInt(posStr, 10, 8)
	if err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", posStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	st, err = strconv.ParseInt(stStr, 10, 8)
	if err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", stStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	if syncStr != "" {
		sync, err = strconv.Atoi(syncStr)
		if err != nil {
			log.Error("strconv.ParseInt(%s) error(%v)", syncStr, err)
			c.JSON(nil, ecode.RequestErr)
			return
		}
	}

	wmSvc.AsyncWaterMarkSet(&watermark.WatermarkParam{
		MID:   mid,
		State: int8(st),
		Ty:    int8(ty),
		Pos:   int8(pos),
		Sync:  int8(sync),
		IP:    ip,
	})
	c.JSON(nil, nil)
}
