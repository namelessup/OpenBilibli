package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/namelessup/bilibili/app/service/main/point/dao"
	"github.com/namelessup/bilibili/app/service/main/point/model"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/bouk/monkey"
	"github.com/smartystreets/goconvey/convey"
)

func TestServiceConfig(t *testing.T) {
	convey.Convey("Config", t, func(ctx convey.C) {
		var (
			c          = context.Background()
			changeType = int(3)
			mid        = int64(4780461)
			bp         = float64(1)
			phs        []*model.PointHistory
			ph         = &model.PointHistory{
				ID:           13,
				Mid:          4780461,
				Point:        60,
				ChangeType:   1,
				PointBalance: 418,
			}
		)
		phs = append(phs, ph)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			monkey.PatchInstanceMethod(reflect.TypeOf(s.dao), "SelPointHistory", func(_ *dao.Dao, _ context.Context, _ int64, _, _ xtime.Time) ([]*model.PointHistory, error) {
				return phs, nil
			})
			point, err := s.Config(c, changeType, mid, bp)
			ctx.Convey("Then err should be nil.point should not be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldBeNil)
				ctx.So(point, convey.ShouldNotBeNil)
			})
		})
	})
}

func TestServiceAllConfig(t *testing.T) {
	convey.Convey("AllConfig", t, func(ctx convey.C) {
		var (
			c = context.Background()
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			p1 := s.AllConfig(c)
			ctx.Convey("Then p1 should not be nil.", func(ctx convey.C) {
				ctx.So(p1, convey.ShouldNotBeNil)
			})
		})
	})
}
