package archive

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/namelessup/bilibili/app/job/main/creative/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/smartystreets/goconvey/convey"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	if os.Getenv("DEPLOY_ENV") != "" {
		flag.Set("app_id", "main.archive.creative-job")
		flag.Set("conf_token", "43943fda0bb311e8865c66d44b23cda7")
		flag.Set("tree_id", "16037")
		flag.Set("conf_version", "docker-1")
		flag.Set("deploy_env", "uat")
		flag.Set("conf_host", "config.bilibili.co")
		flag.Set("conf_path", "/tmp")
		flag.Set("region", "sh")
		flag.Set("zone", "sh001")
	} else {
		flag.Set("conf", "../../cmd/creative-job.toml")
	}
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	os.Exit(m.Run())
}

func Test_UpCount(t *testing.T) {
	var (
		err error
		c   = context.TODO()
		mid = int64(123)
	)
	convey.Convey("UpCount", t, func(ctx convey.C) {
		ctx.Convey("When everything gose positive", func(ctx convey.C) {
			ctx.Convey("Then err should be nil.res should not be nil.", func(ctx convey.C) {
				_, err = d.UpCount(c, mid)
				ctx.So(err, convey.ShouldEqual, err)
			})
		})
	})
}
