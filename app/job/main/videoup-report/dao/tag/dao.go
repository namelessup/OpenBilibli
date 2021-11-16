package tag

import (
	tagClient "github.com/namelessup/bilibili/app/interface/main/tag/rpc/client"
	"github.com/namelessup/bilibili/app/job/main/videoup-report/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

//Dao tag dao
type Dao struct {
	c                       *conf.Config
	client                  *bm.Client
	upBindURL, adminBindURL string
	tagDisRPC               *tagClient.Service
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:            c,
		client:       bm.NewClient(c.HTTPClient.Write),
		upBindURL:    c.Host.API + _upBindURI,
		adminBindURL: c.Host.API + _adminBindURI,
		tagDisRPC:    tagClient.New2(c.TagDisConf),
	}
	return d
}
