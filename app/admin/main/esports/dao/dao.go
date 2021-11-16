package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/esports/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/database/orm"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/jinzhu/gorm"
)

const _esports = "esports"

// Dao .
type Dao struct {
	c       *conf.Config
	DB      *gorm.DB
	Elastic *elastic.Elastic
	// client
	replyClient *bm.Client
}

// New .
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		// conf
		c: c,
		// db
		DB: orm.NewMySQL(c.ORM),
		// elastic
		Elastic:     elastic.NewElastic(nil),
		replyClient: bm.NewClient(c.HTTPReply),
	}
	return
}

// Ping .
func (d *Dao) Ping(c context.Context) error {
	return d.DB.DB().PingContext(c)
}
