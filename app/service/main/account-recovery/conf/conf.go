package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/elastic"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	xtime "github.com/namelessup/bilibili/library/time"

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
	Log    *log.Config
	BM     *bm.ServerConfig
	Verify *verify.Config
	Tracer *trace.Config
	Redis  *redis.Config
	MySQL  *sql.Config
	Ecode  *ecode.Config

	// HTTPClientConfig
	HTTPClientConfig *bm.ClientConfig
	// AccRecover  request URL info
	AccRecover *AccRecover
	// MailConfig
	MailConf *Mail
	// CaptchaConf
	CaptchaConf *Captcha

	// RPC config
	LocationRPC *rpc.ClientConfig

	// grpc
	MemberGRPC  *warden.ClientConfig
	AccountGRPC *warden.ClientConfig

	// ChanSize
	ChanSize *ChanSize
	// Auth
	Auth *permit.Config

	AESEncode *AESEncode

	// elastic config
	Elastic *elastic.Config

	// Bfs
	Bfs *Bfs
	// DataBus databus
	DataBus *DataBus
}

// ChanSize mail send channel size.
type ChanSize struct {
	MailMsg int64
}

// AccRecover is a url config to request java api
type AccRecover struct {
	MidInfoURL         string
	UpPwdURL           string
	UpBatchPwdURL      string
	CheckSafeURL       string
	GameURL            string
	CheckRegURL        string
	CheckUserURL       string
	CheckCardStatusURL string
	CheckCardURL       string
	CheckPwdURL        string
	GetLoginIPURL      string
	GetUserInfoURL     string
}

// Mail 邮件配置
type Mail struct {
	Host               string
	Port               int
	Username, Password string
}

// Captcha 验证码配置
type Captcha struct {
	TokenBID  string
	TokenURL  string
	VerifyURL string
}

// AESEncode aes encode
type AESEncode struct {
	AesKey string
	Salt   string
}

// Bfs Bfs.
type Bfs struct {
	Timeout xtime.Duration
	Bucket  string
	Addr    string
	Key     string
	Secret  string
}

// DataBus is
type DataBus struct {
	UserActLog *databus.Config
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
