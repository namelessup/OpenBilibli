package http

import (
	"time"

	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func updateCheat(c *bm.Context) {
	log.Info("begin updateCheat")
	v := new(struct {
		Date string `form:"date" validate:"required"`
	})

	if err := c.Bind(v); err != nil {
		return
	}
	t, err := time.Parse("2006-01-02", v.Date)
	if err != nil {
		log.Error("updateCheat date error!date:%s", v.Date)
		return
	}
	err = svr.UpdateCheatHTTP(c, t)
	if err != nil {
		log.Error("Exec UpdateCheat error!(%v)", err)
	} else {
		log.Info("Exec UpdateCheat succeed!")
	}
	c.JSON(nil, err)
}
