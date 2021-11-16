package conf

import (
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/antispam"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/supervisor"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
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
	// rpc server2
	RPCServer *rpc.ServerConfig
	// db
	MySQL *MySQL
	// redis
	Redis *Redis
	// memcache
	Memcache *Memcache
	// databus
	JobDatabus *databus.Config
	// verify
	Verify *verify.Config
	// rpc client
	RPCClient2 *RPC
	// tracer
	Tracer *trace.Config
	// http client
	HTTPClient *bm.ClientConfig
	// ecode
	Ecode *ecode.Config
	// TopicClient
	Topic *Topic
	// Antispam
	Antispam *antispam.Config
	// Supervisior
	Supervisor *supervisor.Config
	// collector
	Infoc2 *infoc.Config
	//grpc warden
	WardenServer *warden.ServerConfig
}

// RPC contain all rpc conf
type RPC struct {
	Account *warden.ClientConfig
	Archive *rpc.ClientConfig
	Filter  *rpc.ClientConfig
	Rank    *rpc.ClientConfig
}

// Topic Topic
type Topic struct {
	TopicURL   string
	HTTPClient *bm.ClientConfig
}

// Fav config
type Fav struct {
	// the max of the num of favorite folders
	MaxFolders      int
	MaxPagesize     int
	MaxBatchSize    int
	MaxDataSize     int
	MaxParallelSize int
	MaxRecentSize   int
	MaxNameLen      int
	MaxDescLen      int
	// the num of operation
	MaxOperationNum int
	// the num of default favorite
	DefaultFolderLimit int
	NormalFolderLimit  int
	// ApiHost api.bilibili.co .
	APIHost string
	// cache expire
	Expire time.Duration
	// cdtime cool down time
	CleanCDTime time.Duration
	// real-name switch
	RealNameOn bool
}

// Platform config
type Platform struct {
	MaxFolders int
	MaxNameLen int
	MaxDescLen int
}

// MySQL is mysql conf
type MySQL struct {
	// favorite db
	Fav  *sql.Config
	Read *sql.Config
	Push *sql.Config
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
