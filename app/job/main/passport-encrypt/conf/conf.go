package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

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
	// common conf
	// log
	Xlog *log.Config
	// Tracer tracer
	Tracer *trace.Config
	// http
	BM *BM
	// Database
	DB *DB
	// Step
	StepGroup *StepGroup
	// DataSwitch
	DataSwitch *DataSwitch
	// Group group
	Group *Group
	//Databus databus
	DataBus *DataBus
}

// BM http server config
type BM struct {
	Inner *bm.ServerConfig
	Local *bm.ServerConfig
}

// DB db config
type DB struct {
	OriginDB  *sql.Config
	EncryptDB *sql.Config
}

// Group multi group config collection.
type Group struct {
	AsoBinLog *GroupConfig
}

// GroupConfig group config.
type GroupConfig struct {
	// Size merge size
	Size int
	// Num merge goroutine num
	Num int
	// Ticker duration of submit merges when no new message
	Ticker time.Duration
	// Chan size of merge chan and done chan
	Chan int
}

// DataBus databus.
type DataBus struct {
	AsoBinLogSub *databus.Config
}

// StepGroup group
type StepGroup struct {
	Group1 *Step
	Group2 *Step
	Group3 *Step
	Group4 *Step
}

// Step data step
type Step struct {
	Start, End, Inc, Limit int64
}

// DataSwitch data trans swtich
type DataSwitch struct {
	Full, Inc bool
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
