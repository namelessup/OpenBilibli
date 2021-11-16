package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/net/rpc/liverpc"

	"github.com/namelessup/bilibili/library/database/sql"

	"github.com/namelessup/bilibili/library/queue/databus"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// Config .
type Config struct {
	Log      *log.Config
	BM       *bm.ServerConfig
	Verify   *verify.Config
	Tracer   *trace.Config
	Redis    *Redis
	Database *Database
	Ecode    *ecode.Config
	Cfg      *Cfg
	// databus
	GiftPaySub    *databus.Config
	GiftFreeSub   *databus.Config
	AddCapsuleSub *databus.Config
	UserReport    *databus.Config
	LiveRpc       map[string]*liverpc.ClientConfig
	HTTPClient    *bm.ClientConfig
	CouponConf    *CouponConfig
}

// CouponConfig .
type CouponConfig struct {
	Url    string
	Coupon map[string]string
}

// Database mysql
type Database struct {
	Lottery *sql.Config
}

// Redis redis
type Redis struct {
	Lottery *redis.Config
}

// Cfg def
type Cfg struct {
	// ExpireCountFrequency crontab frequency
	ExpireCountFrequency string
	CouponRetryFrequency string
	ConsumerProcNum      int64
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init conf
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
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
