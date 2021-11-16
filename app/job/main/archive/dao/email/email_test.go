package email

import (
	"github.com/namelessup/bilibili/app/job/main/archive/model/result"
	"github.com/namelessup/bilibili/app/service/main/archive/api"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_PGCNotifyMail(t *testing.T) {
	Convey("PGCNotifyMail", t, func() {
		d.PGCNotifyMail(&api.Arc{}, &result.Archive{}, &result.Archive{})
	})
}
