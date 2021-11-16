package conf

import (
	"errors"
	"flag"
	"github.com/namelessup/bilibili/library/net/rpc/warden"

	"github.com/namelessup/bilibili/library/queue/databus"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/database/sql"
	eCode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"

	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	liverRPC "github.com/namelessup/bilibili/library/net/rpc/liverpc"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// TraceInit weather need init trace
	TraceInit bool
	client    *conf.Client
	// Conf config
	Conf = &Config{}
)

// Config .
type Config struct {
	Log                 *log.Config
	BM                  *bm.ServerConfig
	BMClient            *bm.ClientConfig
	Verify              *verify.Config
	Tracer              *trace.Config
	VipRedis            *redis.Config
	GuardRedis          *redis.Config
	Redis               *redis.Config
	Memcache            *memcache.Config
	ExpMemcache         *memcache.Config
	LiveUserMysql       *sql.Config
	LiveAppMySQL        *sql.Config
	LiveAppORM          *orm.Config
	Ecode               *eCode.Config
	LiveVipChangePub    *databus.Config
	UserExpMySQL        *sql.Config
	LiveRPC             map[string]*liverRPC.ClientConfig
	LiveEntryEffectPub  *databus.Config
	GuardCfg            *GuardCfg
	AccountRPC          *rpc.ClientConfig
	Switch              *ConfigSwitch
	UserExpExpire       *UserExpExpireConf
	UserDaHangHaiExpire *UserDhhExpireConf
	// report
	Report        *databus.Config
	XanchorClient *warden.ClientConfig
}

// GuardCfg config for guard
type GuardCfg struct {
	OpenEntryEffectDatabus bool
	EnableGuardBroadcast   bool
	DanmuHost              string
}

// ConfigSwitch config for query
type ConfigSwitch struct {
	QueryExp int
}

// UserExpExpireConf config for cache expire
type UserExpExpireConf struct {
	ExpireTime int32
}

// UserDhhExpireConf config for cache expire
type UserDhhExpireConf struct {
	ExpireTime int32
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
	flag.BoolVar(&TraceInit, "traceInit", true, "default trace init")
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
