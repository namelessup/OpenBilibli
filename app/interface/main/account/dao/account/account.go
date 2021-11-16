package account

import (
	"github.com/namelessup/bilibili/app/interface/main/account/conf"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	xhttp "net/http"
	"time"
)

// consts
const (
	_cmIsBusinessAccount   = "/basc/api/open_api/v1/up/business_account/is_sign_up"
	_cmBusinessAccountInfo = "/basc/api/open_api/v1/up/bus_account_info"
)

// Dao dao
type Dao struct {
	c          *conf.Config
	client     *bm.Client
	accCom     string
	accCo      string
	vipHost    string
	smsSendURI string
	bfsClient  *xhttp.Client
	mc         *memcache.Pool
	redis      *redis.Pool
}

// New new
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:          c,
		client:     bm.NewClient(c.HTTPClient.Normal),
		accCom:     c.Host.AccCom,
		accCo:      c.Host.AccCo,
		vipHost:    c.Host.Vip,
		smsSendURI: c.Host.API + _smsSendURI,
		bfsClient:  &xhttp.Client{Timeout: time.Duration(c.BFS.Timeout)},
		mc:         memcache.NewPool(c.AccMemcache),
		redis:      redis.NewPool(c.AccRedis),
	}
	return
}
