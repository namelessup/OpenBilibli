package caldiff

import (
	"github.com/namelessup/bilibili/app/job/main/appstatic/conf"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao .
type Dao struct {
	c      *conf.Config
	db     *xsql.DB
	client *bm.Client
}

// New creates a dao instance.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:      c,
		db:     xsql.NewMySQL(c.MySQL),
		client: bm.NewClient(c.HTTPClient),
	}
	return
}
