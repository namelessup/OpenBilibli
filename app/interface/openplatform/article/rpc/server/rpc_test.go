package server

import (
	"context"
	"testing"

	artmdl "github.com/namelessup/bilibili/app/interface/openplatform/article/model"
	artsrv "github.com/namelessup/bilibili/app/interface/openplatform/article/rpc/client"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	ctx = context.TODO()
)

func WithRPC(f func(client *artsrv.Service)) func() {
	return func() {
		client := artsrv.New(nil)
		f(client)
	}
}

func Test_ArticleRemainCount(t *testing.T) {
	Convey("ArticleRemainCount", t, WithRPC(func(client *artsrv.Service) {
		arg := &artmdl.ArgMid{Mid: 27515310}
		res, err := client.ArticleRemainCount(ctx, arg)
		So(err, ShouldBeNil)
		t.Logf("count: %d", res)
	}))
}
