package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// Conf info.
var (
	confPath string
	Conf     = &Config{}
	client   *conf.Client
)

// Config struct.
type Config struct {
	// base
	// tick
	Tick time.Duration
	// max assist count
	MaxAssCnt  int
	MaxTypeCnt int64
	// app
	App *bm.App
	// host
	Host *Host
	// elk
	Xlog *log.Config
	// tracer
	Tracer *trace.Config
	BM     *bm.ServerConfig
	// rpc server2
	RPCServer *rpc.ServerConfig
	// http client
	HTTPClient *HTTPClient
	// db
	DB *DB
	// ecode
	Ecode *ecode.Config
	// rpc client2
	ArchiveRPC *rpc.ClientConfig
	// mc
	Memcache *Memcache
	// redis
	Redis *Redis
	// databus sub
	RelationSub *databus.Config
	AccClient   *warden.ClientConfig
}

// HTTPServers Http Servers
type HTTPServers struct {
	Inner *bm.ServerConfig
}

// Memcache conf
type Memcache struct {
	Assist *struct {
		*memcache.Config
		SubmitExpire time.Duration
	}
}

// Redis conf
type Redis struct {
	Assist *struct {
		*redis.Config
		Expire time.Duration
	}
}

// Host conf.
type Host struct {
	Message string
	Account string
}

// DB conf.
type DB struct {
	Assist *sql.Config
}

// HTTPClient conf.
type HTTPClient struct {
	Normal *bm.ClientConfig
	Slow   *bm.ClientConfig
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init conf.
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
