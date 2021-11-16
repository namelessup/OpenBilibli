package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
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
	Log           *log.Config
	BM            *bm.ServerConfig
	Verify        *verify.Config
	Tracer        *trace.Config
	Redis         *redis.Config
	MallMySQL     *sql.Config
	TicketMySQL   *sql.Config
	MallUgcMySQL  *sql.Config
	PgcSub        *databus.Config
	ElasticSearch *ElasticSearch
	Env           string
	SourcePath    string
	Bfs           *Bfs
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

// ElasticSearch config
type ElasticSearch struct {
	Addr    []string
	Check   xtime.Duration
	Timeout string
	Season  EsIndex
}

// EsIndex config
type EsIndex struct {
	Index   string
	Type    string
	Alias   string
	Mapping string
}

// Bfs config
type Bfs struct {
	Timeout     xtime.Duration
	MaxFileSize int
	Bucket      string
	Addr        string
	Key         string
	Secret      string
}
