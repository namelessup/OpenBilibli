package http

import (
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"
	"strconv"
)

func upSwitch(c *bm.Context) {
	ip := metadata.String(c, metadata.RemoteIP)
	// check user
	midI, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	mid, _ := midI.(int64)
	rs, err := upSvc.UpSwitch(c, mid, ip)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	rs.Show = 1 //播放开关全量
	c.JSON(rs, nil)
}

func setUpSwitch(c *bm.Context) {
	ip := metadata.String(c, metadata.RemoteIP)
	params := c.Request.Form
	stateStr := params.Get("state")
	fromStr := params.Get("from")
	midI, ok := c.Get("mid")
	if !ok {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	state, err := strconv.Atoi(stateStr)
	if err != nil {
		log.Error("strconv.Atoi(%s) error(%v)", stateStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	from, err := strconv.Atoi(fromStr)
	if err != nil {
		log.Error("strconv.Atoi(%s) error(%v)", fromStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	mid, _ := midI.(int64)
	id, err := upSvc.SetUpSwitch(c, mid, state, from, ip)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(map[string]interface{}{
		"id": id,
	}, nil)
}
