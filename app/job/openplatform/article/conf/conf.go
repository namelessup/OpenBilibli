package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	xlog "github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// Config .
type Config struct {
	// Env
	Env string
	// App
	App *bm.App
	// Xlog is github.com/namelessup/bilibili log.
	Xlog *xlog.Config
	// Tracer .
	Tracer *trace.Config
	// ArchiveSub databus
	ArticleSub *databus.Config
	// ArticleStatSub databus
	ArticleStatSub *databus.Config
	// LikeStatSub databus
	LikeStatSub     *databus.Config
	ReplyStatSub    *databus.Config
	FavoriteStatSub *databus.Config
	CoinStatSub     *databus.Config
	// DynamicDbus pub databus
	DynamicDbus *databus.Config
	// BM
	BM *bm.ServerConfig
	// HTTPClient .
	HTTPClient     *bm.ClientConfig
	GameHTTPClient *bm.ClientConfig
	// RPC .
	ArticleRPC *rpc.ClientConfig
	TagRPC     *rpc.ClientConfig
	// DB
	DB *sql.Config
	// Redis
	Redis *redis.Config
	// SMS text message.
	SMS *sms
	// CheatInfoc
	CheatInfoc *infoc.Config
	// ReadInfoc
	ReadInfoc *infoc.Config
	// article interface redis
	ArtRedis *redis.Config
	// Job params
	Job *job
	// Sitemap
	Sitemap Sitemap
}

// Sitemap .
type Sitemap struct {
	Interval int64
	Size     int
}

type job struct {
	ViewCacheTTL          xtime.Duration
	DupViewCacheTTL       xtime.Duration
	UpdateDbInterval      xtime.Duration
	UpdateSortInterval    xtime.Duration
	GameCacheExpire       xtime.Duration
	ListReadCountInterval xtime.Duration
	HotspotInterval       xtime.Duration
	HotspotForceInterval  xtime.Duration
	ExpireSortArts        xtime.Duration
	TTLSortArts           xtime.Duration
	SortLimitTime         xtime.Duration
	RecommendExpire       xtime.Duration
	Words                 int64
	StatDays              int64
	ActLikeURL            string
	FlowURL               string
	MaxNewArtsNum         int64
	MaxSortArtsNum        int64
}

type sms struct {
	Phone string
	Token string
}

var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init .
func Init() (err error) {
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
			xlog.Info("config reload")
			if load() != nil {
				xlog.Error("config reload err")
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
