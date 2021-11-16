package dao

import (
	"context"

	"github.com/namelessup/bilibili/library/log"

	"github.com/namelessup/bilibili/app/job/live/xlottery/internal/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	xsql "github.com/namelessup/bilibili/library/database/sql"
)

// Dao dao
type Dao struct {
	c     *conf.Config
	redis *redis.Pool
	db    *xsql.DB
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:     c,
		redis: redis.NewPool(c.Redis.Lottery),
		db:    xsql.NewMySQL(c.Database.Lottery),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
	d.db.Close()
}

// Ping dao ping
func (d *Dao) Ping(ctx context.Context) error {
	// TODO: add mc,redis... if you use
	return d.db.Ping(ctx)
}

func (d *Dao) execSqlWithBindParams(c context.Context, sql *string, bindParams ...interface{}) (affect int64, err error) {
	res, err := d.db.Exec(c, *sql, bindParams...)
	if err != nil {
		log.Error("db.Exec(%s) error(%v)", *sql, err)
		return
	}
	return res.RowsAffected()
}
