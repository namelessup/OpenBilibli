package dao

import (
	"context"
	"github.com/namelessup/bilibili/app/service/bbq/topic/api"
	"github.com/namelessup/bilibili/library/sync/pipeline/fanout"
	"time"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf/paladin"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	xtime "github.com/namelessup/bilibili/library/time"
)

//go:generate $GOPATH/src/github.com/namelessup/bilibili/app/tool/cache/gen
type _cache interface {
	// cache: -sync=true -batch=10 -max_group=10 -batch_err=break -nullcache=&api.VideoExtension{Svid:-1} -check_null_code=$==nil||$.Svid==-1
	VideoExtension(c context.Context, ids []int64) (map[int64]*api.VideoExtension, error)
	// cache: -sync=true -batch=10 -max_group=10 -batch_err=break -nullcache=&api.TopicInfo{TopicId:-1} -check_null_code=$==nil||$.TopicId==-1
	TopicInfo(c context.Context, ids []int64) (map[int64]*api.TopicInfo, error)
}

// Dao dao.
type Dao struct {
	cache       *fanout.Fanout
	db          *sql.DB
	redis       *redis.Pool
	topicExpire int32
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New() (dao *Dao) {
	var (
		dc struct {
			Topic *sql.Config
		}
		rc struct {
			Topic       *redis.Config
			TopicExpire xtime.Duration
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	checkErr(paladin.Get("redis.toml").UnmarshalTOML(&rc))
	dao = &Dao{
		cache: fanout.New("cache", fanout.Worker(1), fanout.Buffer(1024)),
		// mysql
		db: sql.NewMySQL(dc.Topic),
		// redis
		redis:       redis.NewPool(rc.Topic),
		topicExpire: int32(time.Duration(rc.TopicExpire) / time.Second),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.pingRedis(ctx); err != nil {
		return
	}
	return d.db.Ping(ctx)
}

func (d *Dao) pingRedis(ctx context.Context) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
