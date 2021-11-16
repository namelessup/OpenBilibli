package dao

import (
	"time"

	"github.com/namelessup/bilibili/app/service/main/secure/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/namelessup/bilibili/library/database/hbase.v2"
)

// Dao struct info of Dao.
type Dao struct {
	db                *sql.DB
	ddldb             *sql.DB
	c                 *conf.Config
	redis             *redis.Pool
	hbase             *hbase.Client
	locsExpire        int32
	expire            int64
	doubleCheckExpire int64
	mc                *memcache.Pool
	// http
	httpClient *bm.Client
}

// New new a Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:                 c,
		redis:             redis.NewPool(c.Redis.Config),
		expire:            int64(time.Duration(c.Redis.Expire) / time.Second),
		doubleCheckExpire: int64(time.Duration(c.Redis.DoubleCheck) / time.Second),
		db:                sql.NewMySQL(c.Mysql.Secure),
		ddldb:             sql.NewMySQL(c.Mysql.DDL),
		hbase:             hbase.NewClient(c.HBase.Config),
		mc:                memcache.NewPool(c.Memcache.Config),
		locsExpire:        int32(time.Duration(c.Memcache.Expire) / time.Second),
		httpClient:        bm.NewClient(c.HTTPClient),
	}
	return
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
