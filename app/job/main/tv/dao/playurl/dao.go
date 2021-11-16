package playurl

import (
	"github.com/namelessup/bilibili/app/job/main/tv/conf"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao.
type Dao struct {
	conf   *conf.Config
	client *httpx.Client
}

// New create a instance of Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		conf:   c,
		client: httpx.NewClient(conf.Conf.HTTPClient),
	}
	return
}
