package tag

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/videoup/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is elec dao.
type Dao struct {
	c *conf.Config
	// http
	httpW *bm.Client
	// uri
	upBindURL   string
	TagCheckURL string
}

// New new a elec dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		// http client
		httpW: bm.NewClient(c.HTTPClient.Write),
		// uri
		upBindURL:   c.Host.Tag + _upBindURI,
		TagCheckURL: c.Host.Tag + _tagCheck,
	}
	return d
}

// Ping ping success.
func (d *Dao) Ping(c context.Context) (err error) {
	return
}
