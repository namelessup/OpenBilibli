package v1

import (
	"context"
	"flag"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	api "github.com/namelessup/bilibili/app/interface/live/app-interface/api/http/v1"
	"github.com/namelessup/bilibili/app/interface/live/app-interface/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

var (
	s *Service
)

func init() {
	flag.Set("conf", "../../cmd/test.toml")
	var err error
	if err = conf.Init(); err != nil {
		panic(err)
	}
	s = New(conf.Conf)
}

// go test  -test.v -test.run TestGetAllList
func TestGetAllList(t *testing.T) {
	Convey("TestGetAllList", t, func() {
		res, err := s.GetAllList(&bm.Context{Context: context.TODO()}, &api.GetAllListReq{})
		t.Logf("%v", res)
		So(err, ShouldBeNil)
	})
}

func TestChange(t *testing.T) {
	Convey("TestChange", t, func() {
		res, err := s.Change(&bm.Context{Context: context.TODO()}, &api.ChangeReq{})
		t.Logf("%v", res)
		So(err, ShouldBeNil)
	})
}
