package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/namelessup/bilibili/app/admin/ep/tapd/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/log"
	xhttp "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/sync/pipeline/fanout"

	"github.com/jinzhu/gorm"
	"gopkg.in/gomail.v2"
)

// Dao dao.
type Dao struct {
	c          *conf.Config
	httpClient *xhttp.Client
	db         *gorm.DB
	email      *gomail.Dialer
	mc         *memcache.Pool
	cache      *fanout.Fanout
	expire     int32
}

// New init mysql db.
func New(c *conf.Config) *Dao {
	return &Dao{
		c:          c,
		httpClient: xhttp.NewClient(c.HTTPClient),
		db:         orm.NewMySQL(c.ORM),
		email:      gomail.NewDialer(c.Mail.Host, c.Mail.Port, c.Mail.Username, c.Mail.Password),
		mc:         memcache.NewPool(c.Memcache.Config),
		cache:      fanout.New("cache", fanout.Worker(5), fanout.Buffer(10240)),
		expire:     int32(time.Duration(c.Memcache.Expire) / time.Second),
	}
}

// Close close the resource.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
	if d.mc != nil {
		d.mc.Close()
	}
}

// Ping verify server is ok.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.DB().Ping(); err != nil {
		log.Info("dao.cloudDB.Ping() error(%v)", err)
	}
	return
}

// tokenCacheSave The err does not need to return, because this method is irrelevant.
func (d *Dao) tokenCacheSave(c context.Context, cacheItem *memcache.Item) {
	var f = func(c context.Context) {
		var (
			conn = d.mc.Get(c)
			err  error
		)
		defer conn.Close()
		if err = conn.Set(cacheItem); err != nil {
			log.Error("AddCache conn.Set(%s) error(%v)", cacheItem.Key, err)
		}
	}
	if err := d.cache.Do(c, f); err != nil {
		log.Error("Token cache save err(%v)", err)
	}
}

func (d *Dao) newRequest(method, url string, v interface{}) (req *http.Request, err error) {
	body := &bytes.Buffer{}
	if method != http.MethodGet {
		if err = json.NewEncoder(body).Encode(v); err != nil {
			log.Error("json encode value(%s) err(?) ", v, err)
			return
		}
	}
	if req, err = http.NewRequest(method, url, body); err != nil {
		log.Error("http new request url(?) err(?)", url, err)
	}
	return
}
