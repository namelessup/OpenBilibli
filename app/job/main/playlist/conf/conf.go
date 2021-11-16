package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	xlog "github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// Config .
type Config struct {
	// Env .
	Env string
	// App .
	App *bm.App
	// Log is github.com/namelessup/bilibili log.
	Log *xlog.Config
	// Tracer .
	Tracer *trace.Config
	// PlaylistStatSub databus.
	PlaylistViewSub  *databus.Config
	PlaylistFavSub   *databus.Config
	PlaylistReplySub *databus.Config
	PlaylistShareSub *databus.Config
	// HTTPServer .
	HTTPServer *bm.ServerConfig
	// HTTPClient .
	HTTPClient *bm.ClientConfig
	// RPC .
	PlaylistRPC *rpc.ClientConfig
	// Mysql .
	Mysql *sql.Config
	// Redis .
	Redis *redis.Config
	// Job params .
	Job *job
}

type job struct {
	InterceptOn      bool
	ViewCacheTTL     xtime.Duration
	UpdateDbInterval xtime.Duration
}

var (
	confPath string
	client   *conf.Client
	// Conf config.
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
