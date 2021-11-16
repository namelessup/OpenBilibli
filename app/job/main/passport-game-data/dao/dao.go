package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/passport-game-data/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

// Dao dao
type Dao struct {
	c       *conf.Config
	localDB *sql.DB
	cloudDB *sql.DB
}

// New new dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:       c,
		localDB: sql.NewMySQL(c.DB.Local),
		cloudDB: sql.NewMySQL(c.DB.Cloud),
	}
	return
}

// Ping check dao ok.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.localDB.Ping(c); err != nil {
		log.Info("dao.localDB.Ping() error(%v)", err)
	}
	if err = d.cloudDB.Ping(c); err != nil {
		log.Info("dao.cloudDB.Ping() error(%v)", err)
	}
	return
}

// Close close connections.
func (d *Dao) Close() (err error) {
	if d.localDB != nil {
		d.localDB.Close()
	}
	if d.cloudDB != nil {
		d.cloudDB.Close()
	}
	return
}
