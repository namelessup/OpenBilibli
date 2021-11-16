package operation

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/web-show/conf"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao struct
type Dao struct {
	db *xsql.DB
}

// PromError err
func PromError(name string, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.Error(format, args...)
}

// New Conf
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{db: xsql.NewMySQL(c.MySQL.Operation)}
	return
}

// Close Dao
func (dao *Dao) Close() {
	dao.db.Close()
}

// Ping Dao
func (dao *Dao) Ping(c context.Context) error {
	return dao.db.Ping(c)
}
