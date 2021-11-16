package upcrmservice

import (
	"context"
	"github.com/namelessup/bilibili/app/admin/main/up/model"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUpcrmserviceCommandRefreshUpRank(t *testing.T) {
	convey.Convey("CommandRefreshUpRank", t, func(ctx convey.C) {
		var (
			con = context.Background()
			arg = &model.CommandCommonArg{}
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			result, err := s.CommandRefreshUpRank(con, arg)
			ctx.Convey("Then err should be nil.result should not be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldBeNil)
				ctx.So(result, convey.ShouldNotBeNil)
			})
		})
	})
}
