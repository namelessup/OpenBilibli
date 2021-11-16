package ugc

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/job/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/sql"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao.
type Dao struct {
	conf     *conf.Config
	DB       *sql.DB
	client   *httpx.Client
	mc       *memcache.Pool
	mcExpire int32 // expire for ugc cache, same as pgc auth, because it's daily refresh
	criCID   int64 // critical cid for ugc video sync
	redis    *redis.Pool
}

// New create a instance of Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		conf:     c,
		DB:       sql.NewMySQL(c.Mysql),
		client:   httpx.NewClient(conf.Conf.HTTPClient),
		mc:       memcache.NewPool(c.Memcache.Config),
		mcExpire: int32(time.Duration(c.Memcache.Expire) / time.Second),
		criCID:   c.UgcSync.Cfg.CriticalCid,
		redis:    redis.NewPool(c.Redis.Config),
	}
	return
}

// Close close the redis and kafka resource.
func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}

// BeginTran begin mysql transaction
func (d *Dao) BeginTran(c context.Context) (*sql.Tx, error) {
	return d.DB.Begin(c)
}
