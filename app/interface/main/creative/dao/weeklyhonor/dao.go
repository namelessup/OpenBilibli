package weeklyhonor

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	up "github.com/namelessup/bilibili/app/service/main/up/api/v1"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/hbase.v2"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

// Dao is data dao.
type Dao struct {
	c *conf.Config
	// hbase
	hbase        *hbase.Client
	hbaseTimeOut time.Duration
	// db
	db *sql.DB
	// mc
	mc            *memcache.Pool
	mcExpire      int32
	mcClickExpire int32
	// grpc
	upClient up.UpClient
}

// New init dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		// hbase
		hbase:        hbase.NewClient(c.HBaseOld.Config),
		hbaseTimeOut: time.Duration(time.Millisecond * 200),
		// db
		db: sql.NewMySQL(c.DB.Creative),
		// mc
		mc:            memcache.NewPool(c.Memcache.Honor.Config),
		mcExpire:      int32(time.Duration(c.Memcache.Honor.HonorExpire) / time.Second),
		mcClickExpire: int32(time.Duration(c.Memcache.Honor.ClickExpire) / time.Second),
	}
	var err error
	if d.upClient, err = up.NewClient(c.UpClient); err != nil {
		panic(err)
	}
	return
}

// Ping ping success.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.pingMySQL(c); err != nil {
		log.Error("s.pingMySQL.Ping err(%v)", err)
	}
	return
}

// Close hbase close
func (d *Dao) Close() (err error) {
	if d.hbase != nil {
		d.hbase.Close()
	}
	if d.db != nil {
		d.db.Close()
	}
	return
}
