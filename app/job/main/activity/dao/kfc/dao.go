package kfc

import (
	"github.com/namelessup/bilibili/app/job/main/activity/conf"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao .
type Dao struct {
	httpClient *blademaster.Client
	kfcDelURL  string
}

// New init
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		httpClient: blademaster.NewClient(c.HTTPClient),
		kfcDelURL:  c.Host.APICo + _kfcDelURI,
	}
	return
}
