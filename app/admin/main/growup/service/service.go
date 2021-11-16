package service

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/growup/conf"
	"github.com/namelessup/bilibili/app/admin/main/growup/dao"
	"github.com/namelessup/bilibili/app/admin/main/growup/dao/message"
	"github.com/namelessup/bilibili/app/admin/main/growup/dao/resource"
	"github.com/namelessup/bilibili/app/admin/main/growup/dao/shell"
	"github.com/namelessup/bilibili/app/admin/main/growup/model/offlineactivity"
	"github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Service struct
type Service struct {
	conf                *conf.Config
	dao                 *dao.Dao
	msg                 *message.Dao
	chanCheckDb         chan int
	chanCheckShellOrder chan *offlineactivity.OfflineActivityResult
	chanCheckActivity   chan int64 // it's result id in this channel
	shellClient         *shell.Client
}

// New fn
func New(c *conf.Config) (s *Service) {
	s = &Service{
		conf:                c,
		dao:                 dao.New(c),
		msg:                 message.New(c),
		chanCheckDb:         make(chan int, 1),
		chanCheckShellOrder: make(chan *offlineactivity.OfflineActivityResult, 10240),
		chanCheckActivity:   make(chan int64, 1000),
		shellClient:         shell.New(c.ShellConf, blademaster.NewClient(c.HTTPClient)),
	}
	resource.Init(c)
	if c.OtherConf.OfflineOrderConsume {
		go s.offlineactivityCheckSendDbProc()
	}
	go s.offlineactivityCheckShellOrderProc()
	return s
}

// Ping fn
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close dao
func (s *Service) Close() {
	s.dao.Close()
}
