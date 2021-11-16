package dao

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/service/main/passport-auth/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
)

// Dao dao struct
type Dao struct {
	db       *sql.DB
	mc       *memcache.Pool
	mcExpire int32
	c        *conf.Config
}

// New create new dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:        c,
		db:       sql.NewMySQL(c.Mysql),
		mc:       memcache.NewPool(c.Memcache.Config),
		mcExpire: int32(time.Duration(c.Memcache.Expire) / time.Second),
	}
	return
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		return
	}
	return d.pingMC(c)
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.mc != nil {
		d.mc.Close()
	}
	if d.db != nil {
		d.db.Close()
	}
}
