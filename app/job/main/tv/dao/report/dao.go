package report

import (
	"github.com/namelessup/bilibili/app/job/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao .
type Dao struct {
	conf  *conf.Config
	httpR *bm.Client
	mc    *memcache.Pool
	DB    *sql.DB
}

// New create a instance of Dao and return .
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		conf:  c,
		httpR: bm.NewClient(c.DpClient),
		mc:    memcache.NewPool(c.Memcache.Config),
		DB:    sql.NewMySQL(c.Mysql),
	}
	return
}
