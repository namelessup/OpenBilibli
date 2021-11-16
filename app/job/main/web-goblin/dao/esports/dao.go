package esports

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/web-goblin/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

const _pushURL = "/x/internal/push-strategy/task/add"

// Dao dao
type Dao struct {
	c *conf.Config
	// http client
	http              *bm.Client
	messageHTTPClient *bm.Client
	// push service URL
	pushURL string
	// db
	db *sql.DB
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:                 c,
		http:              bm.NewClient(c.HTTPClient),
		messageHTTPClient: bm.NewClient(c.MessageHTTPClient),
		db:                sql.NewMySQL(c.DB.Esports),
		pushURL:           c.Host.API + _pushURL,
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
}

// Ping ping dao
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		return
	}
	return
}
