package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/activity/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	xhttp "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/jinzhu/gorm"
)

const (
	_actURLAddTags = "/x/internal/tag/activity/add"
	_songsURL      = "/x/internal/v1/audio/songs/activity/filter/info"
)

// Dao struct user of Dao.
type Dao struct {
	c             *conf.Config
	DB            *gorm.DB
	client        *xhttp.Client
	actURLAddTags string
	songsURL      string
}

// New create a instance of Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:             c,
		DB:            orm.NewMySQL(c.ORM),
		client:        xhttp.NewClient(c.HTTPClient),
		actURLAddTags: c.Host.API + _actURLAddTags,
		songsURL:      c.Host.API + _songsURL,
	}
	d.initORM()
	return
}

func (d *Dao) initORM() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if defaultTableName == "act_matchs" {
			return defaultTableName
		}
		return defaultTableName
	}
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
