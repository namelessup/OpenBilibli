package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/cache/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/jinzhu/gorm"
)

// Dao dao.
type Dao struct {
	c      *conf.Config
	DB     *gorm.DB
	client *bm.Client
}

// New new a dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:      c,
		DB:     orm.NewMySQL(c.MySQL),
		client: bm.NewClient(c.HTTPClient),
	}
	return
}

// Ping check connection of db , mc.
func (d *Dao) Ping(c context.Context) (err error) {
	return
}

// Close close connection of db , mc.
func (d *Dao) Close() {

}
