package service

import (
	"context"
	"flag"
	"path/filepath"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/history/conf"
	"github.com/namelessup/bilibili/library/cache/redis"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	s *Service
)

func CleanCache() {
	c := context.TODO()
	pool := redis.NewPool(conf.Conf.Redis.Config)
	pool.Get(c).Do("FLUSHDB")
}

func WithService(f func(s *Service)) func() {
	return func() {
		Reset(func() { CleanCache() })
		f(s)
	}
}
func init() {
	dir, _ := filepath.Abs("../cmd/history-interface-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	s = New(conf.Conf)
	time.Sleep(time.Second)
}
