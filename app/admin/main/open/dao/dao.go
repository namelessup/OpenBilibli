package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/open/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/log"

	"github.com/jinzhu/gorm"
)

// Dao .
type Dao struct {
	DB *gorm.DB
}

// New new a instance.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// db
		DB: orm.NewMySQL(c.ORM),
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
		if err = d.DB.DB().PingContext(c); err != nil {
			log.Error("d.PingContext error (%v)", err)
		}
	}
	return
}

// Close close connection of db , mc.
func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
