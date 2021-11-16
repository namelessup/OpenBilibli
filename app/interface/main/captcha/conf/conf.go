package conf

import (
	"errors"
	"flag"
	"image/color"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/rate"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf .
	Conf = &Config{}
)

// Config captcha service config struct.
type Config struct {
	XLog     *log.Config
	Tracer   *trace.Config
	Ecode    *ecode.Config
	BM       *HTTPServers
	Verify   *verify.Config
	Rate     *rate.Config
	Memcache *Memcache
	Captcha  *Captcha
	Business []*Business
}

// Memcache represent mc conf
type Memcache struct {
	*memcache.Config
	Expire time.Duration
}

// HTTPServers Http Servers
type HTTPServers struct {
	Outer *bm.ServerConfig
}

// Business third business confs.
type Business struct {
	BusinessID string
	LenStart   int
	LenEnd     int
	Width      int
	Length     int
	TTL        time.Duration
}

// Captcha captcha service conf.
type Captcha struct {
	OuterHost    string
	Capacity     int
	DisturbLevel int    // 4 normal, 8 medium, 16 high
	Ext          string // jpeg
	Fonts        []string
	BkgColors    []color.RGBA
	FrontColors  []color.RGBA
	Interval     time.Duration
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init captcha service init.
func Init() (err error) {
	if confPath == "" {
		return configCenter()
	}
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

// configCenter connect to config center, get configs.
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
