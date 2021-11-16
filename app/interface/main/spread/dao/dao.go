package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/spread/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao dao
type Dao struct {
	c          *conf.Config
	httpClient *bm.Client
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:          c,
		httpClient: bm.NewClient(c.HTTPClient),
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

// PromError .
func PromError(name string) {
	prom.BusinessErrCount.Incr(name)
}
