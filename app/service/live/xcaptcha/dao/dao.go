package dao

import (
	"context"
	livecaptchaApi "github.com/namelessup/bilibili/app/service/live/captcha/api/liverpc"
	"github.com/namelessup/bilibili/app/service/live/xcaptcha/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
	"github.com/namelessup/bilibili/library/queue/databus"
)

// Dao dao
type Dao struct {
	c           *conf.Config
	redis       *redis.Pool
	geeClient   *GeeClient
	liveCaptcha *livecaptchaApi.Client
	captchaAnti *databus.Databus
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:           c,
		redis:       redis.NewPool(c.Redis),
		geeClient:   NewGeeClient(c.GeeTest),
		liveCaptcha: livecaptchaApi.New(getConf("captcha")),
		captchaAnti: databus.New(c.DataBus.CaptchaAnti),
	}
	return
}

// getConf get liveRpc conf
func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRpc
	if c != nil {
		return c[appName]
	}
	return nil
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return nil
}
