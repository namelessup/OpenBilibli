package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/antispam"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
	"github.com/namelessup/bilibili/app/service/main/videoup/model/archive"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

var (
	confPath string
	// Conf is global config
	Conf   = &Config{}
	client *conf.Client
)

// Config service config
type Config struct {
	// base
	// host
	Host *Host
	// channal len
	ChanSize int64
	// ecode
	Ecode *ecode.Config
	// Xlog
	Xlog *log.Config
	// tracer
	Tracer *trace.Config
	// tick load pgc
	Tick time.Duration
	// db
	DB *DB
	// databus
	VideoupPub    *databus.Config
	VideoupPGCPub *databus.Config
	// redis
	Redis    *Redis
	Memcache *Memcache
	// http client test
	HTTPClient *HTTPClient
	// keep Aids
	KeepArc *KeepArc
	// DmVerifyKey dm_key
	DmVerifyKey string
	// Monitor
	Monitor *Monitor
	// PubAgent
	PubAgent *PubAgent
	// AsyncThreshold
	AsyncThreshold  int
	SplitThreshold  int
	SplitGroupCount int
	FailThreshold   int
	EditTimeout     time.Duration
	GrayGroup       int
	//BM bladermaster config
	BM *bm.ServerConfig
	//AntispamConf limit request
	AntispamConf *antispam.Config
	Property     *Property
	// rpc
	AccRPC *warden.ClientConfig
}

// Property .
type Property struct {
	MSG []*archive.MSG
}

// Memcache conf.
type Memcache struct {
	Archive struct {
		*memcache.Config
		TplExpire time.Duration
	}
}

// Monitor  define sms monitor conf
type Monitor struct {
	Tels  string
	Env   string
	Count int64
}

// PubAgent struct
type PubAgent struct {
	Proxy         string
	PGCSubmit     string
	PGCDRMSubmit  string
	UGCSubmit     string
	UGCFirstRound string
}

// Host define host info
type Host struct {
	Mission string
	Account string
	Monitor string
	APICO   string
	MSG     string
	Manager string
}

// DB define MySQL config
type DB struct {
	Archive        *sql.Config
	ArchiveRead    *sql.Config
	ArchiveSlave   *sql.Config
	CreativeCenter *sql.Config
	Dede           *sql.Config
	Manager        *sql.Config
	Oversea        *sql.Config
}

// Redis define redis config
type Redis struct {
	Track *TrackRedis
}

// TrackRedis define track redis config
type TrackRedis struct {
	*redis.Config
	Expire time.Duration
}

// HTTPClient test
type HTTPClient struct {
	Read  *bm.ClientConfig
	Write *bm.ClientConfig
}

// KeepArc keep archive to mid
type KeepArc struct {
	Aids []int64
	Mid  int64
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init conf
func Init() (err error) {
	if confPath != "" {
		_, err = toml.DecodeFile(confPath, &Conf)
		return
	}
	err = remote()
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
