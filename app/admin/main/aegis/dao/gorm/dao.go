package gorm

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/aegis/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/log"

	"github.com/jinzhu/gorm"
)

// Dao dao
type Dao struct {
	c   *conf.Config
	orm *gorm.DB
}

// New init mysql orm
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:   c,
		orm: orm.NewMySQL(c.ORM),
	}
	dao.orm.LogMode(true)

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.orm.Close()
}

// BeginTx .
func (d *Dao) BeginTx(c context.Context) (tx *gorm.DB, err error) {
	tx = d.orm.Begin()
	if err = tx.Error; err != nil {
		log.Error("orm begin tx error(%v)", err)
	}
	return
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return d.orm.DB().PingContext(c)
}
