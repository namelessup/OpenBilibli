package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/infra/databus/conf"
	"github.com/namelessup/bilibili/library/database/sql"
)

// Dao mysql struct
type Dao struct {
	db *sql.DB
}

// New new a Dao and return
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		db: sql.NewMySQL(c.MySQL),
	}
	return
}

// Ping ping mysql
func (d *Dao) Ping(c context.Context) error {
	return d.db.Ping(c)
}

// Close release mysql connection
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
