package subtitle

import (
	"github.com/namelessup/bilibili/app/interface/main/dm2/rpc/client"
	"github.com/namelessup/bilibili/app/interface/main/videoup/conf"
)

// Dao fn
type Dao struct {
	c   *conf.Config
	sub *client.Service
}

// New fn
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:   c,
		sub: client.New(c.SubRPC),
	}
	return
}
