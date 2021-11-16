package dao

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/job/main/vip/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao struct info of Dao.
type Dao struct {
	// mysql
	db    *sql.DB
	oldDb *sql.DB
	// http
	client *bm.Client
	// conf
	c *conf.Config
	// memcache
	mc       *memcache.Pool
	mcExpire int32
	//redis pool
	redis        *redis.Pool
	redisExpire  int32
	errProm      *prom.Prom
	frozenExpire int32
}

// New new a Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// conf
		c: c,
		// mc
		mc:       memcache.NewPool(c.Memcache.Config),
		mcExpire: int32(time.Duration(c.Memcache.Expire) / time.Second),
		// redis
		redis:       redis.NewPool(c.Redis.Config),
		redisExpire: int32(time.Duration(c.Redis.Expire) / time.Second),
		// db
		db:    sql.NewMySQL(c.NewMysql),
		oldDb: sql.NewMySQL(c.OldMysql),
		// http client
		client:       bm.NewClient(c.HTTPClient),
		errProm:      prom.BusinessErrCount,
		frozenExpire: int32(time.Duration(c.Property.FrozenExpire) / time.Second),
	}
	return
}

// Ping ping health of db.
func (d *Dao) Ping(c context.Context) (err error) {
	return d.db.Ping(c)
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
	if d.redis != nil {
		d.redis.Close()
	}
}

//StartTx start tx
func (d *Dao) StartTx(c context.Context) (tx *sql.Tx, err error) {
	if d.db != nil {
		tx, err = d.db.Begin(c)
	}
	return
}
