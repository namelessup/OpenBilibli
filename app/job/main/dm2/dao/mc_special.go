package dao

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/job/main/dm2/model"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_fmtSpecialDm = "s_special_%d_%d"
)

func (d *Dao) specialDmKey(oid int64, tp int32) string {
	return fmt.Sprintf(_fmtSpecialDm, oid, tp)
}

// DelSpecialDmCache .
func (d *Dao) DelSpecialDmCache(c context.Context, oid int64, tp int32) (err error) {
	var (
		key  = d.specialDmKey(oid, tp)
		conn = d.dmSegMC.Get(c)
	)
	defer conn.Close()
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
		} else {
			log.Error("memcache.Delete(%s) error(%v)", key, err)
		}
	}
	return
}

// AddSpecialDmCache add special content to memcache.
func (d *Dao) AddSpecialDmCache(c context.Context, ds *model.DmSpecial) (err error) {
	conn := d.dmSegMC.Get(c)
	key := d.specialDmKey(ds.Oid, ds.Type)
	defer conn.Close()
	item := &memcache.Item{
		Key:        key,
		Object:     ds,
		Flags:      memcache.FlagJSON,
		Expiration: d.dmSegMCExpire,
	}
	if err = conn.Set(item); err != nil {
		log.Error("conn.Set(%s) error(%v)", key, err)
	}
	return
}
