package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/bfs"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/net/trace"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	client   *conf.Client
	// Conf export config var
	Conf = &Config{}
)

// Config danmaku config
type Config struct {
	// base
	// log
	Xlog   *log.Config
	Infoc2 *infoc.Config
	// tracer
	Tracer *trace.Config
	// http
	HTTPServer *bm.ServerConfig
	// database
	DB *DB
	// redis
	Redis *Redis
	// memcache
	Memcache *Memcache
	// Subtitle Cache
	SubtitleMemcache *Memcache
	DMMemcache       *Memcache
	// archive rpc client
	ArchiveRPC *rpc.ClientConfig
	// seq-server rpc client
	SeqRPC *rpc.ClientConfig
	Seq    *Seq
	// databus config
	Databus *Databus
	// dm list realname
	Realname   map[string]int64
	HTTPClient *bm.ClientConfig
	Host       *Host
	BFSClient  *bm.ClientConfig
	// client
	FliterRPC *warden.ClientConfig
	// MaskCate
	MaskCate *MaskCate
	// Bfs
	Bfs *Bfs
	// cache routine size
	RoutineSize int
	// bnj
	BNJ *BNJ
	// task config
	TaskConf *TaskConf
}

// BNJ .
type BNJ struct {
	Aid          int64
	BnjCounter   *BnjCounter
	BnjLiveDanmu *BnjLiveDanmu
}

// BnjCounter .
type BnjCounter struct {
	SubAids []int64
}

// BnjLiveDanmu .
type BnjLiveDanmu struct {
	RoomID      int64
	Start       string
	IgnoreRate  int64
	Level       int32
	IgnoreBegin time.Duration
	IgnoreEnd   time.Duration
}

// BNJVideo .
type BNJVideo struct {
	Cid      int64
	Duration float64
}

// Bfs .
type Bfs struct {
	Client *bfs.Config
	Dm     string
}

//Seq Conf
type Seq struct {
	BusinessID int64
	Token      string
}

// DB mysql config struct
type DB struct {
	DMReader      *sql.Config
	DMWriter      *sql.Config
	BiliDMWriter  *sql.Config
	QueryPageSize int32
}

// Redis dm redis
type Redis struct {
	DM *struct {
		*redis.Config
		Expire time.Duration
	}
	DMRct *struct {
		*redis.Config
		Expire time.Duration
	}
	DMSeg *struct {
		*redis.Config
		Expire time.Duration
	}
}

// Memcache dm memcache
type Memcache struct {
	*memcache.Config
	Expire time.Duration
}

// Databus databus config
type Databus struct {
	IndexCsmr         *databus.Config
	SubjectCsmr       *databus.Config
	ActionCsmr        *databus.Config
	ReportCsmr        *databus.Config
	VideoupCsmr       *databus.Config
	SubtitleAuditCsmr *databus.Config
	BnjCsmr           *databus.Config
}

// Host hosts used in dm admin
type Host struct {
	Videoup   string
	Mask      string
	DataRank  string
	MerakHost string
	APILive   string
}

// MaskCate .
type MaskCate struct {
	Tids     []int64
	Interval time.Duration
	Limit    int
}

// TaskConf .
type TaskConf struct {
	DelInterval  time.Duration
	ResInterval  time.Duration
	ResFieldLen  int
	DelNum       int
	DelLimit     int64
	MsgCC        []string
	MsgPublicKey string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

//Init int config
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
	var tmpConf *Config
	value, ok := client.Toml2()
	if !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(value, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}
