package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	xlog "github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// Config .
type Config struct {
	// Env
	Env string
	// App
	App *bm.App
	// Xlog is github.com/namelessup/bilibili log.
	Xlog *xlog.Config
	// rpc
	FeedRPC    *rpc.ClientConfig
	AccountRPC *rpc.ClientConfig
	// multihttp
	BM   *bm.ServerConfig
	Auth *auth.Config
	// tracer
	Tracer *trace.Config
	// feed
	Feed *feed
	// memcache
	Memcache *mc
}

// Feed feed controls
type feed struct {
	DefaultSize int
	MaxSize     int
}

type mc struct {
	*memcache.Config
	FeedExpire xtime.Duration
}

var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init .
func Init() (err error) {
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
			xlog.Info("config reload")
			if load() != nil {
				xlog.Error("config reload err")
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
