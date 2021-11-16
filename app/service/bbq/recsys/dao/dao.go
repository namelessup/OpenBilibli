package dao

import (
	"context"
	relation "github.com/namelessup/bilibili/app/service/main/relation/api"
	"time"

	recallv1 "github.com/namelessup/bilibili/app/service/bbq/recsys-recall/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/bbq/recsys/conf"
	searchv1 "github.com/namelessup/bilibili/app/service/bbq/search/api/grpc/v1"
	user "github.com/namelessup/bilibili/app/service/bbq/user/api"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/net/rpc/warden"

	"github.com/Dai0522/workpool"
)

// Dao dao
type Dao struct {
	c              *conf.Config
	mc             *memcache.Pool
	redis          *redis.Pool
	bfRedis        *redis.Pool
	db             *xsql.DB
	wp             *workpool.Pool
	SearchClient   searchv1.SearchClient
	RecallClient   recallv1.RecsysRecallClient
	UserClient     user.UserClient
	RelationClient relation.RelationClient
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	wpConf := &workpool.PoolConfig{
		MaxWorkers:     c.WorkerPool.MaxWorkers,
		MaxIdleWorkers: c.WorkerPool.MaxIdleWorkers,
		MinIdleWorkers: c.WorkerPool.MinIdleWorkers,
		KeepAlive:      time.Duration(c.WorkerPool.KeepAlive),
	}
	wp, err := workpool.NewWorkerPool(1024, wpConf)
	if err != nil {
		panic(err)
	}
	wp.Start()
	dao = &Dao{
		c:              c,
		redis:          redis.NewPool(c.Redis),
		bfRedis:        redis.NewPool(c.BFRedis),
		db:             xsql.NewMySQL(c.MySQL),
		wp:             wp,
		RecallClient:   newRecallClient(c.GRPCClient["recall"]),
		UserClient:     newUserClient(c.GRPCClient["user"]),
		RelationClient: newRelationClient(c.GRPCClient["relation"]),
	}
	return
}

func newRecallClient(cfg *conf.GRPCConfig) recallv1.RecsysRecallClient {
	cc, err := warden.NewClient(cfg.WardenConf).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return recallv1.NewRecsysRecallClient(cc)
}

func newUserClient(cfg *conf.GRPCConfig) user.UserClient {
	cc, err := warden.NewClient(cfg.WardenConf).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return user.NewUserClient(cc)
}

func newRelationClient(cfg *conf.GRPCConfig) relation.RelationClient {
	cc, err := warden.NewClient(cfg.WardenConf).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return relation.NewRelationClient(cc)
}

// Close close the resource.
func (d *Dao) Close() {
	d.mc.Close()
	d.redis.Close()
	d.bfRedis.Close()
	d.db.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	// TODO: if you need use mc,redis, please add
	return d.db.Ping(c)
}
