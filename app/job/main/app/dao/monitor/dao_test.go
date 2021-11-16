package monitor

import (
	"flag"
	"path/filepath"

	"github.com/namelessup/bilibili/app/job/main/app/conf"
)

var (
	d *Dao
)

func init() {
	dir, _ := filepath.Abs("../../cmd/app-job-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	d = New(conf.Conf)
}
