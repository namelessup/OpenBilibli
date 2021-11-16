package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/playlist/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao dao struct.
type Dao struct {
	// config
	c *conf.Config
	// db
	db *sql.DB
	// redis
	redis      *redis.Pool
	statExpire int32
	plExpire   int32
	// http client
	http *bm.Client
	// stmt
	videosStmt map[string]*sql.Stmt
	// databus
	viewDbus  *databus.Databus
	shareDbus *databus.Databus
	// search video URL
	searchURL string
}

// New new dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// config
		c:          c,
		db:         sql.NewMySQL(c.Mysql),
		redis:      redis.NewPool(c.Redis.Config),
		statExpire: int32(time.Duration(c.Redis.StatExpire) / time.Second),
		plExpire:   int32(time.Duration(c.Redis.PlExpire) / time.Second),
		http:       bm.NewClient(c.HTTPClient),
		viewDbus:   databus.New(c.ViewDatabus),
		shareDbus:  databus.New(c.ShareDatabus),
		searchURL:  c.Host.Search + _searchURL,
	}
	d.videosStmt = make(map[string]*sql.Stmt, _plArcSub)
	for i := 0; i < _plArcSub; i++ {
		key := fmt.Sprintf("%02d", i)
		d.videosStmt[key] = d.db.Prepared(fmt.Sprintf(_plArcsSQL, key))
	}
	return
}

// Ping ping dao
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		return
	}
	err = d.pingRedis(c)
	return
}

func (d *Dao) pingRedis(c context.Context) (err error) {
	conn := d.redis.Get(c)
	_, err = conn.Do("SET", "PING", "PONG")
	conn.Close()
	return
}

// PromError stat and log.
func PromError(name string, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.Error(format, args...)
}
