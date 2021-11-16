package dao

import (
	"context"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/log"

	"github.com/namelessup/bilibili/app/admin/live/live-admin/conf"

	relationApi "github.com/namelessup/bilibili/app/service/live/relation/api/liverpc"
)

// Dao dao
type Dao struct {
	c *conf.Config
	// mc    *memcache.Pool
	redis *redis.Pool
	// db    *xsql.DB
	Relation *relationApi.Client
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
		// mc:    memcache.NewPool(c.Memcache),
		redis: redis.NewPool(c.Redis),
		// db:    xsql.NewMySQL(c.MySQL),
		Relation: relationApi.New(getConf("relation")),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	// d.mc.Close()
	d.redis.Close()
	// d.db.Close()
}

// Ping dao ping
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.pingRedis(ctx); err != nil {
		log.Error("Failed to ping redis: %v", err)
	}
	return
}

func (d *Dao) pingRedis(ctx context.Context) (err error) {
	conn := d.redis.Get(ctx)
	_, err = conn.Do("SET", "PING", "PONG")
	conn.Close()
	return
}
