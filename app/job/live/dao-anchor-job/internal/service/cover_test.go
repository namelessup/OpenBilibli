package service

import (
	"flag"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/namelessup/bilibili/app/job/live/dao-anchor-job/internal/conf"
)

var s *Service

func init() {
	flag.Set("conf", "../../cmd/test.toml")
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s = New(conf.Conf)
}
func TestCover(t *testing.T) {
	Convey("testCover", t, func() {
		s.updateKeyFrame()
	})

}
