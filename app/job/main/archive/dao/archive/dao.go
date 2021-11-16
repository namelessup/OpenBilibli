package archive

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/archive/conf"
	"github.com/namelessup/bilibili/library/database/sql"
)

// Dao is redis dao.
type Dao struct {
	c  *conf.Config
	db *sql.DB
}

// New is new redis dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: sql.NewMySQL(c.DB.Archive),
	}
	return d
}

// BeginTran begin transcation.
func (d *Dao) BeginTran(c context.Context) (tx *sql.Tx, err error) {
	return d.db.Begin(c)
}
