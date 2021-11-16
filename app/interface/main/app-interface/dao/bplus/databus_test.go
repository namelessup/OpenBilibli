package bplus

import (
	"testing"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/app-interface/model/space"
	xtime "github.com/namelessup/bilibili/library/time"

	. "github.com/smartystreets/goconvey/convey"
)

// TestNotifyContribute dao ut.
func TestNotifyContribute(t *testing.T) {
	Convey("get DynamicCount", t, func() {
		var attrs *space.Attrs
		err := dao.NotifyContribute(ctx(), 27515258, attrs, xtime.Time(time.Now().Unix()))
		err = nil
		So(err, ShouldBeNil)
	})
}
