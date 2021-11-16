package reply

import (
	"flag"
	"path/filepath"

	"github.com/namelessup/bilibili/app/job/main/archive/conf"
)

var (
	d *Dao
)

func init() {
	dir, _ := filepath.Abs("../../cmd/archive-job-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	d = New(conf.Conf)
}
