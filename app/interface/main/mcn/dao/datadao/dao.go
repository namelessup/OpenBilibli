package datadao

import (
	"github.com/namelessup/bilibili/app/interface/main/mcn/conf"
	"github.com/namelessup/bilibili/app/interface/main/mcn/dao/global"
	"github.com/namelessup/bilibili/app/interface/main/mcn/tool/cache"
	"github.com/namelessup/bilibili/app/interface/main/mcn/tool/datacenter"
	"github.com/namelessup/bilibili/library/cache/memcache"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

//Dao data dao
type Dao struct {
	Client    *datacenter.HttpClient
	Conf      *conf.Config
	mc        *memcache.Pool
	McWrapper *cache.MCWrapper
	bmClient  *bm.Client
}

//New .
func New(c *conf.Config) *Dao {
	return &Dao{
		Client:    datacenter.New(c.DataClientConf),
		Conf:      c,
		mc:        global.GetMc(),
		McWrapper: cache.New(global.GetMc()),
		bmClient:  global.GetBMClient(),
	}
}
