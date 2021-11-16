package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/app/service/main/identify-game/model"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
)

// Conf global variable.
var (
	Conf     = &Config{}
	client   *conf.Client
	confPath string
)

// Config struct of conf.
type Config struct {
	// base
	// log
	Xlog *log.Config
	//Tracer *conf.Tracer
	Tracer *trace.Config
	// Verify
	Verify *verify.Config
	// BM
	BM *blademaster.ServerConfig
	// http client
	HTTPClient *blademaster.ClientConfig
	// memcache
	Memcache *memcache.Config
	// url router map
	Dispatcher *Dispatcher
	// RPCServer rpc server
	RPCServer *rpc.ServerConfig
	// grpc server
	WardenServer *warden.ServerConfig
	// passport
	Passport *PassportConfig
}

// Dispatcher router map
type Dispatcher struct {
	Name        string
	Oauth       map[string]string
	RenewToken  map[string]string
	RegionInfos []*model.RegionInfo
}

// PassportConfig identify config
type PassportConfig struct {
	Host map[string]string
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
			log.Info("config event")
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

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init int config
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
}
