package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// global var
var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// Config config set
type Config struct {
	// base
	// elk
	Log *log.Config
	// mc
	Memcache *Memcache
	// auth
	Auth *permit.Config
	// MultiHTTP
	BM *bm.ServerConfig
	// http client
	HTTPClient *bm.ClientConfig
	// tracer
	Tracer *trace.Config
	// MySQL
	MySQL *sql.Config
	ORM   *ORM
	// VipRPC
	VipRPC *rpc.ClientConfig
	// business config
	Property *Property
	// ecode
	Ecode *ecode.Config
	// pay conf
	PayConf *PayConf
	// bfs
	Bfs *Bfs
}

// Bfs reprensents the bfs config
type Bfs struct {
	Key         string
	Secret      string
	Host        string
	Timeout     int
	MaxFileSize int
}

// ORM .
type ORM struct {
	Vip *orm.Config
}

// Memcache .
type Memcache struct {
	*memcache.Config
	Expire time.Duration
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Property .
type Property struct {
	MsgURI                         string
	PayURI                         string
	AnnualVipBcoinDay              int
	AnnualVipBcoinCouponMoney      int
	AnnualVipBcoinCouponActivityID int
	WelfareBgHost                  string
}

//PayConf pay config
type PayConf struct {
	BaseURL        string
	CustomerID     int64
	Token          string
	OrderNotifyURL string
	SignNotifyURL  string
	RefundURL      string
	PlanID         int32
	ProductID      string
	Version        string
	ReturnURL      string
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
