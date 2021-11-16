package conf

import (
	"flag"
	"github.com/namelessup/bilibili/library/database/elastic"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	// ConfPath config path.
	ConfPath string
	// Conf config.
	Conf *Config
)

// Config config.
type Config struct {
	XLog   *log.Config
	Tracer *trace.Config
	// HTTP
	HTTPClient        *bm.ClientConfig
	DrawyooHTTPClient *bm.ClientConfig
	// databus
	Databus *Databus
	// rpc
	RPCClient2 *RPCClient2
	// bm
	BM *bm.ServerConfig
	// mysql
	MySQL *MySQL
	// redis
	Redis *Redis
	// mc
	Memcache *Memcache
	// Host
	Host          *Host
	Weight        *Weight
	Job           *Job
	StatTypes     map[string]int8
	Es            *elastic.Config
	AccountClient *warden.ClientConfig
}

// Job job.
type Job struct {
	Proc        int
	SearchNum   int
	SearchFlush time.Duration
	MessageMids []int64
	BatchNumber int
}

// Weight weight.
type Weight struct {
	Like int
	Hate int
}

// Host represents host info.
type Host struct {
	Activity  string
	Message   string
	DrawYoo   string
	Search    string
	BlackRoom string
	LiveVC    string
	LiveAct   string
	API       string
	Bangumi   string
}

// MySQL mysql.
type MySQL struct {
	Reply *sql.Config
}

// Redis redis.
type Redis struct {
	*redis.Config
	IndexExpire     time.Duration
	ReportExpire    time.Duration
	UserCntExpire   time.Duration
	StatCacheExpire time.Duration
	UserActExpire   time.Duration
	NotifyExpire    time.Duration
}

// Memcache mc.
type Memcache struct {
	*memcache.Config
	Expire    time.Duration
	TopExpire time.Duration
}

// RPCClient2 rpc client.
type RPCClient2 struct {
	Account *rpc.ClientConfig
	Archive *rpc.ClientConfig
	Article *rpc.ClientConfig
	Assist  *rpc.ClientConfig
}

// Databus databus.
type Databus struct {
	Event    *databus.Config
	Stats    *databus.Config
	Consumer *databus.Config
	Like     *databus.Config
}

// Stats stats.
type Stats struct {
	*databus.Config
	Type  int8
	Field string
}

func init() {
	flag.StringVar(&ConfPath, "conf", "", "config path")
}

// Init init conf
func Init() (err error) {
	if ConfPath == "" {
		return configCenter()
	}
	_, err = toml.DecodeFile(ConfPath, &Conf)
	return
}

func configCenter() (err error) {
	var (
		ok     bool
		value  string
		client *conf.Client
	)
	if client, err = conf.New(); err != nil {
		return
	}
	if value, ok = client.Toml2(); !ok {
		panic(err)
	}
	_, err = toml.Decode(value, &Conf)
	return
}
