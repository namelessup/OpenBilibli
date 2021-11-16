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

var (
	confPath string
	// Conf conf.
	Conf   = &Config{}
	client *conf.Client
)

// Config config.
type Config struct {
	// log
	Xlog *log.Config
	// Tracer tracer
	Tracer *trace.Config
	// http
	BM *bm.ServerConfig
	// Database
	DB *DB
	// mc
	Memcache *Memcache
	//Databus databus
	DataBus *DataBus
	// DataUtil config
	DatabusUtil *databusutil.Config
	// IncSync
	IncSync *IncSync
	// FullSync
	FullSync *FullSync
	// Scheduler
	Scheduler *Scheduler
}

// DB db config
type DB struct {
	OriginDB  *sql.Config
	UserDB    *sql.Config
	EncryptDB *sql.Config
}

// Memcache memcache
type Memcache struct {
	*memcache.Config
	Expire xtime.Duration
}

// DataBus databus.
type DataBus struct {
	AsoBinLogSub *databus.Config
}

// IncSync increment sync
type IncSync struct {
	Switch bool
}

// FullSync full sync
type FullSync struct {
	AsoCountryCodeSwitch bool
	AsoAccount           *SyncConf
	AsoAccountInfo       *SyncConf
	AsoAccountReg        *SyncConf
	AsoAccountSns        *SyncConf
	AsoTelBindLog        *SyncConf
}

// SyncConf sync conf
type SyncConf struct {
	Switch   bool
	ChanNum  int
	ChanSize int
	Start    int64
	End      int64
	Count    int64
}

// Scheduler scheduler
type Scheduler struct {
	Switch             bool
	TelDuplicateCron   string
	EmailDuplicateCron string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init config.
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
		tmpConf = &Config{}
	)
	if s, ok = client.Toml2(); !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(s, tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}
