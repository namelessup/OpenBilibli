package dataplatform

import (
	"github.com/namelessup/bilibili/app/job/main/growup/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	xhttp "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is redis dao.
type Dao struct {
	c        *conf.Config
	db       *sql.DB
	url      string
	spyURL   string
	bgmURL   string
	basicURL string
	client   *xhttp.Client
}

// New is new redis dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: sql.NewMySQL(c.Mysql.Growup),
		// client
		client: xhttp.NewClient(c.DPClient),
		url:    c.Host.DataPlatform + "/avenger/api/38/query",
		spyURL: c.Host.DataPlatform + "/avenger/api/51/query",
		//		bgmURL: c.Host.DataPlatform + "/avenger/api/81/query",
		bgmURL:   c.Host.DataPlatform + "/avenger/api/95/query",
		basicURL: c.Host.DataPlatform + "/avenger/api/200/query",
	}
	return
}
