package conf

import (
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	xlog "github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/namelessup/bilibili/library/database/hbase.v2"

	"github.com/BurntSushi/toml"
)

const (
	configKey = "history.toml"
)

var (
	confPath string
	// Conf global
	Conf = &Config{}
)

// HBaseConfig ...
type HBaseConfig struct {
	*hbase.Config
	WriteTimeout xtime.Duration
	ReadTimeout  xtime.Duration
}

// Config  service conf
type Config struct {
	Tracer       *trace.Config
	History      *History
	BM           *bm.ServerConfig
	RPCClient2   *RPC
	Toview       *Redis
	Redis        *Redis
	Xlog         *xlog.Config
	Info         *HBaseConfig
	DataBus      *Databus
	Auth         *auth.Config
	Verify       *verify.Config
	Collector    *infoc.Config
	Ecode        *ecode.Config
	RPCServer    *rpc.ServerConfig
	GRPC         *warden.ServerConfig
	ThirdBusines *ThirdBusines
	Report       *databus.Config
}

// History history.
type History struct {
	Max         int
	Total       int
	Cache       int
	Page        int
	Size        int
	Ticker      xtime.Duration
	Pub         bool
	ConsumeSize int

	Migration bool
	Rate      int64
	Mids      []int64
}

// ThirdBusines Bangumi favorite.
type ThirdBusines struct {
	BangumiV2URL string
	SeasonURL    string
	HTTPClient   *bm.ClientConfig
}

// Databus .
type Databus struct {
	PlayPro    *databus.Config
	Merge      *databus.Config
	Experience *databus.Config
	Pub        *databus.Config
}

// Redis redis.
type Redis struct {
	*redis.Config
	Expire xtime.Duration
}

// RPC rpc.
type RPC struct {
	Archive  *rpc.ClientConfig
	Favorite *rpc.ClientConfig
	History  *warden.ClientConfig
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
	if value, ok = client.Value(configKey); !ok {
		panic(err)
	}
	_, err = toml.Decode(value, &Conf)
	return
}
