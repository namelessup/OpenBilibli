package mission

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/videoup/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is mission dao.
type Dao struct {
	c                  *conf.Config
	httpR              *bm.Client
	missAllURL         string
	actOnlineByTypeURL string
}

// New new a mission dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		// client
		httpR: bm.NewClient(c.HTTPClient.Read),
		// uri
		missAllURL:         c.Host.WWW + _msAllURL,
		actOnlineByTypeURL: c.Host.WWW + _actOnlineByTypeURI,
	}
	return d
}

// Ping ping success.
func (d *Dao) Ping(c context.Context) (err error) {
	return
}
