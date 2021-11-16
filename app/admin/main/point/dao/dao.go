package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/point/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/elastic"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

const _searchBussinss = "vip_point_change_history"

// Dao dao
type Dao struct {
	c      *conf.Config
	mc     *memcache.Pool
	db     *xsql.DB
	client *bm.Client
	es     *elastic.Elastic
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:      c,
		mc:     memcache.NewPool(c.Memcache),
		db:     xsql.NewMySQL(c.MySQL),
		client: bm.NewClient(c.HTTPClient),
		// es
		es: elastic.NewElastic(nil),
	}
	return
}

// Close close the resource.
func (dao *Dao) Close() {
	dao.mc.Close()
	dao.db.Close()
}

// Ping dao ping
func (dao *Dao) Ping(c context.Context) error {
	return dao.pingMC(c)
}

// pingMc ping
func (dao *Dao) pingMC(c context.Context) (err error) {
	conn := dao.mc.Get(c)
	defer conn.Close()
	return
}
