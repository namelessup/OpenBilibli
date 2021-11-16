package dao

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/service/main/card/conf"
	"github.com/namelessup/bilibili/library/cache"
	"github.com/namelessup/bilibili/library/cache/memcache"
	xsql "github.com/namelessup/bilibili/library/database/sql"
)

// Dao dao
type Dao struct {
	c *conf.Config
	// memcache
	mc       *memcache.Pool
	mcExpire int32
	// db
	db *xsql.DB
	// cache async save
	cache *cache.Cache
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
		// card memcache
		mc:       memcache.NewPool(c.Memcache.Config),
		mcExpire: int32(time.Duration(c.Memcache.CardExpire) / time.Second),
		db:       xsql.NewMySQL(c.MySQL),
		// cache chan
		cache: cache.New(1, 1024),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.mc.Close()
	d.db.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.pingMC(c); err != nil {
		return
	}
	return d.db.Ping(c)
}
