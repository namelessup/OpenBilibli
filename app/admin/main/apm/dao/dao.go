package dao

import (
	"github.com/namelessup/bilibili/app/admin/main/apm/conf"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/orm"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/jinzhu/gorm"
)

// Dao dao.
type Dao struct {
	c         *conf.Config
	DB        *gorm.DB
	DBDatabus *gorm.DB
	DBCanal   *gorm.DB
	// client
	client *bm.Client
	Redis  *redis.Pool
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:         c,
		DB:        orm.NewMySQL(c.ORM),
		DBDatabus: orm.NewMySQL(c.ORMDatabus),
		DBCanal:   orm.NewMySQL(c.ORMCanal),
		client:    bm.NewClient(c.HTTPClient),
		Redis:     redis.NewPool(c.Redis.Config),
	}
	d.initORM()
	return
}

func (d *Dao) initORM() {
	d.DB.LogMode(true)
	d.DBDatabus.LogMode(true)
	d.DBCanal.LogMode(true)
}

// Close close connection of db , mc.
func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
	if d.DBDatabus != nil {
		d.DBDatabus.Close()
	}
	if d.DBCanal != nil {
		d.DBCanal.Close()
	}
}
