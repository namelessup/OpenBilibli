package dao

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/job/main/passport-sns/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

// Dao dao
type Dao struct {
	c        *conf.Config
	asoDB    *xsql.DB
	snsDB    *xsql.DB
	mc       *memcache.Pool
	mcExpire int32
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		asoDB:    xsql.NewMySQL(c.DB.Aso),
		snsDB:    xsql.NewMySQL(c.DB.Sns),
		mc:       memcache.NewPool(c.Memcache.Config),
		mcExpire: int32(time.Duration(c.Memcache.Expire) / time.Second),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() (err error) {
	if err = d.asoDB.Close(); err != nil {
		log.Error("srv.asoDB.Close() error(%v)", err)
	}
	if err = d.snsDB.Close(); err != nil {
		log.Error("srv.snsDB.Close() error(%v)", err)
	}
	return
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return nil
}

// BeginSnsTran begin sns transcation.
func (d *Dao) BeginSnsTran(c context.Context) (tx *sql.Tx, err error) {
	return d.snsDB.Begin(c)
}
