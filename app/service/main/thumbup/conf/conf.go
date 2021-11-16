package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/tidb"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/rate"
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
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// Config config set
type Config struct {
	// elk
	Log *log.Config
	// BM
	BM *bm.ServerConfig
	// rpc server
	RPCServer *rpc.ServerConfig
	GRPC      *warden.ServerConfig
	// tracer
	Tracer *trace.Config
	// verify
	Verify *verify.Config
	Rate   *rate.Config
	// redis
	Redis *Redis
	// memcache
	Memcache *Memcache
	// Tidb
	Tidb *tidb.Config
	// ecode
	Ecode       *ecode.Config
	StatDatabus *databus.Config
	LikeDatabus *databus.Config
	ItemDatabus *databus.Config
	UserDatabus *databus.Config
	StatMerge   *StatMerge
	// ThumbUp
	ThumbUp ThumbUp
}

// StatMerge .
type StatMerge struct {
	Business string
	Target   int64
	Sources  []int64
}

// Memcache config
type Memcache struct {
	*memcache.Config
	StatsExpire time.Duration
}

// Redis config
type Redis struct {
	*redis.Config
	StatsExpire     time.Duration
	UserLikesExpire time.Duration
	ItemLikesExpire time.Duration
}

// ThumbUp thumb up config
type ThumbUp struct {
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
