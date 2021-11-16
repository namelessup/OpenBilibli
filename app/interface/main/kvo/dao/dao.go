package dao

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/kvo/conf"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
)

// Dao kvo data access obj with bfs
type Dao struct {
	cache    *memcache.Pool
	mcExpire int32
	// http client for bfs req
	db *sql.DB
	// sql stmt
	getUserConf *sql.Stmt
	getDocument *sql.Stmt
}

// New new data access
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		cache:    memcache.NewPool(c.Memcache.Kvo),
		mcExpire: int32(time.Duration(c.Memcache.Expire) / time.Second),
		db:       sql.NewMySQL(c.Mysql),
	}
	d.getUserConf = d.db.Prepared(_getUserConf)
	d.getDocument = d.db.Prepared(_getDocument)
	return
}

// Ping check if health
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.pingMemcache(ctx); err != nil {
		return
	}
	if err = d.db.Ping(ctx); err != nil {
		return
	}
	return
}

// BeginTx begin trans
func (d *Dao) BeginTx(c context.Context) (*sql.Tx, error) {
	return d.db.Begin(c)
}
