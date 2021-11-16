package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/up-rating/conf"

	"github.com/namelessup/bilibili/library/database/sql"
)

// Dao dao
type Dao struct {
	c  *conf.Config
	db *sql.DB
}

// New fn
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: sql.NewMySQL(c.DB.Rating),
	}
	return
}

// Ping ping health.
func (d *Dao) Ping(c context.Context) (err error) {
	return d.db.Ping(c)
}

// Close close connections of db.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
