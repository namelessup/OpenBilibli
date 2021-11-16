package conf

import (
	"errors"
	"flag"

	vipverify "github.com/namelessup/bilibili/app/service/main/vip/verify"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/antispam"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/supervisor"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// Conf global variable.
var (
	Conf     = &Config{}
	confPath string
	client   *conf.Client
)

// Config struct of conf.
type Config struct {
	Xlog                 *log.Config
	Tracer               *trace.Config
	BM                   *HTTPServers
	AuthN                *auth.Config
	Verify               *verify.Config
	RPCClient2           *RPC
	HTTPClient           *HTTPClient
	App                  *bm.App
	Ecode                *ecode.Config
	Antispam             *antispam.Config
	BatchRelAntispam     *antispam.Config
	SMSAntispam          *antispam.Config
	FaceAntispam         *antispam.Config
	VIPAntispam          *antispam.Config
	Host                 *Host
	BFS                  *BFS
	FaceBFS              *BFS
	AccMemcache          *memcache.Config
	AccRedis             *redis.Config
	Realname             *Realname
	Supervisor           *supervisor.Config
	NickFreeAppKeys      map[string]string
	Report               *databus.Config
	Switch               *Switch
	Vipproperty          *VipProperty
	AccountNotify        *databus.Config
	CardClient           *warden.ClientConfig
	Geetest              *Geetest
	Account              *Account
	VipThirdVerifyConfig *vipverify.Config
	VipClient            *warden.ClientConfig
	CouponClient         *warden.ClientConfig
}

// Account is
type Account struct {
	RemoveLoginLogCIDR []string
}

// Geetest is
type Geetest struct {
	PC GeetestConfig
	H5 GeetestConfig
}

// GeetestConfig conf.
type GeetestConfig struct {
	CaptchaID  string
	PrivateKEY string
}

// VipProperty .
type VipProperty struct {
	CodeOpenwhiteIPMap         map[string][]string
	OfficialMid                int64
	OAuthClient                *bm.ClientConfig
	EleOAuthURI                string
	EleConsumerKey             string
	EleOAuthCallBackURI        string
	ActivityURI                string
	ActStartTime               int64
	ActEndTime                 int64
	AssociateWhiteIPMap        map[string][]string
	AssociateWhiteMidMap       []int64
	AssociateWhiteOutOpenIDMap []string
}

// Realname .
type Realname struct {
	DataDir                 string
	ImageExpire             time.Duration
	AlipayAntispamTTL       int32
	AlipayAntispamThreshold int
	Geetest                 *struct {
		RegisterURL string
		ValidateURL string
		CaptchaID   string
		PrivateKey  string
	}
	Alipay *struct {
		Gateway string
		AppID   string
	}
	Channel []*struct {
		Name string
		Flag bool
	}
}

// RsaPub .
func RsaPub() (key string) {
	if client == nil {
		return ""
	}
	key, _ = client.Value("realname.rsa.pub")
	return
}

// RsaPriv .
func RsaPriv() (key string) {
	if client == nil {
		return ""
	}
	key, _ = client.Value("realname.rsa.priv")
	return
}

// AlipayPub .
func AlipayPub() (key string) {
	if client == nil {
		return ""
	}
	key, _ = client.Value("realname.alipay.pub")
	return
}

// AlipayBiliPriv bilibili generate rsa private key
func AlipayBiliPriv() (key string) {
	if client == nil {
		return ""
	}
	key, _ = client.Value("realname.alipay.bili.priv")
	return
}

// BFS bfs config
type BFS struct {
	Timeout     time.Duration
	MaxFileSize int
	URL         string
	Method      string
	Bucket      string
	Key         string
	Secret      string
}

// HTTPServers Http Servers
type HTTPServers struct {
	Inner *bm.ServerConfig
	// Local *bm.ServerConfig
}

// Host host.
type Host struct {
	AccCom      string
	AccCo       string
	Passport    string
	API         string
	Vip         string
	WWW         string
	Search      string
	CM          string
	PassportCom string
}

// RPC config
type RPC struct {
	Relation *rpc.ClientConfig
	Member   *rpc.ClientConfig
	Account  *rpc.ClientConfig
	Vip      *rpc.ClientConfig
	Usersuit *rpc.ClientConfig
	Archive  *rpc.ClientConfig
	UP       *rpc.ClientConfig
	Article  *rpc.ClientConfig
	PassPort *rpc.ClientConfig
	Coin     *rpc.ClientConfig
	Location *rpc.ClientConfig
	Secure   *rpc.ClientConfig
	Filter   *rpc.ClientConfig
	Coupon   *rpc.ClientConfig
	Point    *rpc.ClientConfig
	Resource *rpc.ClientConfig
}

// HTTPClient conf.
type HTTPClient struct {
	Normal *bm.ClientConfig
	Slow   *bm.ClientConfig
}

// Switch is.
type Switch struct {
	UpdatePropertyPhoneRequired bool
}

func configCenter() (err error) {
	if client, err = conf.New(); err != nil {
		panic(err)
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

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf.
func Init() (err error) {
	if confPath == "" {
		return configCenter()
	}
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
