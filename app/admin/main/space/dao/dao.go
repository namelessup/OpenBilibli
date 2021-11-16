package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/space/conf"
	"github.com/namelessup/bilibili/library/database/orm"

	"github.com/jinzhu/gorm"
)

// Dao .
type Dao struct {
	c  *conf.Config
	DB *gorm.DB
}

// New .
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// conf
		c: c,
		// db
		DB: orm.NewMySQL(c.ORM),
	}
	d.DB.LogMode(true)
	return
}

// Ping .
func (d *Dao) Ping(c context.Context) error {
	return d.DB.DB().PingContext(c)
}
