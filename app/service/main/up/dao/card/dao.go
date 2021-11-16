package card

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/up/conf"
	"github.com/namelessup/bilibili/app/service/main/up/dao/global"
	"github.com/namelessup/bilibili/library/database/sql"
)

// Dao is redis dao.
type Dao struct {
	c  *conf.Config
	db *sql.DB
}

// New fn
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: global.GetUpCrmDB(),
	}
	return d
}

// Close fn
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

// Ping ping cpdb
func (d *Dao) Ping(c context.Context) (err error) {
	return d.db.Ping(c)
}
