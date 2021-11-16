package wechat

import (
	"github.com/namelessup/bilibili/app/interface/main/web-goblin/conf"
	"github.com/namelessup/bilibili/library/cache"
	"github.com/namelessup/bilibili/library/cache/redis"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao struct.
type Dao struct {
	// config
	c *conf.Config
	// redis
	redis *redis.Pool
	// httpClient
	httpClient *bm.Client
	// url
	wxAccessTokenURL string
	wxQrcodeURL      string
	cache            *cache.Cache
}

// New new dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// config
		c:          c,
		redis:      redis.NewPool(c.Redis.Config),
		httpClient: bm.NewClient(c.HTTPClient),
		cache:      cache.New(1, 1024),
	}
	d.wxAccessTokenURL = d.c.Host.Wechat + _accessTokenURI
	d.wxQrcodeURL = d.c.Host.Wechat + _qrcodeURI
	return
}
