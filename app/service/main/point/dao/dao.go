package dao

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/service/main/point/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao dao
type Dao struct {
	c        *conf.Config
	mc       *memcache.Pool
	db       *xsql.DB
	mcExpire int32
	errProm  *prom.Prom
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		mc:       memcache.NewPool(c.Memcache.Config),
		mcExpire: int32(time.Duration(c.Memcache.Expire) / time.Second),
		db:       xsql.NewMySQL(c.MySQL),
		errProm:  prom.BusinessErrCount,
	}
	return
}

// Close close the resource.
func (dao *Dao) Close() {
	dao.mc.Close()
	dao.db.Close()
}

// Ping dao ping
func (dao *Dao) Ping(c context.Context) (err error) {
	if err = dao.db.Ping(c); err != nil {
		return
	}
	err = dao.pingMC(c)
	return
}

// pingMc ping
func (dao *Dao) pingMC(c context.Context) (err error) {
	conn := dao.mc.Get(c)
	defer conn.Close()
	item := memcache.Item{Key: "ping", Value: []byte{1}, Expiration: 60}
	return conn.Set(&item)
}
