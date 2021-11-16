package block

import (
	"context"
	"runtime/debug"

	"github.com/namelessup/bilibili/app/admin/main/member/conf"
	"github.com/namelessup/bilibili/app/admin/main/member/dao/block"
	account "github.com/namelessup/bilibili/app/service/main/account/api"
	rpcfigure "github.com/namelessup/bilibili/app/service/main/figure/rpc/client"
	rpcspy "github.com/namelessup/bilibili/app/service/main/spy/rpc/client"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/sync/pipeline/fanout"
)

// Service struct
type Service struct {
	conf             *conf.Config
	dao              *block.Dao
	cache            *fanout.Fanout
	spyRPC           *rpcspy.Service
	figureRPC        *rpcfigure.Service
	accountClient    account.AccountClient
	missch           chan func()
	accountNotifyPub *databus.Databus
}

// New init
func New(conf *conf.Config, dao *block.Dao, spyRPC *rpcspy.Service, figureRPC *rpcfigure.Service,
	accountClient account.AccountClient, accountNotifyPub *databus.Databus) (s *Service) {
	s = &Service{
		conf:             conf,
		dao:              dao,
		cache:            fanout.New("memberAdminCache", fanout.Worker(1), fanout.Buffer(10240)),
		missch:           make(chan func(), 10240),
		accountNotifyPub: accountNotifyPub,
		spyRPC:           spyRPC,
		figureRPC:        figureRPC,
		accountClient:    accountClient,
	}
	go s.missproc()
	return s
}

func (s *Service) missproc() {
	defer func() {
		if x := recover(); x != nil {
			log.Error("service.missproc panic(%+v) : %s", x, debug.Stack())
			go s.missproc()
		}
	}()
	for {
		f := <-s.missch
		f()
	}
}

func (s *Service) mission(f func()) {
	select {
	case s.missch <- f:
	default:
		log.Error("s.missch full")
	}
}

// Ping Service
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close Service
func (s *Service) Close() {
	s.dao.Close()
}
