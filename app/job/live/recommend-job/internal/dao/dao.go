package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/live/recommend-job/internal/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
)

// Dao dao
type Dao struct {
	c     *conf.Config
	redis *redis.Pool
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:     c,
		redis: redis.NewPool(c.Redis),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
}

// Ping dao ping
func (d *Dao) Ping(ctx context.Context) error {
	// TODO: add mc,redis... if you use
	c := d.redis.Get(ctx)
	defer c.Close()
	_, err := c.Do("ping")
	return err
}
