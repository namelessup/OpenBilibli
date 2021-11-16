package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/namelessup/bilibili/app/job/main/passport-game-cloud/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao
type Dao struct {
	c               *conf.Config
	getMemberStmt   []*sql.Stmt
	cloudDB         *sql.DB
	mc              *memcache.Pool
	mcExpire        int32
	gameClient      *bm.Client
	delGameCacheURI string
}

// New new dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:               c,
		cloudDB:         sql.NewMySQL(c.DB.Cloud),
		mc:              memcache.NewPool(c.Memcache.Config),
		mcExpire:        int32(time.Duration(c.Memcache.Expire) / time.Second),
		gameClient:      bm.NewClient(c.Game.Client),
		delGameCacheURI: c.Game.DelCacheURI,
	}
	d.getMemberStmt = make([]*sql.Stmt, _memberShard)
	for i := 0; i < _memberShard; i++ {
		d.getMemberStmt[i] = d.cloudDB.Prepared(fmt.Sprintf(_getMemberInfoSQL, i))
	}
	return
}

// Ping check dao ok.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.cloudDB.Ping(c); err != nil {
		log.Info("dao.cloudDB.Ping() error(%v)", err)
	}
	if err = d.pingMC(c); err != nil {
		log.Info("dao.pingMC() error(%v)", err)
	}
	return
}

// Close close connections of mc, cloudDB.
func (d *Dao) Close() (err error) {
	if d.cloudDB != nil {
		d.cloudDB.Close()
	}
	if d.mc != nil {
		d.mc.Close()
	}
	return
}
