package conf

import (
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	xtime "github.com/namelessup/bilibili/library/time"
)

// Redis redis
type Redis struct {
	*redis.Config
	Expire xtime.Duration
}

// Memcache config
type Memcache struct {
	*memcache.Config
	RelateExpire xtime.Duration
	ViewExpire   xtime.Duration
	ArcExpire    xtime.Duration
	CmsExpire    xtime.Duration
	HisExpire    xtime.Duration
	MangoExpire  xtime.Duration
}
