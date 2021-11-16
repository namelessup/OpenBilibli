package manager

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/videoup-report/conf"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

// Dao is manager dao.
type Dao struct {
	c  *conf.Config
	db *xsql.DB
}

// New new a manager dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: xsql.NewMySQL(c.DB.Manager),
	}
	return d
}

//Ping ping
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		log.Error("manager ping error(%v)", err)
	}
	return
}

//Close close
func (d *Dao) Close() (err error) {
	if d.db != nil {
		err = d.db.Close()
	}
	if err != nil {
		log.Error("manager close error(%v)", err)
	}
	return
}
