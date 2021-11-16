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
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"

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
	Log            *log.Config
	BM             *bm.ServerConfig
	Verify         *verify.Config
	Tracer         *trace.Config
	Redis          *redis.Config
	WhiteListRedis *redis.Config
	Memcache       *memcache.Config
	MySQL          *sql.Config
	Ecode          *ecode.Config
	DmRules        *dmRules
	FilterClient   *warden.ClientConfig
	AccClient      *warden.ClientConfig
	XuserClent     *warden.ClientConfig
	LocationClient *warden.ClientConfig
	SpyClient      *warden.ClientConfig
	BcastClient    *warden.ClientConfig
	UExpClient     *warden.ClientConfig
	IsAdminClient  *warden.ClientConfig
	HTTPClient     *bm.ClientConfig
	LiveRPC        map[string]*liverpc.ClientConfig
	BNDatabus      *databus.Config
	Lancer         *lancer
	CacheDatabus   *cache
	BNJRoomList    map[string]bool
}

type cache struct {
	Size int
}

type lancer struct {
	DMErr  *infoc.Config
	DMSend *infoc.Config
}

type dmRules struct {
	AllUserLimit     bool
	AreaLimit        bool
	LevelLimitStatus bool
	LevelLimit       int64
	PhoneLimit       bool
	RealName         bool
	MsgLength        int
	DmNum            int64
	DmPercent        int64
	Nixiang          map[string]bool
	Color            map[string]int64
	DMwhitelist      bool
	DMwhiteListID    string
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
