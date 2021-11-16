package dao

import (
	"context"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf/paladin"
	"github.com/namelessup/bilibili/library/queue/databus"
)

// Dao dao.
type Dao struct {
	redis   *redis.Pool
	pushBus *databus.Databus
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New() (dao *Dao) {
	var (
		rds struct {
			Push *redis.Config
		}
		dbus struct {
			Push *databus.Config
		}
	)
	checkErr(paladin.Get("redis.toml").UnmarshalTOML(&rds))
	checkErr(paladin.Get("databus.toml").UnmarshalTOML(&dbus))
	dao = &Dao{
		redis:   redis.NewPool(rds.Push),
		pushBus: databus.New(dbus.Push),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
}

// Ping dao ping.
func (d *Dao) Ping(c context.Context) error {
	return d.pingRedis(c)
}
