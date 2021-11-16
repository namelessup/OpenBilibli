package goblin

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao .
type Dao struct {
	conf   *conf.Config
	client *bm.Client
	db     *sql.DB
	mc     *memcache.Pool
}

// New .
func New(c *conf.Config) *Dao {
	return &Dao{
		conf:   c,
		client: bm.NewClient(c.PlayurlClient),
		db:     sql.NewMySQL(c.Mysql),
		mc:     memcache.NewPool(c.Memcache.Config),
	}
}
