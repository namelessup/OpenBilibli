package track

import (
	"time"

	"github.com/namelessup/bilibili/app/admin/main/videoup/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/sql"
)

// Dao is track dao.
type Dao struct {
	c *conf.Config
	// db
	db *sql.DB
	// redis
	redis  *redis.Pool
	expire int32
}

var (
	d *Dao
)

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:      c,
		db:     sql.NewMySQL(c.DB.Archive),
		redis:  redis.NewPool(c.Redis.Track.Config),
		expire: int32(time.Duration(c.Redis.Track.Expire) / time.Second),
	}
	return d
}
