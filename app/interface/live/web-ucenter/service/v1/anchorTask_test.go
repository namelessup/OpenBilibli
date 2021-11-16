package v1

import (
	"flag"
	. "github.com/smartystreets/goconvey/convey"
	api "github.com/namelessup/bilibili/app/interface/live/web-ucenter/api/http/v1"
	"github.com/namelessup/bilibili/app/interface/live/web-ucenter/conf"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/metadata"

	"testing"
)

var (
	AnchorTask *AnchorTaskService
)

func init() {
	flag.Set("conf", "../../cmd/test.toml")
	var err error
	if err = conf.Init(); err != nil {
		panic(err)
	}
	AnchorTask = NewAnchorTaskService(conf.Conf)
}

// go test  -test.v -test.run TestServiceAllowanceList
func TestMyReward(t *testing.T) {
	Convey("TestMyReward", t, func() {

		ctx := metadata.NewContext(bm.Context{}, metadata.MD{
			"mid": 10000,
		})

		res, err := AnchorTask.MyReward(ctx, &api.AnchorTaskMyRewardReq{
			Page: 1,
		})
		t.Logf("%+v", res)
		So(err, ShouldBeNil)
	})
}
