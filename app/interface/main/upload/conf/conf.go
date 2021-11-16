package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/app/interface/main/upload/http/antispam"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf conf
	Conf = &Config{}
)

// Config config
type Config struct {
	XLog *log.Config
	// BM
	BM *bm.ServerConfig
	// ecode
	Ecode *ecode.Config
	// orm
	ORM *orm.Config
	// bfs
	Bfs       *Bfs
	BfsBucket *BfsBucket
	Auths     []*Auth
	// Antispam redis
	Antispam *antispam.Config
	// AuthN
	AuthInter *auth.Config
	// VerifyN
	Verify *verify.Config
	// AuthN outside
	AuthOut *auth.Config
}

// Bfs .
type Bfs struct {
	BfsURL          string
	WaterMarkURL    string
	ImageGenURL     string
	TimeOut         xtime.Duration
	WmTimeOut       xtime.Duration
	ImageGenTimeOut xtime.Duration
}

// BfsBucket .
type BfsBucket struct {
	Bucket string
	Key    string
	Sercet string
}

// Auth .
type Auth struct {
	AppKey    string
	AppSercet string
	BfsBucket *BfsBucket
}

// Antispam .
type Antispam struct {
	Redis  *redis.Config
	Switch bool
	Second int
	N      int
	Hour   int
	M      int
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
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
