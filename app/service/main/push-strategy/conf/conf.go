package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	xsql "github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	xhttp "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// Conf info.
var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// Config struct.
type Config struct {
	Ecode      *ecode.Config
	Verify     *verify.Config
	MySQL      *xsql.Config
	Log        *log.Config
	HTTPServer *xhttp.ServerConfig
	HTTPClient *xhttp.ClientConfig
	FilterRPC  *rpc.ClientConfig
	Tracer     *trace.Config
	Redis      *rds
	Memcache   *mc
	Wechat     *wechat
	Cfg        *cfg
	BizID      *bizid
}

type rds struct {
	*redis.Config
	LimitDayExpire xtime.Duration
}

type mc struct {
	*memcache.Config
	UUIDExpire xtime.Duration
	CDExpire   xtime.Duration
}

type wechat struct {
	Token    string
	Secret   string
	Username string
}

type bizid struct {
	Live    int
	Archive int
}

type cfg struct {
	LoadTaskInteval      xtime.Duration
	LoadBusinessInteval  xtime.Duration
	LoadSettingsInteval  xtime.Duration
	NASPath              string
	LimitUserPerDay      int
	HandleTaskGoroutines int
	HandleMidGoroutines  int
	CacheSize            int
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
