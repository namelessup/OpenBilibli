package tag

import (
	"flag"
	"path/filepath"
	"time"

	"github.com/namelessup/bilibili/app/job/main/growup/conf"
	"github.com/namelessup/bilibili/app/job/main/growup/service/ctrl"
)

var (
	s *Service
)

func init() {
	dir, _ := filepath.Abs("../../cmd/growup-job.toml")
	flag.Set("conf", dir)
	conf.Init()
	s = New(conf.Conf, ctrl.NewUnboundedExecutor())
	time.Sleep(time.Second)
}

func WithService(f func(s *Service)) func() {
	return func() {
		// Reset(func() { CleanCache() })
		f(s)
	}
}
