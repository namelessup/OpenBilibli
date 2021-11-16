package music

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/videoup/conf"
	"github.com/namelessup/bilibili/library/database/orm"

	"github.com/jinzhu/gorm"
)

// Dao struct user of Dao.
type Dao struct {
	c *conf.Config
	// db
	DB *gorm.DB
}

var (
	d *Dao
)

// New create a instance of Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// conf
		c: c,
		// db
		DB: orm.NewMySQL(c.ORMArchive),
	}
	d.initORM()
	return
}

func (d *Dao) initORM() {
	d.DB.LogMode(true)
}

// Ping check connection of db , mc.
func (d *Dao) Ping(c context.Context) (err error) {
	if d.DB != nil {
		err = d.DB.DB().PingContext(c)
	}
	return
}

// Close close connection of db , mc.
func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
