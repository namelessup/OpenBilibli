package hbase

import (
	"context"
	"flag"
	"path/filepath"
	"testing"

	"github.com/namelessup/bilibili/app/job/main/videoup-report/conf"
	"github.com/namelessup/bilibili/app/job/main/videoup-report/model/task"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	d *Dao
)

func init() {
	dir, _ := filepath.Abs("../../cmd/videoup-report-job.toml")
	flag.Set("conf", dir)
	conf.Init()
	d = New(conf.Conf)
}

func Test_AddLog(t *testing.T) {
	Convey("AddLog", t, func() {
		err := d.AddLog(context.TODO(), &task.WeightLog{TaskID: 44441})
		So(err, ShouldBeNil)
	})
}
