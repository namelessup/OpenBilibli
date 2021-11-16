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
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"

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
	NotifyURL string

	// base
	// log
	Xlog *log.Config
	// tracer
	Tracer *trace.Config
	// app
	App *APP
	// bm service
	BM *bm.ServerConfig
	// db
	Mysql *sql.Config
	// mecache
	Memcache     *Memcache
	PendantRedis *PendantRedis
	// http client
	HTTPClient *bm.ClientConfig
	PageSize   int64

	Properties *Properties
	SuitRPC    *rpc.ClientConfig

	Databus *Databus
}

// Databus .
type Databus struct {
	AccountNotify *databus.Config
	VipBinLog     *databus.Config
}

// Memcache define memcache conf.
type Memcache struct {
	*memcache.Config
}

// PendantRedis pendant redis
type PendantRedis struct {
	*redis.Config
}

// APP appkey and sec
type APP struct {
	Key    string
	Secret string
}

// Properties app config.
type Properties struct {
	UpInfoURL string
	MedalCron string
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
