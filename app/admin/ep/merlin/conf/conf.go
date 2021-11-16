package conf

import (
	"errors"
	"flag"

	"github.com/namelessup/bilibili/app/admin/ep/merlin/model"
	"github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/conf"
	"github.com/namelessup/bilibili/library/database/orm"
	ecode "github.com/namelessup/bilibili/library/ecode/tip"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"
	"github.com/namelessup/bilibili/library/time"

	"github.com/BurntSushi/toml"
)

// global var
var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// Memcache memcache.
type Memcache struct {
	*memcache.Config
	Expire time.Duration
}

// Config config set
type Config struct {
	Version string `toml:"version"`
	// base
	// elk
	Log *log.Config
	// http
	BM *bm.ServerConfig
	// ecode
	Ecode *ecode.Config

	HTTPClient *bm.ClientConfig

	Memcache *Memcache

	// orm
	ORM *orm.Config

	Paas *model.PaasConf

	ServiceTree *model.TreeConf

	Mail *Mail

	Scheduler *Scheduler

	Auth *permit.Config

	BiliHub *BiliHub

	DeviceFarm *DeviceFarm

	WeChat *WeChat
}

//WeChat WeChat config
type WeChat struct {
	WeChatHost       string
	WeChatDeviceFarm *WeChatDeviceFarm
}

// WeChatDeviceFarm WeChatDeviceFarm.
type WeChatDeviceFarm struct {
	ChatID      string
	MsgType     string
	Safe        int
	SendMessage bool
}

//DeviceFarm DeviceFarm
type DeviceFarm struct {
	HostList       []string
	SuperOwner     []string
	ApplyMonthTime int
}

// Scheduler scheduler
type Scheduler struct {
	GetExpiredMachinesTime              string
	SendTaskMailMachinesWillExpiredTime string
	DeleteExpiredMachinesInTask         string
	CheckMachinesStatusInTask           string
	UpdateMobileDeviceInTask            string
	UpdateSnapshotStatusInDoing         string
	Active                              bool
	ExpiredDate                         int
}

// Mail mail
type Mail struct {
	Host        string
	Port        int
	Username    string
	Password    string
	NoticeOwner []string
}

// BiliHub BiliHub.
type BiliHub struct {
	Host           string
	HostName       string
	Username       string
	Password       string
	MerlinPub      string
	SharePub       string
	MachineTagPri  string
	SupportNetWork []string
	SuperOwner     []string
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
	return load()
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
