package conf

import (
	"errors"
	"flag"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

type BucketConf struct {
	Key     string
	Secret  string
	Private bool
}

// Config .
type Config struct {
	Log              *log.Config
	BM               *bm.ServerConfig
	Verify           *verify.Config
	Tracer           *trace.Config
	Redis            *redis.Config
	Memcache         *memcache.Config
	MySQL            *sql.Config
	Ecode            *ecode.Config
	ResourceClient   *warden.ClientConfig
	CapsuleClient    *warden.ClientConfig
	RoomMngClient    *liverpc.ClientConfig
	StreamMngClient  *warden.ClientConfig
	LiveRpc          map[string]*liverpc.ClientConfig
	ResourceClientV2 *warden.ClientConfig
	Auth             *auth.Config
	Bucket           map[string]*BucketConf
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
