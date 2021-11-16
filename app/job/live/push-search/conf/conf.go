package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/sql"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/database/hbase.v2"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
	xtime "github.com/namelessup/bilibili/library/time"
)

var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// Config .
type Config struct {
	Log      *log.Config
	BM       *bm.ServerConfig
	Verify   *verify.Config
	Tracer   *trace.Config
	Redis    *redis.Config
	Memcache *memcache.Config
	MySQL    *sql.Config
	Ecode    *ecode.Config
	LiveRpc  map[string]*liverpc.ClientConfig
	//DataBus
	DataBus     *DataBus
	Group       *Group
	SearchHBase *hbaseConf
	MigrateNum	int
}

type DataBus struct {
	RoomInfo *databus.Config
	Attention *databus.Config
	UserName *databus.Config
	PushSearch *databus.Config
}

// Group group.
type Group struct {
	RoomInfo *GroupConf
	Attention *GroupConf
	UserInfo *GroupConf
}

// GroupConf group conf.
type GroupConf struct {
	Num    int
	Chan   int
}

type hbaseConf struct {
	hbase.Config
	ReadTimeout   xtime.Duration
	ReadsTimeout  xtime.Duration
	WriteTimeout  xtime.Duration
	WritesTimeout xtime.Duration
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
