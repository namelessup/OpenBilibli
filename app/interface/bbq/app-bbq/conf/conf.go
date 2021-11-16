package conf

import (
	"github.com/namelessup/bilibili/app/interface/bbq/app-bbq/api/http/v1"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf/paladin"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/antispam"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
)

var (
	// Conf config
	Conf = &Config{}
	// App setting
	App = &AppSetting{}
	// Filter .
	Filter = &UploadFilter{}
)

// Config .
type Config struct {
	Log        *log.Config
	BM         *bm.ServerConfig
	Verify     *verify.Config
	Auth       *auth.Config
	Tracer     *trace.Config
	Redis      *redis.Config
	MySQL      *sql.Config
	DMMySQL    *sql.Config
	Ecode      *ecode.Config
	HTTPClient *HTTPClient
	GRPCClient map[string]*GRPCConf
	AntiSpam   map[string]*antispam.Config
	Tmap       map[string]string
	URLs       map[string]string
	Comment    *Comment
	Infoc      *infoc.Config
	Search     *Search
	Notices    []*v1.NoticeOverview
	Upload     *Upload
}

//Upload ..
type Upload struct {
	HTTPSchema string
}

// Set .
func (c *Config) Set(text string) error {
	if _, err := toml.Decode(text, c); err != nil {
		panic(err)
	}
	if c.Redis != nil {
		for _, anti := range c.AntiSpam {
			anti.Redis = c.Redis
		}
	}
	return nil
}

// Comment 评论配置
type Comment struct {
	Type       int64
	DebugID    int64
	CloseRead  bool
	CloseWrite bool
}

// Search 搜索配置
type Search struct {
	Host string
}

// HTTPClient conf
type HTTPClient struct {
	Normal *bm.ClientConfig
	Slow   *bm.ClientConfig
}

//GRPCConf .
type GRPCConf struct {
	WardenConf *warden.ClientConfig
	Addr       string
}

// Init init conf
func Init() (err error) {
	if err = paladin.Init(); err != nil {
		return
	}
	if err = paladin.Watch("video-c.toml", Conf); err != nil {
		return
	}
	if err = paladin.Watch("app_setting.toml", App); err != nil {
		return
	}
	if err = paladin.Watch("upload_filter.toml", Filter); err != nil {
		return
	}
	return
}
