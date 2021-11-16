package web

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/web-goblin/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/stat/prom"
)

const _broadURL = "/x/internal/broadcast/push/all"

// Dao dao
type Dao struct {
	c *conf.Config
	// http client
	http *bm.Client
	// broadcast URL
	broadcastURL string
	ela          *elastic.Elastic
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:            c,
		http:         bm.NewClient(c.HTTPClient),
		broadcastURL: c.Host.API + _broadURL,
		ela:          elastic.NewElastic(nil),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return nil
}

// PromError stat and log.
func PromError(name string, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.Error(format, args...)
}
