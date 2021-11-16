package order

import (
	"context"
	"flag"
	"github.com/namelessup/bilibili/app/interface/main/videoup/conf"
	"os"
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	gock "gopkg.in/h2non/gock.v1"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	if os.Getenv("DEPLOY_ENV") != "" {
		flag.Set("app_id", "main.archive.videoup")
		flag.Set("conf_token", "9772c9629b00ac09af29a23004795051")
		flag.Set("tree_id", "2306")
		flag.Set("conf_version", "docker-1")
		flag.Set("deploy_env", "uat")
		flag.Set("conf_host", "config.bilibili.co")
		flag.Set("conf_path", "/tmp")
		flag.Set("region", "sh")
		flag.Set("zone", "sh001")
	} else {
		flag.Set("conf", "../cmd/videoup.toml")
	}
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	m.Run()
	os.Exit(0)
}

func httpMock(method, url string) *gock.Request {
	r := gock.New(url)
	r.Method = strings.ToUpper(method)
	d.client.SetTransport(gock.DefaultTransport)
	return r
}

func Test_PubTime(t *testing.T) {
	convey.Convey("PubTime", t, func(ctx convey.C) {
		var (
			err     error
			c       = context.Background()
			mid     = int64(2089809)
			orderID = int64(111)
			ip      = "127.0.0.1"
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			defer gock.OffAll()
			httpMock("Post", d.launchTimeURI).Reply(200).JSON(`{"code":21022,"data":""}`)
			_, err = d.PubTime(c, mid, orderID, ip)
			ctx.Convey("Then err should be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldNotBeNil)
			})
		})
	})
}

func Test_BindOrder(t *testing.T) {
	convey.Convey("BindOrder", t, func(ctx convey.C) {
		var (
			err     error
			c       = context.Background()
			mid     = int64(2089809)
			aid     = int64(10110826)
			orderID = int64(111)
			ip      = "127.0.0.1"
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			defer gock.OffAll()
			httpMock("Post", d.useExeOrderURI).Reply(200).JSON(`{"code":21022,"data":""}`)
			err = d.BindOrder(c, mid, aid, orderID, ip)
			ctx.Convey("Then err should be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldNotBeNil)
			})
		})
	})
}

func Test_Ups(t *testing.T) {
	convey.Convey("Ups", t, func(ctx convey.C) {
		var (
			err error
			c   = context.Background()
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			defer gock.OffAll()
			httpMock("GET", d.upsURI).Reply(200).JSON(`{"code":21022,"data":""}`)
			_, err = d.Ups(c)
			ctx.Convey("Then err should be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldNotBeNil)
			})
		})
	})
}

func Test_ExecuteOrders(t *testing.T) {
	convey.Convey("ExecuteOrders", t, func(ctx convey.C) {
		var (
			err error
			c   = context.Background()
			mid = int64(2089809)
			ip  = "127.0.0.1"
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			defer gock.OffAll()
			httpMock("GET", d.executeOrdersURI).Reply(200).JSON(`{"code":21022,"data":""}`)
			_, err = d.ExecuteOrders(c, mid, ip)
			ctx.Convey("Then err should be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldNotBeNil)
			})
		})
	})
}
