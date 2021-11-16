package upper

import (
	"runtime"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/stat/prom"
)

var (
	missedCount = prom.CacheMiss
	cachedCount = prom.CacheHit
)

// Dao is account dao.
type Dao struct {
	mc       *memcache.Pool
	mcExpire int32
	DB       *sql.DB
	mCh      chan func()
}

// New account dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		mc:       memcache.NewPool(c.Memcache.Config),
		mcExpire: int32(time.Duration(c.Memcache.CmsExpire) / time.Second),
		DB:       sql.NewMySQL(c.Mysql),
		mCh:      make(chan func(), 10240),
	}
	for i := 0; i < runtime.NumCPU()*2; i++ {
		go d.cacheproc()
	}
	return
}

// addCache add archive to mc or redis
func (d *Dao) addCache(f func()) {
	select {
	case d.mCh <- f:
	default:
		log.Warn("cacheproc chan full")
	}
}

// cacheproc write memcache and stat redis use goroutine
func (d *Dao) cacheproc() {
	for {
		f := <-d.mCh
		f()
	}
}
