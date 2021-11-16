package dao

import (
	"context"
	xhttp "net/http"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/member/conf"
	"github.com/namelessup/bilibili/app/admin/main/member/dao/block"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/database/sql"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/queue/databus"

	"github.com/namelessup/bilibili/library/database/hbase.v2"

	"github.com/jinzhu/gorm"
)

// Dao dao
type Dao struct {
	c *conf.Config
	// db
	member         *gorm.DB
	memberRead     *gorm.DB
	account        *gorm.DB
	block          *block.Dao
	fhbyophbase    *hbase.Client
	fhbymidhbase   *hbase.Client
	httpClient     *bm.Client
	passportClient *bm.Client
	bfsclient      *xhttp.Client
	upUnameURL     string
	queryByMidsURL string
	msgURL         string
	expMsgDatabus  *databus.Databus
	es             *elastic.Elastic
	redis          *redis.Pool
	memcache       *memcache.Pool
	merakURL       string
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:              c,
		member:         orm.NewMySQL(c.ORM.Member),
		memberRead:     orm.NewMySQL(c.ORM.MemberRead),
		account:        orm.NewMySQL(c.ORM.Account), // account-php 库
		fhbyophbase:    hbase.NewClient(c.FHByOPHBase),
		fhbymidhbase:   hbase.NewClient(c.FHByMidHBase),
		httpClient:     bm.NewClient(c.HTTPClient.Read),
		passportClient: bm.NewClient(c.HTTPClient.Passport),
		bfsclient:      &xhttp.Client{Timeout: time.Duration(c.FacePriBFS.Timeout)},
		msgURL:         c.Host.Message + _msgURL,
		upUnameURL:     c.Host.Passport + _updateUname,
		queryByMidsURL: c.Host.Passport + _queryByMids,
		merakURL:       c.Host.Merak + "/",
		expMsgDatabus:  databus.New(c.ExpMsgDatabus),
		es:             elastic.NewElastic(c.ES),
		redis:          redis.NewPool(c.Redis),
		memcache:       memcache.NewPool(c.Memcache),
	}
	dao.block = block.New(c, dao.httpClient, memcache.NewPool(c.BlockMemcache), sql.NewMySQL(c.BlockMySQL))
	dao.initORM()
	return
}

// Close close the resource.
func (d *Dao) Close() {
	if d.member != nil {
		d.member.Close()
	}
	if d.redis != nil {
		d.redis.Close()
	}
}

func (d *Dao) initORM() {
	d.member.LogMode(true)
}

// Ping dao ping
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.member.DB().PingContext(ctx); err != nil {
		return
	}
	if err = d.pingRedis(ctx); err != nil {
		return
	}
	return
}

// BlockImpl is
func (d *Dao) BlockImpl() *block.Dao {
	return d.block
}
