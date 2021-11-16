package conf

import (
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/hbase.v2"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

const (
	configKey = "filter-admin.toml"
)

var (
	confPath string
	// client   *conf.Client
	// Conf .
	Conf *Config
)

// Config represent filter config
type Config struct {
	// MySQL .
	MySQL *sql.Config
	// HBase .
	HBase *HBaseConfig
	// Memcache mc.
	Memcache *Memcache
	// MultiHTTP http server.
	BM *bm.ServerConfig
	// Log log.
	Log *log.Config
	// Tracer
	Tracer *trace.Config
	// Property .
	Property *Property
	Ai       *Ai
	// HTTPClient .
	HTTPClient *HTTPClient
	// ecode
	Ecode *ecode.Config
	// Host
	Host *Host
	// Auth
	Auth *permit.Config
}

// HBaseConfig extra hbase config
type HBaseConfig struct {
	*hbase.Config
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Memcache cache.
type Memcache struct {
	Mc     *memcache.Config
	Expire *McExpire
}

// McExpire expore.
type McExpire struct {
	Expire time.Duration
}

// Expire expire.
type Expire struct {
	Expire time.Duration
}

// Property app properties
type Property struct {
	SourceMask []int64
	FilterType []int64
	Level      []int64
	// filter expired tick
	ExpiredTick time.Duration
	// 正常文本测试
	NormalContents []string
	// 危险文本测试
	RiskContents []string
	// 正常文本失效阈值
	NormalHitRate int
}

// Ai struct
type Ai struct {
	// AI阀值
	Threshold float64
	// AI真实分标准
	TrueScore float64
}

// HTTPClient conf.
type HTTPClient struct {
	Off          bool
	SearchDomain string
	Normal       *bm.ClientConfig
}

// Host is Host config
type Host struct {
	AI  string
	API string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf
func Init() (err error) {
	if confPath == "" {
		return configCenter()
	}
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func configCenter() (err error) {
	var (
		client *conf.Client
		value  string
		ok     bool
	)
	if client, err = conf.New(); err != nil {
		return
	}
	if value, ok = client.Value(configKey); !ok {
		panic(err)
	}
	_, err = toml.Decode(value, &Conf)
	return
}
