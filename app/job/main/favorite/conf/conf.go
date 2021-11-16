package conf

import (
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf .
	Conf *Config
)

// Config .
type Config struct {
	// base
	// fav
	Fav *Fav
	// log
	Log *log.Config
	// db
	DB *DB
	// memcache
	Memcache *Memcache
	// redis
	Redis *Redis
	// stat fav databus
	StatFavDatabus *StatFavDatabus
	// job databus
	JobDatabus *databus.Config
	// BinlogDatabus databus.
	BinlogDatabus *databus.Config
	// rpc client
	RPCClient2 *RPC
	// BM blademaster
	BM *bm.ServerConfig
	// StatMerge for bnj
	StatMerge *StatMerge
	// playlist stat
	FavStatDatabus      *databus.Config
	ShareStatDatabus    *databus.Config
	MediaListCntDatabus *databus.Config
	// http client
	HTTPClient *bm.ClientConfig
}

// Fav favorite
type Fav struct {
	Proc        int64
	MaxPageSize int
	CleanCDTime time.Duration
	SleepTime   time.Duration
	GreyMod     int64
	WhiteMids   []int64
}

// StatMerge .
type StatMerge struct {
	Business int
	Target   int64
	Sources  []int64
}

// RPC rpc cliens.
type RPC struct {
	Archive *rpc.ClientConfig
	Article *rpc.ClientConfig
	Coin    *rpc.ClientConfig
}

// DB mysql.
type DB struct {
	// favorite db
	Fav  *sql.Config
	Read *sql.Config
}

// Redis redis conf.
type Redis struct {
	*redis.Config
	Expire      time.Duration
	IPExpire    time.Duration
	BuvidExpire time.Duration
}

// Memcache mc conf.
type Memcache struct {
	*memcache.Config
	Expire time.Duration
}

// StatsDatabus stats.
type StatsDatabus struct {
	*databus.Config
	Field string
	Type  int8
}

// StatFavDatabus new stats.
type StatFavDatabus struct {
	*databus.Config
	Consumers map[string]int8
}

// init
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

// configCenter remote config.
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
