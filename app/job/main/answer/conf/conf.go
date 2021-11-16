package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/netutil"
	"github.com/namelessup/bilibili/library/queue/databus"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf service config
	Conf = &Config{}
)

// Config def.
type Config struct {
	Log        *log.Config
	Databus    *DataSource
	Mysql      *sql.Config
	BFS        *BFS
	HTTPClient *bm.ClientConfig
	Properties *Properties
	Backoff    *netutil.BackoffConfig
	BM         *bm.ServerConfig
}

// BFS bfs config
type BFS struct {
	Timeout     xtime.Duration
	MaxFileSize int
	Bucket      string
	URL         string
	Method      string
	Key         string
	Secret      string
	Host        string
}

// Properties def.
type Properties struct {
	UploadInterval     xtime.Duration
	AccountIntranetURI string
	MaxRetries         int
	FontFilePath       string
}

// DataSource databus source
type DataSource struct {
	Labour  *databus.Config
	Account *databus.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf.
func Init() (err error) {
	if confPath == "" {
		return configCenter()
	}
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func configCenter() (err error) {
	if client, err = conf.New(); err != nil {
		panic(err)
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
