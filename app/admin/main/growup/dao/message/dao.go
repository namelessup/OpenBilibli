package message

import (
	"github.com/namelessup/bilibili/app/admin/main/growup/conf"
	xhttp "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Dao is message dao
type Dao struct {
	c                *conf.Config
	uri, creativeURL string
	client           *xhttp.Client
}

// New a message dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:           c,
		client:      xhttp.NewClient(c.HTTPClient),
		uri:         c.Host.Message + "/api/notify/send.user.notify.do",
		creativeURL: c.Host.Creative + "/x/internal/creative/join/growup/account",
	}
	return
}
