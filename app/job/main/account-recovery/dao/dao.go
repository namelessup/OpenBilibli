package dao

import (
	"github.com/namelessup/bilibili/app/job/main/account-recovery/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao dao
type Dao struct {
	c *conf.Config
	// httpClient
	httpClient *bm.Client
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
		// httpClient
		httpClient: bm.NewClient(c.HTTPClientConfig),
	}
	return
}
