package service

import (
	"fmt"
	"testing"

	"github.com/namelessup/bilibili/app/service/main/vip/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceAssociateVips(t *testing.T) {
	Convey(" TestServiceAssociateVips ", t, func() {
		res := s.AssociateVips(c, &model.ArgAssociateVip{
			Platform: "android",
			MobiApp:  "android",
		})
		fmt.Println("res", res)
		So(res, ShouldNotBeNil)
	})
}
