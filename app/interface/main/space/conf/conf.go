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
	"github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/supervisor"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/time"

	"github.com/namelessup/bilibili/library/database/hbase.v2"

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
	// elk
	Log *log.Config
	// App
	App *blademaster.App
	// tracer
	Tracer *trace.Config
	// Auth
	Auth *auth.Config
	// Verify
	Verify *verify.Config
	// Supervisor
	Supervisor *supervisor.Config
	// BM
	BM *httpServers
	// HTTPServer
	HTTPServer *blademaster.ServerConfig
	// Ecode
	Ecode *ecode.Config
	// ArchiveRPC
	AccountRPC  *rpc.ClientConfig
	ArticleRPC  *rpc.ClientConfig
	AssistRPC   *rpc.ClientConfig
	TagRPC      *rpc.ClientConfig
	FavoriteRPC *rpc.ClientConfig
	FilterRPC   *rpc.ClientConfig
	ThumbupRPC  *rpc.ClientConfig
	RelationRPC *rpc.ClientConfig
	MemberRPC   *rpc.ClientConfig
	// grpc
	AccClient  *warden.ClientConfig
	ArcClient  *warden.ClientConfig
	CoinClient *warden.ClientConfig
	UpClient   *warden.ClientConfig
	// Mysql
	Mysql *sql.Config
	// Redis
	Redis *redisConf
	// Mc
	Memcache *memConf
	// Rule
	Rule *rule
	// HTTP client
	HTTPClient *httpClient
	// Host
	Host *host
	// HBase hbase config
	HBase *Hbase
}

type redisConf struct {
	*redis.Config
	ClExpire    time.Duration
	UpArtExpire time.Duration
	UpArcExpire time.Duration
}

type memConf struct {
	*memcache.Config
	SettingExpire time.Duration
	NoticeExpire  time.Duration
	TopArcExpire  time.Duration
	MpExpire      time.Duration
	ThemeExpire   time.Duration
	TopDyExpire   time.Duration
}

type rule struct {
	MaxChNameLen     int
	MaxChIntroLen    int
	MaxChLimit       int
	MaxChArcLimit    int
	MaxChArcAddLimit int
	MaxChArcsPs      int
	MaxRiderPs       int
	MaxArticlePs     int
	ChIndexCnt       int
	MaxNoticeLen     int
	MaxTopReasonLen  int
	MaxMpReasonLen   int
	MaxMpLimit       int
	// RealNameOn
	RealNameOn bool
	// No limit notice mids
	NoNoticeMids []int64
	// default top photo
	TopPhoto string
	// dynamic list switch
	Merge   bool
	ActFold bool
	// block mids
	BlockMids []int64
	//BlackFre space blacklist frequency
	BlackFre time.Duration
}

type host struct {
	Bangumi string
	API     string
	Mall    string
	APIVc   string
	APILive string
	Acc     string
	Game    string
	AppGame string
	Search  string
	Elec    string
	Space   string
}

type httpClient struct {
	Read  *blademaster.ClientConfig
	Write *blademaster.ClientConfig
	Game  *blademaster.ClientConfig
}

type httpServers struct {
	Outer *blademaster.ServerConfig
}

// Hbase .
type Hbase struct {
	*hbase.Config
	ReadTimeout time.Duration
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
