package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
)

// Conf global var
var (
	Conf     = &Config{}
	confPath string
	client   *conf.Client
)

// Config .
type Config struct {
	Log      *log.Config
	BM       *bm.ServerConfig
	Verify   *verify.Config
	Tracer   *trace.Config
	Memcache *memcache.Config
	Ecode    *ecode.Config
	Seo      *Seo
	Pages    []*Page
	Sitemaps []*Sitemap
}

// Seo config
type Seo struct {
	Expire  int32
	MaxAge  int32
	BotList []string
}

// Page pro, item ...
type Page struct {
	Name string
	Url  string
	Bfs  string
	Path string
}

// Sitemap app sitemap
type Sitemap struct {
	Host string
	Url  string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init .
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

// GetPage get page config by name
func GetPage(name string) *Page {
	for _, p := range Conf.Pages {
		if p.Name == name {
			return p
		}
	}
	return nil
}

// GetSitemap get sitemap url
func GetSitemap(host string) *Sitemap {
	for _, s := range Conf.Sitemaps {
		if s.Host == host {
			return s
		}
	}
	return nil
}
