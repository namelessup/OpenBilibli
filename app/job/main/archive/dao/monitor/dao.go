package monitor

import (
	"github.com/namelessup/bilibili/app/job/main/archive/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is redis dao.
type Dao struct {
	c      *conf.Config
	client *bm.Client
}

// New is new redis dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:      c,
		client: bm.NewClient(c.HTTPClient),
	}
	return d
}
