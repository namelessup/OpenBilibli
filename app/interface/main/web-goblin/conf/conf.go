package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/time"

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
	Tracer   *trace.Config
	Memcache *memcache.Config
	Ecode    *ecode.Config
	// ArchiveRPC
	ArchiveRPC *rpc.ClientConfig
	TagRPC     *rpc.ClientConfig
	SuitRPC    *rpc.ClientConfig
	AccountRPC *rpc.ClientConfig
	// auth
	Auth *auth.Config
	// Mysql
	DB *DB
	// Redis
	Redis *Redis
	// http client
	HTTPClient   *bm.ClientConfig
	SearchClient *bm.ClientConfig
	// Rule
	Rule *Rule
	// Pendants
	Pendants []*Pendant
	// Host
	Host      host
	Wechat    wechat
	Es        *elastic.Config
	OutSearch OutSearch
	Recruit   *Recruit
}

// Recruit .
type Recruit struct {
	MokaURI string
	Orgid   string
}

// OutSearch search out .
type OutSearch struct {
	Rspan        int64
	AcPgcFull    []string
	AcPgcIncre   []string
	AcUgcFull    []string
	AcUgcIncre   []string
	DealCommFull int32
	DealLikeFull int32
}

// Redis redis struct .
type Redis struct {
	*redis.Config
}

// Rule .
type Rule struct {
	Gid            int64
	ChCardInterval time.Duration
}

// Pendant .
type Pendant struct {
	Pid   int64
	Level int64
}

// DB .
type DB struct {
	Goblin *sql.Config
	Show   *sql.Config
}

type host struct {
	Wechat string
	PgcURI string
}

type wechat struct {
	AppID  string
	Secret string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init conf .
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
