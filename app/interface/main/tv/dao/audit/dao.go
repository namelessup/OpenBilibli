package audit

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
)

// Dao is account dao.
type Dao struct {
	mc   *memcache.Pool
	conf *conf.Config
	db   *sql.DB
}

// New account dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		conf: c,
		mc:   memcache.NewPool(c.Memcache.Config),
		db:   sql.NewMySQL(c.Mysql),
	}
	return
}
