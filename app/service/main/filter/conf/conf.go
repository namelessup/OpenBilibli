package conf

import (
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/hbase.v2"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

const (
	configKey = "filter-service.toml"
)

var (
	confPath string
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
	// BM
	BM *bm.ServerConfig
	// RPCServer rpc server.
	RPCServer *rpc.ServerConfig
	// grpc server
	WardenServer *warden.ServerConfig
	// Log log.
	Log *log.Config
	// Tracer
	Tracer *trace.Config
	// Property .
	Property *Property
	// http client
	HTTPClient *bm.ClientConfig
	Infoc      *infoc.Config
}

// HBaseConfig .
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
	FilterKeyExpire time.Duration
	FilterExpire    time.Duration
}

// Expire expire.
type Expire struct {
	Expire time.Duration
}

// Property app properties
type Property struct {
	SourceMask []int64
	// reload
	ReloadTick       time.Duration
	SecondReloadTick time.Duration
	// lru reload
	LruTick time.Duration
	// lru length
	LruLen int
	// parallelSize
	ParallelSize int
	// filter stage tick
	ExpiredTick time.Duration
	// it means <= criticalFilterLevel won't replace "*" , in the other will replace it.
	CriticalFilterLevel int8
	// cpu并发数
	GoMaxProce int
	// 文本分片最大长度
	MaxMpostSplitSize int
	MaxHitSplitSize   int
	// 分片过滤url
	FilterMpostURL string
	FilterHitURL   string
	// 短文本cache最大长度 (Byte)
	FilterCacheShortMaxSize int
	// 长文本cache最小长度 (Byte)
	FilterCacheLongMinSize int
	// 全level过滤名单
	FilterFullLevelList []string

	AI *struct {
		// AI分值，如果>0则以此阀值
		Threshold float64
		// AI真实分标准
		TrueScore float64
	}
	AIHost *struct {
		AI      string
		API     string
		Manager string
	}
	AIDelayTick time.Duration
}

// HTTPClient conf.
type HTTPClient struct {
	Off          bool
	SearchDomain string
	Normal       *bm.ClientConfig
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
