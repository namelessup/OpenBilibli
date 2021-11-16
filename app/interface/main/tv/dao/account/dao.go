package account

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	accwar "github.com/namelessup/bilibili/app/service/main/account/api"
)

// Dao is account dao.
type Dao struct {
	// rpc
	accClient accwar.AccountClient
}

// New account dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{}
	var err error
	if d.accClient, err = accwar.NewClient(c.AccClient); err != nil {
		panic(err)
	}
	return
}
