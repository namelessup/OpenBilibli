package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	xlog "github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// Config .
type Config struct {
	Env         string
	Log         *xlog.Config
	Tracer      *trace.Config
	PushRPC     *warden.ClientConfig
	HTTPServer  *bm.ServerConfig
	HTTPClient  *bm.ClientConfig
	DpClient    *bm.ClientConfig
	ReportSub   *databus.Config
	CallbackSub *databus.Config
	MySQL       *sql.Config
	Memcache    *mc
	Wechat      *wechat
	Job         *job
}

// mc config
type mc struct {
	*memcache.Config
}

type wechat struct {
	Token    string
	Secret   string
	Username string
}

type job struct {
	ReportTicker             xtime.Duration
	DelInvalidReportInterval xtime.Duration
	LoadTaskInteval          xtime.Duration
	PullResultIntervalHour   int
	DelCallbackInterval      int
	DelTaskInterval          int
	SyncReportCacheWeek      int
	SyncReportCacheHour      int
	ReportShard              int
	CallbackShard            int
	PretreatmentTaskShard    int
	TaskGoroutines           int
	LimitPerTask             int
	PushPartSize             int
	PushPartChanSize         int
	MountDir                 string
	PretreatTask             bool
	DpPollingTime            xtime.Duration
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
	err = load()
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
