package dao

import (
	"context"
	"net/http"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/esports/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/sync/pipeline/fanout"
)

const (
	_searchURL = "/esports/search"
)

// Dao dao struct.
type Dao struct {
	// config
	c *conf.Config
	// db
	db *sql.DB
	// redis
	redis                    *redis.Pool
	filterExpire, listExpire int32
	// http client
	http      *bm.Client
	ldClient  *http.Client
	searchURL string
	ela       *elastic.Elastic
	cache     *fanout.Fanout
}

// New new dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// config
		c:            c,
		db:           sql.NewMySQL(c.Mysql),
		redis:        redis.NewPool(c.Redis.Config),
		filterExpire: int32(time.Duration(c.Redis.FilterExpire) / time.Second),
		listExpire:   int32(time.Duration(c.Redis.ListExpire) / time.Second),
		http:         bm.NewClient(c.HTTPClient),
		ldClient:     http.DefaultClient,
		searchURL:    c.Host.Search + _searchURL,
		ela:          elastic.NewElastic(nil),
		cache:        fanout.New("fanout"),
	}
	return
}

// Ping ping dao
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		return
	}
	return
}
