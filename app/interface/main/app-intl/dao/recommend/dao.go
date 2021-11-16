package recommend

import (
	"time"

	"github.com/namelessup/bilibili/app/interface/main/app-intl/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is show dao.
type Dao struct {
	// http client
	client     *httpx.Client
	clientAsyn *httpx.Client
	// hetongzi
	hot string
	// bigdata
	rcmd  string
	group string
	top   string
	// redis
	redis     *redis.Pool
	expireRds int
	// mc
	mc       *memcache.Pool
	expireMc int32
}

// New new a show dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// http client
		client:     httpx.NewClient(c.HTTPData),
		clientAsyn: httpx.NewClient(c.HTTPClientAsyn),
		// hetongzi
		hot: c.Host.Hetongzi + _hot,
		// bigdata
		rcmd:  c.Host.Data + _rcmd,
		group: c.Host.BigData + _group,
		top:   c.Host.Data + _top,
		// redis
		redis:     redis.NewPool(c.Redis.Feed.Config),
		expireRds: int(time.Duration(c.Redis.Feed.ExpireRecommend) / time.Second),
		// mc
		mc:       memcache.NewPool(c.Memcache.Cache.Config),
		expireMc: int32(time.Duration(c.Memcache.Cache.Expire) / time.Second),
	}
	return
}

// Close close resource.
func (d *Dao) Close() {
	if d.redis != nil {
		d.redis.Close()
	}
	if d.mc != nil {
		d.mc.Close()
	}
}
