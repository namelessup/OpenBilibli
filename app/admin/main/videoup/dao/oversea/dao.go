package manager

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/namelessup/bilibili/app/admin/main/videoup/conf"
	"github.com/namelessup/bilibili/library/database/orm"
)

// Dao is redis dao.
type Dao struct {
	c *conf.Config
	// db
	OverseaDB *gorm.DB
}

var (
	d *Dao
)

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:         c,
		OverseaDB: orm.NewMySQL(c.DB.Oversea),
	}
	return d
}

// Close close.
func (d *Dao) Close() {
	if d.OverseaDB != nil {
		d.OverseaDB.Close()
	}
}

// Ping ping cpdb
func (d *Dao) Ping(c context.Context) (err error) {
	err = d.OverseaDB.DB().PingContext(c)
	return
}
