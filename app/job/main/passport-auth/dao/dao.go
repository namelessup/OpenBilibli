package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/passport-auth/conf"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao
type Dao struct {
	c     *conf.Config
	db    *xsql.DB
	olddb *xsql.DB
	// httpClient
	httpClient *bm.Client
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:     c,
		db:    xsql.NewMySQL(c.MySQL),
		olddb: xsql.NewMySQL(c.OldMySQL),
		// httpClient
		httpClient: bm.NewClient(c.HTTPClientConfig),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
	d.olddb.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return d.pingMC(c)
}

// pingMc ping
func (d *Dao) pingMC(c context.Context) (err error) {
	return
}
