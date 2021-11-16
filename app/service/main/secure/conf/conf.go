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

	"github.com/namelessup/bilibili/library/database/hbase.v2"

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
	// app
	App *bm.App
	// tracer
	Tracer *trace.Config
	// goroutine sleep
	Tick time.Duration
	// http server
	BM *bm.ServerConfig
	// log
	Log *log.Config
	// verify
	Verify      *verify.Config
	LocationRPC *rpc.ClientConfig
	DataBus     *databus.Config
	Mysql       *Mysql
	Expect      *Expect
	// redis
	Redis *Redis
	// httpClient
	HTTPClient *bm.ClientConfig
	// rpc
	RPCServer *rpc.ServerConfig
	// HBase
	HBase *HBaseConfig
	// mc
	Memcache *Memcache
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// HBaseConfig extra hbase config for compatible
type HBaseConfig struct {
	*hbase.Config
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

// Memcache mc config.
type Memcache struct {
	*memcache.Config
	Expire time.Duration
}

// Mysql conf.
type Mysql struct {
	Secure *sql.Config
	DDL    *sql.Config
}

// Redis redis conf.
type Redis struct {
	*redis.Config
	Expire      time.Duration
	DoubleCheck time.Duration
}

// Expect Login expection config.
type Expect struct {
	Top         int64 // login loc count.
	Count       int64 // login time count.
	CloseCount  int64 // user close count.
	Rand        int64 // rand ratio
	DoubleCheck int64 // double login check location count.
}

// Init create config instance.
func Init() (err error) {
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
