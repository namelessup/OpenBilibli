// Code generated by $GOPATH/src/github.com/namelessup/bilibili/app/tool/cache/gen. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// cache: -batch=50 -max_group=10 -batch_err=continue -nullcache=&model.UserEquip{CardID:-1} -check_null_code=$!=nil&&$.CardID==-1
		Equips(c context.Context, keys []int64) (map[int64]*model.UserEquip, error)
		// cache: -nullcache=&model.UserEquip{CardID:-1} -check_null_code=$!=nil&&$.CardID==-1 -singleflight=true
		Equip(c context.Context, key int64) (*model.UserEquip, error)
	}
*/

package dao

import (
	"context"
	"sync"

	"github.com/namelessup/bilibili/app/service/main/card/model"
	"github.com/namelessup/bilibili/library/net/metadata"
	"github.com/namelessup/bilibili/library/stat/prom"
	"github.com/namelessup/bilibili/library/sync/errgroup"

	"golang.org/x/sync/singleflight"
)

var _ _cache
var cacheSingleFlights = [1]*singleflight.Group{{}}

// Equips get data from cache if miss will call source method, then add to cache.
func (d *Dao) Equips(c context.Context, keys []int64) (res map[int64]*model.UserEquip, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	res, err = d.CacheEquips(c, keys)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("Equips", int64(len(keys)-len(miss)))
	defer func() {
		for k, v := range res {
			if v != nil && v.CardID == -1 {
				delete(res, k)
			}
		}
	}()
	if len(miss) == 0 {
		return
	}
	var missData map[int64]*model.UserEquip
	missLen := len(miss)
	prom.CacheMiss.Add("Equips", int64(missLen))
	mutex := sync.Mutex{}
	for i := 0; i < missLen; i += 50 * 10 {
		var subKeys []int64
		group := &errgroup.Group{}
		ctx := c
		if (i + 50*10) > missLen {
			subKeys = miss[i:]
		} else {
			subKeys = miss[i : i+50*10]
		}
		missSubLen := len(subKeys)
		for j := 0; j < missSubLen; j += 50 {
			var ks []int64
			if (j + 50) > missSubLen {
				ks = subKeys[j:]
			} else {
				ks = subKeys[j : j+50]
			}
			group.Go(func() (err error) {
				data, err := d.RawEquips(ctx, ks)
				mutex.Lock()
				for k, v := range data {
					if missData == nil {
						missData = make(map[int64]*model.UserEquip, len(keys))
					}
					missData[k] = v
				}
				mutex.Unlock()
				return
			})
		}
		err1 := group.Wait()
		if err1 != nil {
			err = err1
		}
	}
	if res == nil {
		res = make(map[int64]*model.UserEquip)
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	for _, key := range keys {
		if res[key] == nil {
			if missData == nil {
				missData = make(map[int64]*model.UserEquip, len(keys))
			}
			missData[key] = &model.UserEquip{CardID: -1}
		}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheEquips(metadata.WithContext(c), missData)
	})
	return
}

// Equip get data from cache if miss will call source method, then add to cache.
func (d *Dao) Equip(c context.Context, id int64) (res *model.UserEquip, err error) {
	addCache := true
	res, err = d.CacheEquip(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.CardID == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("Equip")
		return
	}
	var rr interface{}
	sf := d.cacheSFEquip(id)
	rr, err, _ = cacheSingleFlights[0].Do(sf, func() (r interface{}, e error) {
		prom.CacheMiss.Incr("Equip")
		r, e = d.RawEquip(c, id)
		return
	})
	res = rr.(*model.UserEquip)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.UserEquip{CardID: -1}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheEquip(metadata.WithContext(c), id, miss)
	})
	return
}
