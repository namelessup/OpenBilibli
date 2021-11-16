package ad

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/web-show/conf"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/stat/prom"
)

// Dao define db struct
type Dao struct {
	cpt *xsql.DB
	// sql
	selAdsStmt *xsql.Stmt
	// cpt
	httpClient *httpx.Client
	cpmURL     string
}

const (
	_cpmURL = "/bce/api/bce/pc"
)

// PromError err
func PromError(name string, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.Error(format, args...)
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		cpt:        xsql.NewMySQL(c.MySQL.Cpt),
		httpClient: httpx.NewClient(c.HTTPClient),
		cpmURL:     c.Host.Ad + _cpmURL,
	}
	dao.selAdsStmt = dao.cpt.Prepared(_selAds)
	return
}

// Close close the resource.
func (dao *Dao) Close() {
	dao.cpt.Close()
}

// Ping ping mysql
func (dao *Dao) Ping(c context.Context) error {
	return dao.cpt.Ping(c)
}
