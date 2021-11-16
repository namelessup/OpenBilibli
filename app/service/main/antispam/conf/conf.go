package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

const (
	configKey = "antispam-service.toml"
)

var (
	// Conf .
	Conf *Config
	// ConfPath .
	ConfPath string
)

// Config .
type Config struct {
	RPC                *rpc.ServerConfig
	App                *bm.App
	BM                 *bm.ServerConfig
	HTTPClient         *bm.ClientConfig
	MySQL              *MySQL
	Redis              *Redis
	Tracer             *trace.Config
	Log                *log.Config
	Verify             *verify.Config
	Auth               *auth.Config
	Ecode              *ecode.Config
	AppkeyType         map[string][]int8
	ReplyURL           string
	ServiceOption      *ServiceOption
	MaxSpawnGoroutines int
	MaxAllowedCounts   int64
	MaxDurationSec     int64
	AutoWhite          *AutoWhite
}

// AutoWhite .
type AutoWhite struct {
	KeywordHitCounts int64
	NumOfSenders     int64
	Derivation       float64
}

// ServiceOption .
type ServiceOption struct {
	GcOpt                    *GcOpt
	BuildTrieIntervalMinute  int64
	BuildTrieMaxRowsPerQuery int64
	AsyncTaskChanSize        int64

	RefreshTrieIntervalSec    int64
	RefreshRulesIntervalSec   int64
	RefreshRegexpsIntervalSec int64

	MinKeywordLen          int
	MaxSenderNum           int64
	DefaultExpireSec       int64
	DefaultChanSize        int64
	MaxExportRows          int64
	MaxRegexpCountsPerArea int64
	MaxSpawnGoroutines     int64

	RuleDefaultExpireSec   int64
	RegexpDefaultExpireSec int64
}

// GcOpt .
type GcOpt struct {
	Open            bool
	IntervalSec     int
	MaxRowsPerQuery int64
}

// MySQL .
type MySQL struct {
	AntiSpam *sql.Config
}

// Redis .
type Redis struct {
	*redis.Config
	IndexExpire time.Duration
}

func init() {
	flag.StringVar(&ConfPath, "conf", "", "config path")
}

// Init .
func Init(path string) error {
	if len(Areas) == 0 {
		panic(errors.New("areas must be set"))
	}
	if path == "" {
		return configCenter()
	}
	_, err := toml.DecodeFile(path, &Conf)
	return err
}

func configCenter() error {
	client, err := conf.New()
	if err != nil {
		return err
	}
	value, ok := client.Value(configKey)
	if !ok {
		return errors.New("empty value")
	}
	_, err = toml.Decode(value, &Conf)
	return err
}

// Areas .
var Areas = map[string]int{
	"reply":   1,
	"im":      2,
	"live_dm": 3,
	"danmu":   4,
}
