package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf of config
	Conf = &Config{}
)

// Memcache memcache.
type Memcache struct {
	*memcache.Config
}

// Bfs Bfs.
type Bfs struct {
	Timeout     xtime.Duration
	MaxFileSize int
	Bucket      string
	Addr        string
	Key         string
	Secret      string
}

// Cfg def
type Cfg struct {
	// HotCroFre hotword crontab frequency
	HotCroFre string
	// DarkCroFre darkword crontab frequency
	DarkCroFre string
	//RunCront is run crontab
	RunCront bool
}

//Host host
type Host struct {
	Manager string
}

// Config def.
type Config struct {
	// base
	// http
	HTTPServer *bm.ServerConfig
	// httpClinet
	HTTPClient *bm.ClientConfig
	// host
	Host *Host
	// auth
	Auth *permit.Config
	// db
	ORM *orm.Config
	// db
	ORMResource *orm.Config
	// log
	Log *log.Config
	// tracer
	Tracer *trace.Config
	//mc
	Memcache *Memcache
	// Bfs
	Bfs *Bfs
	// log
	ManagerReport *databus.Config
	// BroadcastRPC grpc
	PGCRPC *warden.ClientConfig
	// rpc client
	AccountRPC *rpc.ClientConfig
	ArchiveRPC *rpc.ClientConfig
	// Cfg
	Cfg *Cfg
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
