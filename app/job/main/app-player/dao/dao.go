package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/app-player/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
)

// Dao is dao.
type Dao struct {
	// mc
	mc *memcache.Pool
	// redis
	redis *redis.Pool
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// mc
		mc: memcache.NewPool(c.Memcache),
		// reids
		redis: redis.NewPool(c.Redis),
	}
	return
}

// PingMc is
func (d *Dao) PingMc(c context.Context) (err error) {
	conn := d.mc.Get(c)
	item := &memcache.Item{Key: "ping", Value: []byte{1}, Flags: memcache.FlagRAW, Expiration: 0}
	err = conn.Set(item)
	conn.Close()
	return
}
