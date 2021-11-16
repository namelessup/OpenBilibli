package dao

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/namelessup/bilibili/app/job/live/wallet/conf"
	"github.com/namelessup/bilibili/library/log"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/namelessup/bilibili/app/job/live/wallet/model"
)

var (
	once    sync.Once
	d       *Dao
	ctx     = context.TODO()
	testUid int64
)

func initConf() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	defer log.Close()
}

func startService() {
	initConf()
	d = New(conf.Conf)
	time.Sleep(time.Second * 2)
}

func TestInitWallet(t *testing.T) {
	Convey("Init Wallet", t, func() {
		once.Do(startService)
		user := &model.User{}
		user.Uid = testUid
		_, err := d.InitWallet(ctx, user)
		So(err, ShouldBeNil)
	})
}

func TestMergeWallet(t *testing.T) {
	Convey("Merge Wallet", t, func() {
		once.Do(startService)
		_, err := d.MergeWallet(ctx, testUid, 1, 2, 3)
		So(err, ShouldBeNil)
	})
}
