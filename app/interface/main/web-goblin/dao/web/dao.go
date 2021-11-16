package web

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/web-goblin/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/stat/prom"
)

const (
	_pgcFullURL  = "/ext/internal/archive/channel/content"
	_pgcIncreURL = "/ext/internal/archive/channel/content/change"
)

// Dao dao .
type Dao struct {
	c                       *conf.Config
	db                      *sql.DB
	showDB                  *sql.DB
	httpR                   *bm.Client
	pgcFullURL, pgcIncreURL string
	ela                     *elastic.Elastic
}

// New init mysql db .
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:           c,
		db:          sql.NewMySQL(c.DB.Goblin),
		showDB:      sql.NewMySQL(c.DB.Show),
		httpR:       bm.NewClient(c.SearchClient),
		pgcFullURL:  c.Host.PgcURI + _pgcFullURL,
		pgcIncreURL: c.Host.PgcURI + _pgcIncreURL,
		ela:         elastic.NewElastic(c.Es),
	}
	return
}

// Close close the resource .
func (d *Dao) Close() {
}

// Ping dao ping .
func (d *Dao) Ping(c context.Context) error {
	return nil
}

// PromError stat and log .
func PromError(name string, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.Error(format, args...)
}
