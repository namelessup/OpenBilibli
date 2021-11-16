package pay

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/videoup/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao str
type Dao struct {
	c         *conf.Config
	client    *bm.Client
	assRegURI string
	assURI    string
	es        *elastic.Elastic
}

// New fn
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:         c,
		client:    bm.NewClient(c.HTTPClient.Write),
		assRegURI: c.Host.APICo + _assRegURI,
		assURI:    c.Host.APICo + _assURI,
		es: elastic.NewElastic(&elastic.Config{
			Host:       c.Host.MainSearch,
			HTTPClient: c.HTTPClient.Read,
		}),
	}
	return d
}

// Ping fn
func (d *Dao) Ping(c context.Context) (err error) {
	return
}
