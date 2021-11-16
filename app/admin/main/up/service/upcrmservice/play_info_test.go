package upcrmservice

import (
	"context"
	"github.com/namelessup/bilibili/app/admin/main/up/model/upcrmmodel"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUpcrmservicePlayQueryInfo(t *testing.T) {
	convey.Convey("PlayQueryInfo", t, func(ctx convey.C) {
		var (
			context = context.Background()
			arg     = &upcrmmodel.PlayQueryArgs{}
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			result, err := s.PlayQueryInfo(context, arg)
			ctx.Convey("Then err should be nil.result should not be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldBeNil)
				ctx.So(result, convey.ShouldNotBeNil)
			})
		})
	})
}
