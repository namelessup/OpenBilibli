package medal

import (
	"github.com/namelessup/bilibili/app/job/main/usersuit/conf"
	"github.com/namelessup/bilibili/app/service/main/usersuit/rpc/client"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	_updateinfo = "/mingpai/api/updateinfo/%s"
)

// Dao struct info of Dao.
type Dao struct {
	db         *sql.DB
	c          *conf.Config
	client     *bm.Client
	suitRPC    *client.Service2
	updateInfo string
	// memcache
	mc *memcache.Pool
}

// New new a Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:          c,
		db:         sql.NewMySQL(c.Mysql),
		client:     bm.NewClient(c.HTTPClient),
		updateInfo: c.Properties.UpInfoURL + _updateinfo,
		suitRPC:    client.New(c.SuitRPC),
		// memcache
		mc: memcache.NewPool(c.Memcache.Config),
	}
	return
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
