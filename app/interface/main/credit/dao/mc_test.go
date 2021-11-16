package dao

import (
	"context"
	"testing"

	credit "github.com/namelessup/bilibili/app/interface/main/credit/model"

	. "github.com/smartystreets/goconvey/convey"
)

// TestMcJury .

func TestMcJury(t *testing.T) {
	var (
		c  = context.TODO()
		op = &credit.Opinion{
			Mid:     1,
			OpID:    632,
			Content: "aaaaa",
		}
	)
	Convey("return someting", t, func() {
		d.AddOpinionCache(c, op)
		mop, miss, err := d.OpinionsCache(c, []int64{632, 631})
		So(err, ShouldBeNil)
		So(miss, ShouldNotBeNil)
		So(mop, ShouldNotBeNil)
	})
}
