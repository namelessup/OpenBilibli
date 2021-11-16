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
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

const _configKey = "tag-job.toml"

var (
	confPath string
	// Conf .
	Conf = &Config{}
)

// Config config .
type Config struct {
	Log        *log.Config
	Tag        *Tag
	Platform   *Platform
	Redis      *Redis
	Databus    *Databus
	ArchiveRPC *rpc.ClientConfig
	FTP        *FTP
	BM         *bm.ServerConfig
	RedisRank  *redis.Config
	RedisTag   *redis.Config
	Memcache   *memcache.Config
	//  GRPCClient
	FilterGRPC *warden.ClientConfig
}

// Tag tag.
type Tag struct {
	MaxArcsLimit   int
	ArcTagSharding int
	Tick           time.Duration
	TagInfoPath    string
}

// FTP FTP.
type FTP struct {
	Addr     string
	User     string
	Password string
	HomeDir  string
	Timeout  time.Duration
}

// Redis redis .
type Redis struct {
	Rank *RankRedis
}

// RankRedis rank redis .
type RankRedis struct {
	Redis  *redis.Config
	Expire *RankExpire
}

// RankExpire rank redis expire .
type RankExpire struct {
	TagNewArc time.Duration
}

// Databus databus .
type Databus struct {
	Archive *databus.Config
	Tag     *databus.Config
}

// Platform .
type Platform struct {
	MySQL *sql.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init conf .
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
	if value, ok = client.Value(_configKey); !ok {
		panic(err)
	}
	_, err = toml.Decode(value, &Conf)
	return
}
