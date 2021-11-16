package dao

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/job/live-userexp/model"
	mc "github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_expKey = "level:%d"
)

func key(uid int64) string {
	return fmt.Sprintf(_expKey, uid)
}

// SetLevelCache 设置等级缓存
func (d *Dao) SetLevelCache(c context.Context, level *model.Level) (err error) {
	key := key(level.Uid)
	conn := d.expMc.Get(c)
	defer conn.Close()

	if conn.Set(&mc.Item{
		Key:        key,
		Object:     level,
		Flags:      mc.FlagProtobuf,
		Expiration: d.cacheExpire,
	}); err != nil {
		log.Error("[dao.mc_exp|SetLevelCache] conn.Set(%s, %v) error(%v)", key, level, err)
	}
	return
}
