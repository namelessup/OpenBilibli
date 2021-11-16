package dao

import (
	"flag"
	"path/filepath"

	"github.com/namelessup/bilibili/app/job/main/click/conf"
)

var (
	d *Dao
)

func init() {
	dir, _ := filepath.Abs("../cmd/click-job-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	d = New(conf.Conf)
}
