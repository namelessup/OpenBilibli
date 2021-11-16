package service

import (
	"context"
	"flag"
	"log"
	"testing"

	"github.com/namelessup/bilibili/app/infra/notify/conf"
	"github.com/namelessup/bilibili/app/infra/notify/model"
	"github.com/namelessup/bilibili/library/ecode"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	s *Service
)

func TestMain(m *testing.M) {
	var err error
	flag.Set("conf", "../cmd/notify-test.toml")
	if err = conf.Init(); err != nil {
		log.Println(err)
		return
	}
	s = New(conf.Conf)
	m.Run()
}
func TestPub(t *testing.T) {
	s.pubConfs = map[string]*model.Pub{
		"test-test": &model.Pub{
			Topic: "test",
			Group: "test",
		},
	}
	Convey("test pub", t, func() {
		err := s.Pub(context.TODO(), &model.ArgPub{Topic: "test", Group: "test", AppSecret: "test"})
		So(err, ShouldEqual, ecode.AccessDenied)
	})
}
