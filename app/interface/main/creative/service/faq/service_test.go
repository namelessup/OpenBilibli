package faq

import (
	"context"
	"flag"
	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	faqMdl "github.com/namelessup/bilibili/app/interface/main/creative/model/faq"
	"path/filepath"
	"testing"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/creative/service"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	s *Service
)

func init() {
	dir, _ := filepath.Abs("../../cmd/creative.toml")
	flag.Set("conf", dir)
	conf.Init()
	rpcdaos := service.NewRPCDaos(conf.Conf)
	s = New(conf.Conf, rpcdaos)
	time.Sleep(time.Second)
}

func WithService(f func(s *Service)) func() {
	return func() {
		Reset(func() {})
		f(s)
	}
}
func Test_Pre(t *testing.T) {
	var (
		c   = context.TODO()
		res map[string]*faqMdl.Faq
	)
	Convey("Pre", t, WithService(func(s *Service) {
		res = s.Pre(c)
		So(res, ShouldNotBeNil)
	}))
}
