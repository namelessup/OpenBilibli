package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/queue/databus/databusutil"
	xtime "github.com/namelessup/bilibili/library/time"

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
	Log *log.Config
	//Tracer *conf.Tracer
	Tracer *trace.Config
	//Databus databus
	DataBus *DataBus
	// Databusutil
	Databusutil *databusutil.Config
	// memcache
	Memcaches map[string]*Memcache
	// BM
	BM *bm.ServerConfig
	// AuthDB
	AuthDB *sql.Config
	// AuthMC
	AuthMC *memcache.Config
	// CheckConf
	CheckConf *CheckConf
}

// DataBus databus.
type DataBus struct {
	IdentifySub *databus.Config
	AuthDataBus *databus.Config
}

// Memcache contains prefix
type Memcache struct {
	Prefix string
	*memcache.Config
}

// CheckConf .
type CheckConf struct {
	Switch   bool
	ChanNum  int
	ChanSize int
	Ticker   xtime.Duration
	Count    int64
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
