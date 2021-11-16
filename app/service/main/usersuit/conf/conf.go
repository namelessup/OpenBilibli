package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// Conf global variable.
var (
	Conf     = &Config{}
	client   *conf.Client
	confPath string
)

// Config struct of conf.
type Config struct {
	AccountIntranetURI string
	AccountURI         string
	VipURI             string
	PayURL             string
	NotifyURL          string

	PayInfo *PayInfo
	// base
	// host
	Host *Host
	// log
	Xlog *log.Config
	// tracer
	Tracer *trace.Config
	// Verify
	Verify *verify.Config
	// http
	BM *bm.ServerConfig
	// RPCServer rpc server
	RPCServer *rpc.ServerConfig
	// MySQL mysql
	MySQL *sql.Config
	// Redis .
	Redis *Redis
	// Memcache memcache.
	Memcache *Memcache
	// HTTPClient http client
	HTTPClient *bm.ClientConfig
	// GORPCClient
	GORPCClient *GORPCClient
	// MedalCache
	MedalCache *LocalCache
	// EquipCache
	EquipCache *LocalCache
	// AccountNotify account notify.
	AccountNotify *databus.Config
}

// GORPCClient .
type GORPCClient struct {
	Member *rpc.ClientConfig
	Coin   *rpc.ClientConfig
	Point  *rpc.ClientConfig
}

// Host define host conf.
type Host struct {
	MessageCo    string
	AccountCoURI string
	APICoURI     string
	LiveAPICo    string
}

// PayInfo pay basic info
type PayInfo struct {
	MerchantID        string
	MerchantProductID string
	CallBackURL       string
}

// Redis .
type Redis struct {
	*redis.Config
	InviteExpire  time.Duration
	PendantExpire time.Duration
}

// LocalCache .
type LocalCache struct {
	Size   int
	Expire time.Duration
}

// Memcache define memcache conf.
type Memcache struct {
	*memcache.Config
	MedalExpire time.Duration
	PointExpire time.Duration
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func remote() (err error) {
	if client, err = conf.New(); err != nil {
		return
	}
	if err = load(); err != nil {
		return
	}
	go func() {
		for range client.Event() {
			log.Info("config reload")
			if load() != nil {
				log.Error("config reload error (%v)", err)
			}
		}
	}()
	return
}

func load() (err error) {
	var (
		s       string
		ok      bool
		tmpConf *Config
	)
	if s, ok = client.Toml2(); !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init int config
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
}
