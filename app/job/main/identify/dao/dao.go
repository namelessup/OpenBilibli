package dao

import (
	"github.com/namelessup/bilibili/app/job/main/identify/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	xsql "github.com/namelessup/bilibili/library/database/sql"
)

// Dao dao
type Dao struct {
	c      *conf.Config
	authDB *xsql.DB
	authMC *memcache.Pool
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:      c,
		authDB: xsql.NewMySQL(c.AuthDB),
		authMC: memcache.NewPool(c.AuthMC),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.authDB.Close()
}
