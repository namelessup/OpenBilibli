package dede

import (
	"github.com/namelessup/bilibili/app/service/main/videoup/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
)

// Dao is redis dao.
type Dao struct {
	c *conf.Config
	// redis
	redis *redis.Pool
}

// New new
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:     c,
		redis: redis.NewPool(c.Redis.Track.Config),
	}
	return d
}
