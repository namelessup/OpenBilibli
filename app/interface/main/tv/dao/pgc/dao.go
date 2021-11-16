package pgc

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is account dao.
type Dao struct {
	conf   *conf.Config
	client *bm.Client
	mc     *memcache.Pool
}

// New account dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		conf:   c,
		client: bm.NewClient(c.HTTPClient),
		mc:     memcache.NewPool(c.Memcache.Config),
	}
	return
}
