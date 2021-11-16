package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/infra/notify/conf"
	xsql "github.com/namelessup/bilibili/library/database/sql"
)

// Dao dao
type Dao struct {
	c  *conf.Config
	db *xsql.DB
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:  c,
		db: xsql.NewMySQL(c.MySQL),
	}
	return
}

// Close close the resource.
func (dao *Dao) Close() {
	dao.db.Close()
}

// Ping dao ping
func (dao *Dao) Ping(c context.Context) error {
	return dao.db.Ping(c)
}
