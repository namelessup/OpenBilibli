package monitor

import (
	"context"
	"github.com/namelessup/bilibili/app/job/main/aegis/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

type Dao struct {
	c           *conf.Config
	redis       *redis.Pool
	db          *xsql.DB
	http        *bm.Client
	URLArcAddit string
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:           c,
		redis:       redis.NewPool(c.Redis),
		db:          xsql.NewMySQL(c.MySQL.Fast),
		http:        bm.NewClient(c.HTTP.Fast),
		URLArcAddit: c.Host.Videoup + _arcAdditURL,
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
	d.db.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	// TODO: if you need use mc,redis, please add
	return d.db.Ping(c)
}
