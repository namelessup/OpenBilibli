package http

import (
	"strconv"

	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/xstr"
)

// arcStat get archive stat.
func arcStat(c *bm.Context) {
	params := c.Request.Form
	aidStr := params.Get("aid")
	// check params
	aid, err := strconv.ParseInt(aidStr, 10, 64)
	if err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(arcSvc.Stat3(c, aid))
}

// arcStats get archives stat.
func arcStats(c *bm.Context) {
	params := c.Request.Form
	aidsStr := params.Get("aids")
	// check params
	aids, err := xstr.SplitInts(aidsStr)
	if err != nil {
		log.Error("query aids(%s) split error(%v)", aidsStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(arcSvc.Stats3(c, aids))
}
