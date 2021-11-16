package conf

import (
	"context"
	"fmt"
	"sync"

	"github.com/namelessup/bilibili/app/interface/openplatform/monitor-end/model/kafka"
	"github.com/namelessup/bilibili/app/interface/openplatform/monitor-end/model/monitor"
	"github.com/namelessup/bilibili/library/cache/redis"
	"github.com/namelessup/bilibili/library/conf/paladin"
	"github.com/namelessup/bilibili/library/database/sql"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf is global config
	Conf = &Config{Prom: &Prom{}}
)

// Config service config
type Config struct {
	// http
	BM *bm.ServerConfig
	// XLog
	Log *log.Config
	// monitor
	Monitor *monitor.MonitorConfig
	// kafka
	Kafka       *kafka.Config
	NeedConsume bool
	// Prometheus
	Prom  *Prom
	mutex sync.Mutex
	// mysql
	MySQL *sql.Config
	// redis
	Redis *redis.Config
	// infoc
	CollectInfoc *infoc.Config
	// collect fronend log
	CollectFE bool
	// products
	Products string
}

// Prom .
type Prom struct {
	Promed   bool
	IgnoreNA bool
	Factor   int
	Limit    uint64
}

// Set .
func (e *Config) Set(text string) error {
	var ec Config
	if err := toml.Unmarshal([]byte(text), &ec); err != nil {
		return err
	}
	if ec.Prom == nil {
		e.Prom = &Prom{}
		return nil
	}
	e.mutex.Lock()
	e.Prom = &Prom{
		Factor:   ec.Prom.Factor,
		Promed:   ec.Prom.Promed,
		Limit:    ec.Prom.Limit,
		IgnoreNA: ec.Prom.IgnoreNA,
	}
	e.CollectFE = ec.CollectFE
	e.Products = ec.Products
	e.mutex.Unlock()
	// *e = ec
	return nil
}

// Init .
func Init() {
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	// var ec exampleConf
	// var setter
	if err := paladin.Watch("monitor-end.toml", Conf); err != nil {
		panic(err)
	}
	if err := paladin.Get("monitor-end.toml").UnmarshalTOML(Conf); err != nil {
		panic(err)
	}

	// use exampleConf
	// watch key
	go func() {
		for event := range paladin.WatchEvent(context.TODO(), "monitor-end.toml") {
			fmt.Println(event, Conf)
		}
	}()
}

// func init() {
// 	flag.StringVar(&confPath, "conf", "", "config path")
// }

// func local() (err error) {
// 	_, err = toml.DecodeFile(confPath, &Conf)
// 	return
// }

// Init int config
// func Init() error {
// 	if confPath != "" {
// 		return local()
// 	}
// 	return remote()
// }

// func remote() (err error) {
// 	if client, err = conf.New(); err != nil {
// 		return
// 	}
// 	if err = load(); err != nil {
// 		return
// 	}
// 	go func() {
// 		for range client.Event() {
// 			log.Info("config reload")
// 			if load() != nil {
// 				log.Error("config reload error (%v)", err)
// 			}
// 		}
// 	}()
// 	return
// }

// func load() (err error) {
// 	var (
// 		s       string
// 		ok      bool
// 		tmpConf *Config
// 	)
// 	if s, ok = client.Toml2(); !ok {
// 		return errors.New("load config center error")
// 	}
// 	if _, err = toml.Decode(s, &tmpConf); err != nil {
// 		return errors.New("could not decode config")
// 	}
// 	*Conf = *tmpConf
// 	return
// }
