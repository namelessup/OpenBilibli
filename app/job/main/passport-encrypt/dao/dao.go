package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/passport-encrypt/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
)

// Dao dao
type Dao struct {
	c         *conf.Config
	originDB  *sql.DB
	encryptDB *sql.DB
}

// New new dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:         c,
		originDB:  sql.NewMySQL(c.DB.OriginDB),
		encryptDB: sql.NewMySQL(c.DB.EncryptDB),
	}
	return
}

// Ping check dao ok.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.originDB.Ping(c); err != nil {
		log.Info("dao.originDB.Ping() error(%v)", err)
	}
	if err = d.encryptDB.Ping(c); err != nil {
		log.Info("dao.encryptDB.Ping() error(%v)", err)
	}
	return
}

// Close close connections of mc, cloudDB.
func (d *Dao) Close() (err error) {
	if d.originDB != nil {
		d.originDB.Close()
	}
	if d.encryptDB != nil {
		d.encryptDB.Close()
	}
	return
}
