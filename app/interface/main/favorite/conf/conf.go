package conf

import (
	"flag"
	"github.com/namelessup/bilibili/library/net/rpc/warden"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/antispam"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/supervisor"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf Config
	Conf *Config
)

// Config is favorte conf
type Config struct {
	// base
	// log
	Log *log.Config
	App *bm.App
	// favorite config
	Fav      *Fav
	Platform *Platform
	// BM blademaster
	BM *bm.ServerConfig
	// redis
	Redis *Redis
	// memcache
	Memcache *Memcache
	// databus
	JobDatabus *databus.Config
	// Verify
	Verify *verify.Config
	// Auth
	Auth *auth.Config
	// rpc client
	RPCClient2 *RPC
	// tracer
	Tracer *trace.Config
	// http client
	HTTPClient *bm.ClientConfig
	// ecode
	Ecode *ecode.Config
	// Antispam
	Antispam *antispam.Config
	// Supervisior
	Supervisor *supervisor.Config
	// collector
	Infoc2 *infoc.Config
}

// RPC contain all rpc conf
type RPC struct {
	Archive   *rpc.ClientConfig
	Favorite  *rpc.ClientConfig
	FavClient *warden.ClientConfig
}

// Fav config
type Fav struct {
	// the max of the num of favorite folders
	MaxFolders  int
	MaxPagesize int
	MaxNameLen  int
	MaxDescLen  int
	// the num of operation
	MaxOperationNum int
	// the num of default favorite
	DefaultFolderLimit int
	NormalFolderLimit  int
	// cache expire
	Expire time.Duration
}

// Platform config
type Platform struct {
	MaxFolders int
	MaxNameLen int
	MaxDescLen int
}

// Redis redis conf
type Redis struct {
	*redis.Config
	Expire      time.Duration
	CoverExpire time.Duration
}

// Memcache memcache conf
type Memcache struct {
	*memcache.Config
	Expire time.Duration
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf
func Init() (err error) {
	if confPath == "" {
		return configCenter()
	}
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func configCenter() (err error) {
	var (
		ok     bool
		value  string
		client *conf.Client
	)
	if client, err = conf.New(); err != nil {
		return
	}
	if value, ok = client.Toml2(); !ok {
		panic(err)
	}
	_, err = toml.Decode(value, &Conf)
	return
}
