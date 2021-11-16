package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/coupon/conf"
	seasongrpc "github.com/namelessup/bilibili/app/service/openplatform/pgc-season/api/grpc/season/v1"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/elastic"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao
type Dao struct {
	c  *conf.Config
	db *xsql.DB
	es *elastic.Elastic
	mc *memcache.Pool
	// grpc
	rpcClient seasongrpc.SeasonClient
	client    *bm.Client
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:  c,
		db: xsql.NewMySQL(c.MySQL),
		// es
		es:     elastic.NewElastic(nil),
		mc:     memcache.NewPool(c.Memcache.Config),
		client: bm.NewClient(c.HTTPClient),
	}
	var err error
	if dao.rpcClient, err = seasongrpc.NewClient(c.PGCRPC); err != nil {
		log.Error("seasongrpc NewClientt error(%v)", err)
	}
	return
}

// Close close the resource.
func (dao *Dao) Close() {
	dao.db.Close()
}

// Ping dao ping
func (dao *Dao) Ping(c context.Context) error {
	return dao.db.Ping(c)
}
