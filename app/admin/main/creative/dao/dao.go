package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/creative/conf"
	article "github.com/namelessup/bilibili/app/interface/openplatform/article/rpc/client"
	accapi "github.com/namelessup/bilibili/app/service/main/account/api"
	archive "github.com/namelessup/bilibili/app/service/main/archive/api/gorpc"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/jinzhu/gorm"
)

const (
	_msgURL = "/api/notify/send.user.notify.do"
)

// Dao dao.
type Dao struct {
	c         *conf.Config
	DB        *gorm.DB
	DBArchive *gorm.DB
	acc       accapi.AccountClient
	arc       *archive.Service2
	art       *article.Service
	es        *elastic.Elastic
	msgURL    string
	// http
	client *bm.Client
}

// New new a dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:         c,
		DB:        orm.NewMySQL(c.ORM),
		DBArchive: orm.NewMySQL(c.ORMArchive),
		arc:       archive.New2(c.ArchiveRPC),
		art:       article.New(c.ArticleRPC),
		es:        elastic.NewElastic(nil),
		// http client
		client: bm.NewClient(c.HTTPClient),
	}
	d.msgURL = c.Host.Msg + _msgURL
	d.initORM()
	var err error
	if d.acc, err = accapi.NewClient(c.AccClient); err != nil {
		panic(err)
	}
	return
}

func (d *Dao) initORM() {
	d.DB.LogMode(true)
	d.DBArchive.LogMode(true)
	d.DB.SingularTable(true)
}

// Ping check connection of db , mc.
func (d *Dao) Ping(c context.Context) (err error) {
	if d.DB != nil {
		err = d.DB.DB().PingContext(c)
	}
	if d.DBArchive != nil {
		err = d.DBArchive.DB().PingContext(c)
	}
	return
}

// Close close connection of db , mc.
func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
	if d.DBArchive != nil {
		d.DBArchive.Close()
	}
}

// ProfileStat get account.
func (d *Dao) ProfileStat(c context.Context, mid int64) (res *accapi.ProfileStatReply, err error) {
	arg := &accapi.MidReq{Mid: mid}
	if res, err = d.acc.ProfileWithStat3(c, arg); err != nil {
		log.Error("d.acc.ProfileWithStat3() error(%v)", err)
		err = ecode.CreativeAccServiceErr
	}
	return
}
