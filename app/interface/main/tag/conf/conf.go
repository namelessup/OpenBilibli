package conf

import (
	"errors"
	"flag"
	"time"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/stat/prom"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf config .
	Conf = &Config{}
)

// Config config .
type Config struct {
	Log         *log.Config
	Ecode       *ecode.Config
	Tag         *Tag
	Supervision *Supervision
	Host        *Host
	Tracer      *trace.Config
	Auth        *auth.Config
	Verify      *verify.Config
	HTTPClient  *bm.ClientConfig
	HTTPSimilar *bm.ClientConfig
	BM          *bm.ServerConfig
	GRPCServer  *warden.ServerConfig
	RPCServer   *rpc.ServerConfig
	Redis       *Redis
	ArchiveRPC  *rpc.ClientConfig
	TagDisRPC   *rpc.ClientConfig
	FigureRPC   *rpc.ClientConfig
	// Warden Client
	TagGRPClient *warden.ClientConfig
	AccGRPClient *warden.ClientConfig
}

// Host host config .
type Host struct {
	APICo      string
	AI         string
	Account    string
	Archive    string
	BigDataURL string
}

// Redis redis  config .
type Redis struct {
	Tag  *TagRedis
	Rank *RankRedis
}

// TagRedis tag redis  config .
type TagRedis struct {
	Redis  *redis.Config
	Expire *TagExpire
}

// TagExpire expire  config .
type TagExpire struct {
	Sub      xtime.Duration
	ArcTag   xtime.Duration
	ArcTagOp xtime.Duration
	AtLike   xtime.Duration
	AtHate   xtime.Duration
}

// RankRedis rank redis  config .
type RankRedis struct {
	Redis  *redis.Config
	Expire *RankExpire
}

// RankExpire rang expire config .
type RankExpire struct {
	TagNewArc xtime.Duration
}

// Tag tag config .
type Tag struct {
	FeedBackMaxLen int
	// user level
	ArcTagAddLevel  int
	ArcTagDelLevel  int
	ArcTagRptLevel  int
	ArcTagLikeLevel int
	ArcTagHateLevel int

	SubArcMaxNum int
	// arctag
	ArcTagMaxNum     int
	ArcTagAddMaxNum  int
	ArcTagDelMaxNum  int
	ArcTagDelSomeNum int
	ArcTagLikeMaxNum int
	ArcTagHateMaxNum int
	ArcTagRptMaxNum  int
	LikeLimitToLock  int64

	MaxArcsPageSize int
	MaxArcsLimit    int
	// select tag number
	MaxSelTagNum       int
	White              []int64 // 用户账号白名单
	ChannelRefreshTime xtime.Duration
	AITimeout          int
}

// Supervision supervision .
type Supervision struct {
	SixFour *struct {
		Button bool
		Begin  time.Time
		End    time.Time
	}
	RealName *struct {
		Button bool
	}
}

// PromError stat and log.
func PromError(name string, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init intt conf .
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
	if value, ok = client.Toml2(); !ok {
		return errors.New("load config center error")
	}
	_, err = toml.Decode(value, &Conf)
	return
}
