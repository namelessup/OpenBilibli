package dao

import (
	"context"
	"strconv"

	"github.com/namelessup/bilibili/app/service/main/history/model"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// AddHistoryMessage .
func (d *Dao) AddHistoryMessage(c context.Context, k int, msg []*model.Merge) (err error) {
	key := strconv.Itoa(k)
	prom.BusinessInfoCount.Add("dbus-"+key, int64(len(msg)))
	if err = d.mergeDbus.Send(c, key, msg); err != nil {
		log.Error("Pub(%s,%+v) error(%v)", key, msg, err)
	}
	return
}
