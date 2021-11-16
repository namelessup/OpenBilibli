package conf

import (
	"github.com/namelessup/bilibili/app/common/openplatform/encoding"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	// Conf common conf
	Conf = &Config{}
)

//Config config struct
type Config struct {
	//数据库配置
	DB *DB
	// redis
	Redis *Redis
	// http client
	HTTPClient HTTPClient
	// http
	BM *blademaster.ServerConfig
	// tracer
	Tracer *trace.Config
	// log
	Log *log.Config
	// UT
	UT         *UT
	GRPCClient map[string]*warden.ClientConfig
	Encrypt    *encoding.EncryptConfig
	URLs       map[string]string
	//basecenter配置
	BaseCenter *BaseCenter
	Databus    map[string]*databus.Config

	TestProject *TestProject
}

// HTTPClient config
type HTTPClient struct {
	Read  *blademaster.ClientConfig
	Write *blademaster.ClientConfig
}

// HTTPServers Http Servers
type HTTPServers struct {
	Inner *blademaster.ServerConfig
	Local *blademaster.ServerConfig
}

// Redis config
type Redis struct {
	Master *redis.Config
	Expire time.Duration
}

// DB config
type DB struct {
	Master *sql.Config
}

// UT config
type UT struct {
	DistPrefix string
}

//BaseCenter 的配置
type BaseCenter struct {
	AppID string
	Token string
}

// TestProject 测试项目配置
type TestProject struct {
	IDs        []int64
	CheckQuery string
}

// Set set config and decode.
func (c *Config) Set(text string) error {
	var tmp Config
	if _, err := toml.Decode(text, &tmp); err != nil {
		return err
	}
	*c = tmp
	return nil
}
