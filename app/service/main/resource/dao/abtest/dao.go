package abtest

import (
	"github.com/namelessup/bilibili/app/service/main/resource/conf"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao define db struct
type Dao struct {
	c *conf.Config
	// cpt
	httpClient *httpx.Client
	testURL    string
}

const (
	_abTestURL = "/abserver/v1/app/query-exp"
)

// New init mysql db
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:          c,
		httpClient: httpx.NewClient(c.HTTPClient),
		testURL:    c.Host.DataPlat + _abTestURL,
	}
	return
}
