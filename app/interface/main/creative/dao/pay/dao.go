package pay

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao str
type Dao struct {
	c      *conf.Config
	client *bm.Client
	assURI string
	es     *elastic.Elastic
}

// New fn
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:      c,
		client: bm.NewClient(c.HTTPClient.Normal),
		assURI: c.Host.API + _assURI,
		es: elastic.NewElastic(&elastic.Config{
			Host:       c.Host.MainSearch,
			HTTPClient: c.HTTPClient.Slow,
		}),
	}
	return d
}

// Ping fn
func (d *Dao) Ping(c context.Context) (err error) {
	return
}
