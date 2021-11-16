package bnj

import (
	"time"

	"github.com/namelessup/bilibili/app/interface/main/activity/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	xhttp "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao bnj dao.
type Dao struct {
	c              *conf.Config
	mc             *memcache.Pool
	redis          *redis.Pool
	client         *xhttp.Client
	resetExpire    int32
	rewardExpire   int32
	grantCouponURL string
}

// New init bnj dao.
func New(c *conf.Config) *Dao {
	d := &Dao{
		c:            c,
		mc:           memcache.NewPool(c.Memcache.Like),
		redis:        redis.NewPool(c.Redis.Config),
		client:       xhttp.NewClient(c.HTTPClientBnj),
		resetExpire:  int32(time.Duration(c.Redis.ResetExpire) / time.Second),
		rewardExpire: int32(time.Duration(c.Redis.RewardExpire) / time.Second),
	}
	d.grantCouponURL = d.c.Host.Mall + _grantCouponURL
	return d
}

// Close .
func (d *Dao) Close() {
	if d.mc != nil {
		d.mc.Close()
	}
	if d.redis != nil {
		d.redis.Close()
	}
}
