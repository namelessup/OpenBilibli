package dao

import (
	"github.com/namelessup/bilibili/app/service/main/msm/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao.
type Dao struct {
	client     *bm.Client
	db         *sql.DB
	treeHost   string
	platformID string
}

// New new dao.
func New(c *conf.Config) *Dao {
	d := &Dao{
		db:         sql.NewMySQL(c.Mysql),
		client:     bm.NewClient(c.HTTPClient),
		treeHost:   c.Tree.Host,
		platformID: c.Tree.PlatformID,
	}
	return d
}

// Close close mysql resource.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
