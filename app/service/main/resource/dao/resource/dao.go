package resource

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/resource/conf"
	xsql "github.com/namelessup/bilibili/library/database/sql"
)

// Dao is resource dao.
type Dao struct {
	db *xsql.DB
	c  *conf.Config
}

// New init mysql db
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: xsql.NewMySQL(c.DB.Res),
	}
	return
}

// BeginTran begin transcation.
func (d *Dao) BeginTran(c context.Context) (tx *xsql.Tx, err error) {
	return d.db.Begin(c)
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
}

// Ping check dao health.
func (d *Dao) Ping(c context.Context) error {
	return d.db.Ping(c)
}
