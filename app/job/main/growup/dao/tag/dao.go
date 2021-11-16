package tag

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/growup/conf"

	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is redis dao.
type Dao struct {
	c            *conf.Config
	db           *sql.DB
	client       *bm.Client
	archiveURL   string
	typeURL      string
	columnURL    string
	columnActURL string
}

// New is new redis dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:            c,
		db:           sql.NewMySQL(c.Mysql.Growup),
		client:       bm.NewClient(c.HTTPClient),
		archiveURL:   c.Host.Archive + "/manager/search",
		typeURL:      c.Host.VideoType + "/videoup/types",
		columnURL:    c.Host.ColumnType,
		columnActURL: c.Host.ColumnAct,
	}
	return
}

// Ping ping health.
func (d *Dao) Ping(c context.Context) (err error) {
	return d.db.Ping(c)
}

// Close close connections
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

// BeginTran begin transcation
func (d *Dao) BeginTran(c context.Context) (tx *sql.Tx, err error) {
	return d.db.Begin(c)
}
