package dao

import (
	"time"

	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao .
type Dao struct {
	db         *sql.DB
	conf       *conf.Config
	client     *bm.Client
	redis      *redis.Pool
	mc         *memcache.Pool
	dbeiExpire int64
}

// New .
func New(c *conf.Config) *Dao {
	return &Dao{
		db:         sql.NewMySQL(c.Mysql),
		conf:       c,
		client:     bm.NewClient(c.HTTPClient),
		redis:      redis.NewPool(c.Redis.Config),
		dbeiExpire: int64(time.Duration(c.Cfg.Dangbei.Expire) / time.Second),
		mc:         memcache.NewPool(c.Memcache.Config),
	}
}

// Prom
var (
	errorsCount = prom.BusinessErrCount
	infosCount  = prom.BusinessInfoCount
)

// PromError prom error
func PromError(name string) {
	errorsCount.Incr(name)
}

// PromInfo add prom info
func PromInfo(name string) {
	infosCount.Incr(name)
}
