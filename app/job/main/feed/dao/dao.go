package dao

import (
	"context"
	"github.com/namelessup/bilibili/app/job/main/feed/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/stat/prom"
)

var (
	infosCount = prom.BusinessInfoCount
)

// Dao is feed job dao.
type Dao struct {
	c         *conf.Config
	smsClient *bm.Client
}

// New add a feed job dao.
func New(c *conf.Config) *Dao {
	return &Dao{
		c:         c,
		smsClient: bm.NewClient(c.HTTPClient),
	}
}

func (d *Dao) Ping(c context.Context) (err error) {
	return
}

func PromError(name string) {
	prom.BusinessErrCount.Incr(name)
}

// PromInfo add prom info
func PromInfo(name string) {
	infosCount.Incr(name)
}
