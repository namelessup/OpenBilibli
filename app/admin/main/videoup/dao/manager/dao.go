package manager

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/videoup/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is redis dao.
type Dao struct {
	c *conf.Config
	// db
	managerDB  *sql.DB
	httpClient *bm.Client
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:          c,
		managerDB:  sql.NewMySQL(c.DB.Manager),
		httpClient: bm.NewClient(c.HTTPClient.Read),
	}
	return d
}

// Close close.
func (d *Dao) Close() {
	if d.managerDB != nil {
		d.managerDB.Close()
	}
}

// Ping ping cpdb
func (d *Dao) Ping(c context.Context) (err error) {
	return d.managerDB.Ping(c)
}
