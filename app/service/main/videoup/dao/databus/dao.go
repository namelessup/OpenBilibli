package databus

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/videoup/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/log"
)

// Dao is redis dao.
type Dao struct {
	c *conf.Config
	// redis
	redis *redis.Pool
}

//New  .
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:     c,
		redis: redis.NewPool(c.Redis.Track.Config),
	}
	return d
}

// Ping ping redis.
func (d *Dao) Ping(c context.Context) (err error) {
	conn := d.redis.Get(c)
	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
		log.Error("conn.Do(SET) error(%v)", err)
	}
	conn.Close()
	return
}

//Close .
func (d *Dao) Close() {
	d.redis.Close()
}
