package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/library/conf"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
	"github.com/namelessup/bilibili/library/net/trace"

	"github.com/BurntSushi/toml"
	"github.com/namelessup/bilibili/library/database/orm"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	"github.com/namelessup/bilibili/library/queue/databus"
	"path"
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
	Ecode  *ecode.Config
	// rpc client
	AccountRPC *rpc.ClientConfig
	// gorm
	Upcrm            *orm.Config
	ArchiveOrm       *orm.Config
	MailConf         *Mail
	MailTemplateConf *MailTemplateConfig
	DatabusConf      *DataBusConfig
	GRPCClient       *GRPCClient
	// cron job
	Job    *JobCron
	IsTest bool
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
	if err != nil {
		return
	}
	var templateConfPath = path.Join(path.Dir(confPath), "mail-template.toml")
	_, err = toml.DecodeFile(templateConfPath, &Conf)
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

//Mail 邮件配置
type Mail struct {
	Host               string
	Port               int
	Username, Password string
	DueMailReceivers   []string //  []adminname, send to adminname@bilibili.com
}

//JobCron 任务时间配置
type JobCron struct {
	UpCheckDateDueTaskTime string
	TaskScheduleTime       string
	CheckStateJobTime      string
	UpdateUpTidJobTime     string
}

//MailTemplateConfig mail template conf
type MailTemplateConfig struct {
	SignTmplTitle   string
	SignTmplContent string
	PayTmplTitle    string
	PayTmplContent  string
	TaskTmplTitle   string
	TaskTmplContent string
}

//DataBusConfig databus config
type DataBusConfig struct {
	ArchiveNotify *databus.Config
	Archive       *databus.Config
}

// GRPCClient .
type GRPCClient struct {
	Up      *warden.ClientConfig
	Archive *warden.ClientConfig
}
