package service

import (
	"testing"

	"github.com/namelessup/bilibili/app/service/main/vip/model"
	"github.com/namelessup/bilibili/library/time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSyncUserSuccess(t *testing.T) {
	Convey("Test SyncUser Success", t, func() {
		item := &model.VipUserInfo{
			Mid:            7593623,
			VipOverdueTime: time.Time(1672243200),
			IosOverdueTime: time.Time(1672243200),
			Mtime:          time.Time(1672243200),
			PayChannelID:   2,
			IsAutoRenew:    2,
			VipType:        1,
			VipStatus:      1,
			Ver:            7,
		}
		err := s.SyncUser(c, item)
		So(item.VipRecentTime, ShouldEqual, item.Mtime)
		So(err, ShouldBeNil)
	})
}
