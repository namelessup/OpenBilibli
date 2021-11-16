package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
	"github.com/namelessup/bilibili/library/database/hbase.v2"
)

var (
	confPath string
	// Conf .
	Conf   = &Config{}
	client *conf.Client
)

// Config config struct .
type Config struct {
	// API host
	Host *Host
	// channal len
	ChanSize    int
	BeginOffset int64
	// log
	Xlog *log.Config
	// http
	BM *bm.ServerConfig
	// tracer
	Tracer *trace.Config
	// tick load pgc
	Tick time.Duration
	// db
	DB *DB
	// redis
	Redis *Redis
	// hbase
	Hbase *hbaseConf
	// http client test
	HTTPClient HTTPClient
	// databus
	ArchiveSub       *databus.Config
	ArchiveResultSub *databus.Config
	VideoupSub       *databus.Config
	ManagerDBSub     *databus.Config
	// ChanSize aid%ChanSize
	ArchiveRPCGroup2 *rpc.ClientConfig
	TagDisConf       *rpc.ClientConfig
	//grpc
	GRPC *GRPC
	// mail
	Mail *mail
}

//GRPC .
type GRPC struct {
	AccRPC *warden.ClientConfig
	UpsRPC *warden.ClientConfig
}

type hbaseConf struct {
	hbase.Config
	ReadTimeout   time.Duration
	ReadsTimeout  time.Duration
	WriteTimeout  time.Duration
	WritesTimeout time.Duration
}

// Host for httpclient
type Host struct {
	Data    string
	API     string
	Archive string
	Profit  string
	WWW     string
}

// DB db struct
type DB struct {
	Archive *sql.Config
	Manager *sql.Config
}

// Redis redis struct
type Redis struct {
	Track *struct {
		*redis.Config
		Expire time.Duration
	}
	Mail      *redis.Config
	Secondary *struct {
		*redis.Config
		Expire time.Duration
	}
}

// HTTPClient http client struct
type HTTPClient struct {
	Read  *bm.ClientConfig
	Write *bm.ClientConfig
}

//mail 邮件配置
type mail struct {
	Host                                     string
	Port, SpeedThreshold, OverspeedThreshold int
	Username, Password                       string
	Addr, PrivateAddr                        []*MailElemenet
}

//MailElemenet 邮件接收人配置
type MailElemenet struct {
	Type string
	Desc string
	Addr []string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init conf
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
