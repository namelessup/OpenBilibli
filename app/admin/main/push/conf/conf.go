package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	sqlx "github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
)

var (
	// config
	confPath string
	client   *conf.Client
	// Conf .
	Conf = &Config{}
)

// Config def.
type Config struct {
	Ecode      *ecode.Config
	Auth       *permit.Config
	HTTPServer *bm.ServerConfig
	HTTPClient *bm.ClientConfig
	DPClient   *bm.ClientConfig
	ORM        *orm.Config
	MySQL      *sqlx.Config
	Log        *log.Config
	Tracer     *trace.Config
	Wechat     *wechat
	Cfg        *cfg
}

type wechat struct {
	Token    string
	Secret   string
	Username string
}

type cfg struct {
	MountDir          string
	DiskFileExpireDay int64
	LimitPerTask      int
	TaskGoroutines    int
	PartitionsURL     string
	UpimgURL          string
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func remote() (err error) {
	if client, err = conf.New(); err != nil {
		return
	}
	err = load()
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
