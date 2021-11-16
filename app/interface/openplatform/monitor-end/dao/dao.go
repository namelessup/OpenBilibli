package dao

import (
	"context"
	"github.com/namelessup/bilibili/app/interface/openplatform/monitor-end/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

// Dao .
type Dao struct {
	c     *conf.Config
	db    *sql.DB
	redis *redis.Pool
}

// New .
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:     c,
		db:    sql.NewMySQL(c.MySQL),
		redis: redis.NewPool(c.Redis),
	}
	return d
}

// Close .
func (d *Dao) Close() {
	d.db.Close()
	d.redis.Close()
}

// Ping .
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.pingRedis(c); err != nil {
		log.Error("d.pingRedis error(%+v)", err)
		return
	}
	if err = d.db.Ping(c); err != nil {
		log.Error("d.db.Ping error(%+v)", err)
	}
	return
}
