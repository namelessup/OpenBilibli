package dao

import (
	"context"
	"fmt"

	arcGRPC "github.com/namelessup/bilibili/app/service/main/archive/api"
	"github.com/namelessup/bilibili/app/service/main/ugcpay/conf"
	"github.com/namelessup/bilibili/app/service/main/ugcpay/model"
	"github.com/namelessup/bilibili/library/cache"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	xsql "github.com/namelessup/bilibili/library/database/sql"
)

// Dao dao
type Dao struct {
	c          *conf.Config
	mc         *memcache.Pool
	redis      *redis.Pool
	db         *xsql.DB
	cache      *cache.Cache
	archiveAPI arcGRPC.ArchiveClient
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:     c,
		mc:    memcache.NewPool(c.Memcache),
		redis: redis.NewPool(c.Redis),
		db:    xsql.NewMySQL(c.MySQL),
		cache: cache.New(10, 10240),
	}
	var err error
	if dao.archiveAPI, err = arcGRPC.NewClient(nil); err != nil {
		panic(err)
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.mc.Close()
	d.redis.Close()
	d.db.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return d.db.Ping(c)
}

func orderKey(id string) string {
	return fmt.Sprintf("up_o_%s", id)
}

func assetKey(oid int64, otype string, currency string) string {
	return fmt.Sprintf("up_a_%d_%s_%s", oid, otype, currency)
}

func (d *Dao) cacheSFAsset(oid int64, otype string, currency string) string {
	return fmt.Sprintf("up_a_sf_%d_%s_%s", oid, otype, currency)
}

//go:generate $GOPATH/src/github.com/namelessup/bilibili/app/tool/cache/mc
type _mc interface {
	//mc: -key=orderKey -type=get
	CacheOrderUser(c context.Context, id int64) (*model.Order, error)
	//mc: -key=orderKey -expire=d.cacheTTL.OrderTTL
	AddCacheOrderUser(c context.Context, id int64, value *model.Order) error
	//mc: -key=orderKey
	DelCacheOrderUser(c context.Context, id int64) error

	//mc: -key=assetKey -type=get
	CacheAsset(c context.Context, oid int64, otype string, currency string) (*model.Asset, error)
	//mc: -key=assetKey -expire=d.cacheTTL.AssetTTL
	AddCacheAsset(c context.Context, oid int64, otype string, currency string, value *model.Asset) error
	//mc: -key=assetKey
	DelCacheAsset(c context.Context, oid int64, otype string, currency string) error
}

//go:generate $GOPATH/src/github.com/namelessup/bilibili/app/tool/cache/gen
type _cache interface {
	// cache: -nullcache=&model.Order{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	OrderUser(c context.Context, id int64) (*model.Order, error)
	// cache: -nullcache=&model.Asset{ID:-1} -check_null_code=$!=nil&&$.ID==-1 -singleflight=true
	Asset(c context.Context, oid int64, otype string, currency string) (*model.Asset, error)
}
