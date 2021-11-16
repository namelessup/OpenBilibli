package web

import (
	"context"
	"net/url"
	"testing"

	"github.com/namelessup/bilibili/app/interface/main/web-goblin/model/web"

	. "github.com/smartystreets/goconvey/convey"
)

func TestService_Recruit(t *testing.T) {
	Convey("test recruit", t, WithService(func(s *Service) {
		var (
			ctx   = context.Background()
			param = url.Values{}
			ru    = &web.Params{
				Route: "v1/jobs",
			}
		)
		param.Set("mode", "social")
		res, err := s.Recruit(ctx, param, ru)
		So(len(res), ShouldBeGreaterThan)
		So(err, ShouldBeNil)
	}))
}
