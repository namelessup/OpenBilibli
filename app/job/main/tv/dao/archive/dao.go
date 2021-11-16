package archive

import (
	"time"

	"github.com/namelessup/bilibili/app/job/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
)

// Dao is archive dao.
type Dao struct {
	// memcache
	mc         *memcache.Pool
	expireView int32
}

// New new a archive dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// memcache
		mc:         memcache.NewPool(c.Memcache.Config),
		expireView: int32(time.Duration(c.Memcache.ExpireMedia) / time.Second),
	}
	return
}
