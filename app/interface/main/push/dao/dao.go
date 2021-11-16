package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/push/conf"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao .
type Dao struct {
	c           *conf.Config
	reportPub   *databus.Databus
	callbackPub *databus.Databus
}

// New creates a push-service DAO instance.
func New(c *conf.Config) *Dao {
	d := &Dao{
		c:           c,
		reportPub:   databus.New(c.ReportPub),
		callbackPub: databus.New(c.CallbackPub),
	}
	return d
}

// PromError prom error
func PromError(name string) {
	prom.BusinessErrCount.Incr(name)
}

// PromInfo add prom info
func PromInfo(name string) {
	prom.BusinessInfoCount.Incr(name)
}

// Close dao.
func (d *Dao) Close() {}

// Ping check connection status.
func (d *Dao) Ping(c context.Context) (err error) {
	return
}
