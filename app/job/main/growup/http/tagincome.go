package http

import (
	"time"

	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func sendTagIncome(c *bm.Context) {
	log.Info("begin sendTagIncome")
	v := new(struct {
		Date string `form:"date" validate:"required"`
	})

	if err := c.Bind(v); err != nil {
		return
	}
	t, err := time.Parse("2006-01-02", v.Date)
	if err != nil {
		log.Error("sendTagIncome date error!date:%s", v.Date)
		return
	}
	err = svr.SendTagIncomeByHTTP(c, t.Year(), int(t.Month()), t.Day())
	if err != nil {
		log.Error("SendTagIncomeByHTTP error!(%v)", err)
	} else {
		log.Info("SendTagIncomeByHTTP succeed!")
	}
	c.JSON(nil, err)
}
