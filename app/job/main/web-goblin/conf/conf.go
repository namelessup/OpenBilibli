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
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// Config .
type Config struct {
	Log      *log.Config
	BM       *bm.ServerConfig
	Tracer   *trace.Config
	Redis    *redis.Config
	Memcache *memcache.Config
	DB       *DB
	Ecode    *ecode.Config
	// rpc
	DynamicRPC  *rpc.ClientConfig
	FavoriteRPC *rpc.ClientConfig
	ArchiveRPC  *rpc.ClientConfig
	// Host
	Host Host
	// HTTP client
	HTTPClient        *bm.ClientConfig
	MessageHTTPClient *bm.ClientConfig
	// Rule
	Rule *Rule
	//Push push urls
	Push *Push
	//Message
	Message Message
	// App
	App *bm.App
	// databus
	ArchiveNotifySub *databus.Config
	// Warden Client
	ArcClient *warden.ClientConfig
}

// Push push.
type Push struct {
	BusinessID    int
	BusinessToken string
	PartSize      int
	RetryTimes    int
	Title         string
	BodyDefault   string
	BodySpecial   string
	OnlyMids      string
}

// Message .
type Message struct {
	URL string
	MC  string
}

// Rule .
type Rule struct {
	BroadFeed        int
	SleepInterval    time.Duration
	Before           time.Duration
	ScoreSleep       time.Duration
	AlertTitle       string
	AlertBodyDefault string
	AlertBodySpecial string
	CoinPercent      float64
	FavPercent       float64
	DmPercent        float64
	ReplyPercent     float64
	ViewPercent      float64
	LikePercent      float64
	SharePercent     float64
	NewDay           float64
	NewPercent       float64
}

// Host remote host
type Host struct {
	API string
}

// DB define MySQL config
type DB struct {
	Esports *sql.Config
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
