package conf

import (
	"flag"

	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/antispam"
	v "github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf init config
	Conf *Config
)

// Config config.
type Config struct {
	// log
	Log *log.Config
	//rpc server2
	RPCServer *rpc.ServerConfig
	// db
	DB *sql.Config
	// redis
	Redis *redis.Config
	// timeout
	PollTimeout time.Duration
	// local cache
	PathCache string
	// orm
	ORM *orm.Config
	//BM
	BM *bm.ServerConfig
	// Antispam
	Antispam *antispam.Config
	Verify   *v.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "./config-service-example.toml", "config path")
}

// Init init.
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
