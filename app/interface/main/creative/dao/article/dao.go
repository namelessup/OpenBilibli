package article

import (
	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	article "github.com/namelessup/bilibili/app/interface/openplatform/article/rpc/client"
)

// Dao is archive dao.
type Dao struct {
	// config
	c *conf.Config
	// rpc
	art *article.Service
}

// New init api url
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		// rpc
		art: article.New(c.ArticleRPC),
	}
	return
}
