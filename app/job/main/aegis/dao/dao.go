package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/aegis/conf"
	relrpc "github.com/namelessup/bilibili/app/service/main/relation/rpc/client"
	uprpc "github.com/namelessup/bilibili/app/service/main/up/api/v1"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/orm"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/jinzhu/gorm"
)

// Dao dao
type Dao struct {
	c      *conf.Config
	mc     *memcache.Pool
	redis  *redis.Pool
	slowdb *xsql.DB
	fastdb *xsql.DB
	orm    *gorm.DB
	//gorpc
	relRPC *relrpc.Service
	//grpc
	upRPC uprpc.UpClient

	httpFast *bm.Client
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:      c,
		mc:     memcache.NewPool(c.Memcache),
		redis:  redis.NewPool(c.Redis),
		fastdb: xsql.NewMySQL(c.MySQL.Fast),
		slowdb: xsql.NewMySQL(c.MySQL.Slow),
		orm:    orm.NewMySQL(c.ORM),

		httpFast: bm.NewClient(c.HTTP.Fast),
	}

	// rpc
	if !c.Debug {
		dao.relRPC = relrpc.New(c.RPC.Rel)
		var err error
		if dao.upRPC, err = uprpc.NewClient(c.GRPC.Up); err != nil {
			panic(err)
		}
	}

	dao.orm.LogMode(true)
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.mc.Close()
	d.redis.Close()
	d.slowdb.Close()
	d.fastdb.Close()
	d.orm.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return d.fastdb.Ping(c)
}
