package mis

import (
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/conf"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/dao"
)

//Mis http server
type Mis struct {
	c   *conf.Config
	dao *dao.Dao
}

// New for new mis obj
func New(c *conf.Config, d *dao.Dao) *Mis {
	m := &Mis{
		c:   c,
		dao: d,
	}
	return m
}
