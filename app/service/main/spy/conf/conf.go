package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/namelessup/bilibili/library/net/rpc/warden"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf represents a config for spy service.
	Conf = &Config{}
)

// Config def.
type Config struct {
	Account string
	// tracer
	Tracer *trace.Config
	// http client
	HTTPClient *bm.ClientConfig
	// db
	DB *DB
	// rpc server2
	RPCServer *rpc.ServerConfig
	// memcache
	Memcache *Memcache
	// log
	Log *log.Config
	// rpc clients
	RPC *RPC
	// biz property.
	Property *Property
	// redis
	Redis *Redis
	// databus
	DBScoreChange *databus.Config
	// qcloud
	Qcloud *Qcloud
	// bm
	BM *bm.ServerConfig
	// grpc
	GRPC *warden.ServerConfig
}

// DB config.
type DB struct {
	Spy *sql.Config
}

// Redis redis.
type Redis struct {
	*redis.Config
	Expire        xtime.Duration
	VerifyCdTimes xtime.Duration
}

// RPC clients config.
type RPC struct {
	Account *rpc.ClientConfig
}

// Memcache config.
type Memcache struct {
	User       *memcache.Config
	UserExpire time.Duration
}

// Property config for biz logic.
type Property struct {
	TelValidateURL      string
	BlockAccountURL     string
	SecurityLoginURL    string
	TelInfoByMidURL     string
	ProfileInfoByMidURL string
	UnicomGiftStateURL  string
	LoadEventTick       xtime.Duration
	DoubleCheckLevel    int32
	ConfigLoadTick      xtime.Duration
	UserInfoShard       int64
	HistoryShard        int64
	AutoBlockSwitch     bool
	Score               *struct {
		BaseInit  int8
		EventInit int8
	}
	Punishment *struct {
		ScoreThreshold int8
		Times          int8
	}
	Event *struct {
		ServiceName           string
		InitEventID           int64
		BindMailAndTelLowRisk string
		BindMailOnly          string
		BindNothing           string
		BindTelLowRiskOnly    string
		BindTelMediumRisk     string
		BindTelHighRisk       string
		BindTelUnknownRisk    string

		BindTelLowRiskAndIdenAuth       string
		BindTelLowRiskAndIdenUnauth     string
		BindTelUnknownRiskAndIdenAuth   string
		BindTelMediumRiskAndIdenAuth    string
		BindTelUnknownRiskAndIdenUnauth string
		BindTelMediumRiskAndIdenUnauth  string
		BindMailAndIdenUnknown          string
		BindTelHighRiskAndIdenAuth      string
		BindNothingV2                   string
		BindNothingAndIdenAuth          string
		BindTelHighRiskAndIdenUnauth    string
	}
	Block *struct {
		CycleTimes int64 // unit per seconds
	}
	White *struct {
		Tels []struct {
			From int64 // <= from
			To   int64 // >= to
		}
	}
}

// Qcloud def.
type Qcloud struct {
	Path      string
	Region    string
	SecretID  string
	SecretKey string
	Charset   string
	BaseURL   string
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
