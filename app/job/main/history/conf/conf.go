package conf

import (
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/tidb"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/sync/pipeline"
	xtime "github.com/namelessup/bilibili/library/time"

	"github.com/namelessup/bilibili/library/database/hbase.v2"

	"github.com/BurntSushi/toml"
)

const (
	configKey = "history-job.toml"
)

// global conf
var (
	confPath string
	Conf     = &Config{}
)

// Config  service conf
type Config struct {
	App           *bm.App
	Log           *log.Config
	Tracer        *trace.Config
	Ecode         *ecode.Config
	Job           *Job
	Info          *HBaseConfig
	HisSub        *databus.Config
	ServiceHisSub *databus.Config
	Sub           *databus.Config
	BM            *bm.ServerConfig
	Redis         *redis.Config
	Merge         *pipeline.Config
	TiDB          *tidb.Config
	LongTiDB      *tidb.Config
}

// HBaseConfig .
type HBaseConfig struct {
	*hbase.Config
	WriteTimeout xtime.Duration
	ReadTimeout  xtime.Duration
}

// Job job.
type Job struct {
	URL             string
	Client          *bm.ClientConfig
	Expire          xtime.Duration
	Max             int
	Batch           int
	ServiceBatch    int
	DeleteLimit     int
	DeleteStartHour int
	DeleteEndHour   int
	DeleteStep      xtime.Duration
	// 用户最近播放列表长度
	CacheLen  int
	QPSLimit  int
	IgnoreMsg bool
	RetryTime xtime.Duration
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
