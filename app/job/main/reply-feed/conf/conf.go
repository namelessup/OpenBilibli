package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	xtime "github.com/namelessup/bilibili/library/time"

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
	Log            *log.Config
	BM             *bm.ServerConfig
	Verify         *verify.Config
	Tracer         *trace.Config
	Redis          *redis.Config
	RedisExpire    *RedisExpire
	Memcache       *memcache.Config
	MemcacheExpire *MemcacheExpire
	MySQL          *MySQL
	Databus        *Databus
	Ecode          *ecode.Config
	RefreshTime    int64
	HTTPClient     *bm.ClientConfig
}

// RedisExpire Redis
type RedisExpire struct {
	RedisReplySetExpire  xtime.Duration
	RedisReplyZSetExpire xtime.Duration
	RedisRefreshExpire   xtime.Duration
}

// MemcacheExpire Memcache
type MemcacheExpire struct {
	McStatExpire xtime.Duration
}

// Databus databus
type Databus struct {
	Stats *databus.Config
	Event *databus.Config
}

// MySQL mysql config
type MySQL struct {
	DB      *sql.Config
	DBSlave *sql.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init conf
func Init() (err error) {
	if confPath != "" {
		return local()
	}
	if err = remote(); err != nil {
		return
	}
	if Conf.RefreshTime <= 0 {
		panic("refresh time illegal.")
	}
	return
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
