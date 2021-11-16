package charge

import (
	"context"
	"flag"
	"path/filepath"
	"testing"

	"github.com/namelessup/bilibili/app/job/main/growup/conf"
	"github.com/namelessup/bilibili/app/job/main/growup/service/ctrl"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	s *Service
)

func init() {
	dir, _ := filepath.Abs("../../cmd/growup-job.toml")
	flag.Set("conf", dir)
	conf.Init()
	if s == nil {
		s = New(conf.Conf, ctrl.NewUnboundedExecutor())
	}
}

func TestPing(t *testing.T) {
	Convey("Test_Ping", t, func() {
		err := s.Ping(context.Background())
		So(err, ShouldBeNil)
	})
}
