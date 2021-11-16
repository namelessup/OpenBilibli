package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/app/service/main/coin/model"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/antispam"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// global var
var (
	confPath string
	Conf     = &Config{}
	client   *conf.Client
)

// Config config.
type Config struct {
	// rpc server2
	RPCServer *rpc.ServerConfig
	GRPC      *warden.ServerConfig
	// BM
	BM *bm.ServerConfig
	// db
	DB *DB
	// redis
	Redis *Redis
	// databus
	DbBigData *databus.Config
	DbCoinJob *databus.Config
	// new stat.
	Stat *Stat
	// rpc client
	MemberRPC *warden.ClientConfig
	// tracer
	Tracer *trace.Config
	// verify
	Verify *verify.Config
	// Log
	Log    *log.Config
	Report *log.AgentConfig
	// ding url
	TagURL     string
	HTTPClient *bm.ClientConfig
	// Antispam
	Antispam *antispam.Config
	// Memcache .
	Memcache   *Memcache
	Businesses []*model.Business
	Coin       *Coin
	UserReport *databus.Config
	StatMerge  *StatMerge
}

// Coin .
type Coin struct {
	ESLogURL string
}

// StatMerge .
type StatMerge struct {
	Business string
	Target   int64
	Sources  []int64
}

// Stat databus stat conf.
type Stat struct {
	Databus *databus.Config
}

// DB db config.
type DB struct {
	Coin *sql.Config
}

// Memcache mc config.
type Memcache struct {
	*memcache.Config
	Expire    time.Duration
	ExpExpire time.Duration
}

// Redis redis conf.
type Redis struct {
	*redis.Config
	Expire time.Duration
}

func configCenter() (err error) {
	if client, err = conf.New(); err != nil {
		panic(err)
	}
	err = load()
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
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf.
func Init() (err error) {
	if confPath == "" {
		return configCenter()
	}
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
