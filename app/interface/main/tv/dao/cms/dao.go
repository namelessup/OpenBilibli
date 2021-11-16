package cms

import (
	"runtime"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao is account dao.
type Dao struct {
	mc        *memcache.Pool
	conf      *conf.Config
	db        *sql.DB
	mCh       chan func()
	expireCMS int32
}

// New account dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		conf:      c,
		mc:        memcache.NewPool(c.Memcache.Config),
		db:        sql.NewMySQL(c.Mysql),
		mCh:       make(chan func(), 10240),
		expireCMS: int32(time.Duration(c.Memcache.CmsExpire) / time.Second),
	}
	// video db
	for i := 0; i < runtime.NumCPU()*2; i++ {
		go d.cacheproc()
	}
	return
}

// Prom
var (
	errorsCount = prom.BusinessErrCount
	infosCount  = prom.BusinessInfoCount
	cachedCount = prom.CacheHit
	missedCount = prom.CacheMiss
)

// PromError prom error
func PromError(name string) {
	errorsCount.Incr(name)
}

// PromInfo add prom info
func PromInfo(name string) {
	infosCount.Incr(name)
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
