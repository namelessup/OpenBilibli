package mission

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/videoup-report/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is mission dao.
type Dao struct {
	c          *conf.Config
	httpR      *bm.Client
	missAllURL string
}

// New new a mission dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		// client
		httpR: bm.NewClient(c.HTTPClient.Read),
		// uri
		missAllURL: c.Host.WWW + _msAllURL,
	}
	return d
}

// Ping ping success.
func (d *Dao) Ping(c context.Context) (err error) {
	return
}
