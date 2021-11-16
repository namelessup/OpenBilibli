package http

import (
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func expireUpRating(c *bm.Context) {
	arg := new(struct {
		MID int64 `form:"mid"`
	})
	if err := c.Bind(arg); err != nil {
		log.Error("error bind arg")
		c.JSON(nil, ecode.RequestErr)
		return
	}
	mid := arg.MID
	err := svc.ExpireUpRatingCache(c, mid)
	if err != nil {
		log.Error("svc.ExpireUpRatingCache mid(%v) err(%v)", mid, err)
		c.JSON(nil, err)
		return
	}
	c.JSON(true, err)
}
