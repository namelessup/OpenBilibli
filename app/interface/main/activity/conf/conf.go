package conf

import (
	"errors"
	"flag"
	xtime "time"

	"github.com/namelessup/bilibili/app/interface/main/activity/model/bnj"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/stat/prom"
	"github.com/namelessup/bilibili/library/time"

	"github.com/namelessup/bilibili/library/database/hbase.v2"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf is global config
	Conf = &Config{}
)

// Config service config
type Config struct {
	Static string
	// reload
	Reload ReloadInterval
	// auth
	Auth *auth.Config
	// verify
	Verify *verify.Config
	// HTTPServer
	HTTPServer *blademaster.ServerConfig
	// tracer
	Tracer *trace.Config
	// db
	MySQL *MySQL
	// rpc
	RPCClient2 *RPCClient2
	// grpc
	TagClient *warden.ClientConfig
	// acc client
	AccClient *warden.ClientConfig
	// httpClient
	HTTPClient *blademaster.ClientConfig
	// HTTPClientSports
	HTTPClientSports *blademaster.ClientConfig
	// HTTPClientBnj
	HTTPClientBnj *blademaster.ClientConfig
	// HTTPClientKfc
	HTTPClientKfc *blademaster.ClientConfig
	// Rule
	Rule *Rule
	// Host
	Host Host
	// Log
	Log *log.Config
	// ecode
	Ecode *ecode.Config
	// ip
	IPFile string
	// mc
	Memcache *Memcache
	TimeMc   *tmMC
	// redis
	Redis *Redis
	// hbase
	Hbase     *hbase.Config
	RPCServer *rpc.ServerConfig
	// interval
	Interval *Interval
	// Elastic
	Elastic *elastic.Config
	// ArcClient
	ArcClient *warden.ClientConfig
	// Time machine conf
	Timemachine *timemachine
	// Bnj
	Bnj2019 *bnj2019
	// databus
	Databus *Databus
}

// Host remote host.
type Host struct {
	Sports   string
	QqNews   string
	Activity string
	APICo    string
	Mall     string
	LiveCo   string
}

// Rule   rule config.
type Rule struct {
	GuessCount      int
	MaxGuessCoin    int64
	SuitPids        []int64
	SuitExpire      int64
	TickQq          time.Duration
	QqTryCount      int
	DTimeout        time.Duration
	QqStartTime     string
	QqEndTime       string
	QqYear          string
	PlayerYear      string
	BwsMids         []int64
	BwsAwardMids    []int64
	BwsLotteryMids  []int64
	BwsLotteryAids  []int64
	BwsSuitExpire   int64
	NeedInitAchieve bool
	DialectTags     []int64
	DialectRegions  []int16
	DialectSid      int64
	SpecialSids     []int64
	Spylike         int64
	LotteryActID    int64
	MatchLotteryID  int64
	S8Sid           int64
	S8ArcSid        int64
	S8ArtSid        int64
	KingStorySid    int64
	TmMids          []int64
}

// Interval .
type Interval struct {
	NewestSubTsInterval time.Duration
	PullArcTypeInterval time.Duration
	ActSourceInterval   time.Duration
	TmInternal          time.Duration
}

// Prom prom .
type Prom struct {
	LIBClient      *prom.Prom
	LIBClientState *prom.Prom
	APIClient      *prom.Prom
	HTTPServer     *prom.Prom
}

// MySQL define MySQL config
type MySQL struct {
	Like *sql.Config
}

// ReloadInterval define reolad config
type ReloadInterval struct {
	Jobs   time.Duration
	Notice time.Duration
	Ad     time.Duration
}

// RPCClient2 define RPC client config
type RPCClient2 struct {
	Archive *rpc.ClientConfig
	Coin    *rpc.ClientConfig
	Suit    *rpc.ClientConfig
	Spy     *rpc.ClientConfig
	Tag     *rpc.ClientConfig
	Thumbup *rpc.ClientConfig
	Article *rpc.ClientConfig
}

// Redis struct
type Redis struct {
	*redis.Config
	Expire          time.Duration
	MatchExpire     time.Duration
	FollowExpire    time.Duration
	UserAchExpire   time.Duration
	UserPointExpire time.Duration
	AchCntExpire    time.Duration
	HotDotExpire    time.Duration
	RandomExpire    time.Duration
	ResetExpire     time.Duration
	RewardExpire    time.Duration
}

// Memcache struct
type Memcache struct {
	Like             *memcache.Config
	LikeExpire       time.Duration
	LikeIPExpire     time.Duration
	PerpetualExpire  time.Duration
	ItemExpire       time.Duration
	SubStatExpire    time.Duration
	ViewRankExpire   time.Duration
	SourceItemExpire time.Duration
	QqExpire         time.Duration
	BwsExpire        time.Duration
	ProtocolExpire   time.Duration
	KfcExpire        time.Duration
	KfcCodeExpire    time.Duration
}

type tmMC struct {
	Timemachine *memcache.Config
	TmExpire    time.Duration
}

type timemachine struct {
	TagDescID       int64
	TagRegionDescID int64
	RegionDescID    int64
}

type bnj2019 struct {
	ActID         int64
	SubID         int64
	GameCancel    int64
	AdminCheck    int64
	Admins        []int64
	TimelinePic   string
	H5TimelinePic string
	Start         xtime.Time
	Reward        []*bnj.Reward
	Info          []*struct {
		Nav      string
		Pic      string
		H5Pic    string
		Aid      int64
		Detail   string
		H5Detail string
		Nickname string
		Publish  xtime.Time
	}
}

// Databus .
type Databus struct {
	Bnj *databus.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

// Init conf.
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
}

func remote() (err error) {
	if client, err = conf.New(); err != nil {
		return
	}
	if err = load(); err != nil {
		return
	}
	client.Watch("activity.toml")
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
