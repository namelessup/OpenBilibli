package usersuit

import (
	"flag"

	"github.com/namelessup/bilibili/app/interface/main/account/conf"
)

var d *Dao

func init() {
	flag.Parse()

	flag.Set("conf", "../../cmd/account-interface-example.toml")
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
}
