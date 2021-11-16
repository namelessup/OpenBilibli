// Code generated by $GOPATH/src/github.com/namelessup/bilibili/app/tool/cache/gen. DO NOT EDIT.

/*
  Package roomNotice is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// cache: -sync=true -nullcache=&roomNotice.MonthConsume{Amount:-1} -check_null_code=$.Amount==-1
		MonthConsume(c context.Context, UID int64, targetID int64, date string) (*roomNotice.MonthConsume, error)
	}
*/

package notice

import (
	"context"

	"github.com/namelessup/bilibili/app/service/live/xuserex/model/roomNotice"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/stat/prom"
)

var _ _cache

// MonthConsume get data from cache if miss will call source method, then add to cache.
func (d *Dao) MonthConsume(c context.Context, id int64, targetID int64, date string) (res *roomNotice.MonthConsume, err error) {
	addCache := true
	res, err = d.CacheMonthConsume(c, id, targetID, date)

	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if nil != res && res.Amount == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("MonthConsume")
		return
	}
	prom.CacheMiss.Incr("MonthConsume")
	res, err = d.RawMonthConsume(c, id, targetID, date)
	log.Info("MonthConsume_RawMonthConsume uid (%v) targetId (%v) date (%v) res (%+v)", id, targetID, date, res)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &roomNotice.MonthConsume{Amount: -1}
	}
	if !addCache {
		return
	}
	d.AddCacheMonthConsume(c, id, targetID, date, miss)
	return
}
