package video

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/favorite/conf"
	xredis "github.com/namelessup/bilibili/library/cache/redis"
)

// Dao defeine fav Dao
type Dao struct {
	redisPool        *xredis.Pool
	expireRedis      int
	coverExpireRedis int
}

// New return fav dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		redisPool:        xredis.NewPool(c.Redis.Config),
		expireRedis:      int(time.Duration(c.Redis.Expire) / time.Second),
		coverExpireRedis: int(time.Duration(c.Redis.CoverExpire) / time.Second),
	}
	return
}

// Close close all connection
func (d *Dao) Close() {
	if d.redisPool != nil {
		d.redisPool.Close()
	}
}

// Ping check connection used in dao
func (d *Dao) Ping(c context.Context) (err error) {
	return d.pingRedis(c)
}
