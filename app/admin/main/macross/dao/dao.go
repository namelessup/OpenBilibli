package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/macross/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

// Dao macross dao.
type Dao struct {
	// conf
	c *conf.Config
	// db
	db *sql.DB
}

// New dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: sql.NewMySQL(c.DB.Macross),
	}
	return
}

// Ping dao.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		log.Error("d.db error(%v)", err)
	}
	return
}

// Close close kafka connection.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
