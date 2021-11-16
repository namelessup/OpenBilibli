package conf

import (
	"flag"

	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"

	"github.com/BurntSushi/toml"
)

const (
	configKey = "report-click.toml"
)

// global conf
var (
	confPath string
	Conf     *Config
)

// Infoc2 Infoc2.
type Infoc2 struct {
	RealTime   *infoc.Config
	Statistics *infoc.Config
}

// Config config .
type Config struct {
	Infoc2  *Infoc2
	Env     string
	Tracer  *trace.Config
	Xlog    *log.Config
	App     *bm.App
	BM      *bm.ServerConfig
	Auth    *auth.Config
	Verify  *verify.Config
	AccRPC  *rpc.ClientConfig
	HisRPC  *rpc.ClientConfig
	DataBus *Databus
	Click   *Click
}

// Databus .
type Databus struct {
	Merge *databus.Config
}

// Click click config.
type Click struct {
	WebSecret string
	OutSecret string
	// aes
	AesKey  string
	AesIv   string
	AesSalt string
	// aes2
	AesKey2        string
	AesIv2         string
	AesSalt2       string
	From           []int64
	FromInline     []int64
	InlineDuration int64 // inline play duration line, 10s
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf
func Init() (err error) {
	if confPath == "" {
		return configCenter()
	}
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func configCenter() (err error) {
	var (
		ok     bool
		value  string
		client *conf.Client
	)
	if client, err = conf.New(); err != nil {
		return
	}
	if value, ok = client.Value(configKey); !ok {
		panic(err)
	}
	_, err = toml.Decode(value, &Conf)
	return
}
