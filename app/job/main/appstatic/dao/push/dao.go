package push

import (
	appres "github.com/namelessup/bilibili/app/interface/main/app-resource/api/v1"
	"github.com/namelessup/bilibili/app/job/main/appstatic/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao .
type Dao struct {
	c         *conf.Config
	db        *xsql.DB
	client    *bm.Client
	redis     *redis.Pool
	appresCli appres.AppResourceClient
}

// New creates a dao instance.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:      c,
		db:     xsql.NewMySQL(c.MySQL),
		client: bm.NewClient(c.HTTPClient),
		redis:  redis.NewPool(c.Redis),
	}
	var err error
	if d.appresCli, err = appres.NewClient(c.AppresClient); err != nil {
		panic(err)
	}
	return
}
