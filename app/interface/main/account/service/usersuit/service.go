package usersuit

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/account/conf"
	"github.com/namelessup/bilibili/app/interface/main/account/dao/usersuit"
	"github.com/namelessup/bilibili/app/interface/main/account/dao/vip"
	accrpc "github.com/namelessup/bilibili/app/service/main/account/rpc/client"
	coinrpc "github.com/namelessup/bilibili/app/service/main/coin/api/gorpc"
	memrpc "github.com/namelessup/bilibili/app/service/main/member/api/gorpc"
	usrpc "github.com/namelessup/bilibili/app/service/main/usersuit/rpc/client"
)

// Service struct.
type Service struct {
	c       *conf.Config
	dao     *usersuit.Dao
	vipDao  *vip.Dao
	usRPC   *usrpc.Service2
	accRPC  *accrpc.Service3
	coinRPC *coinrpc.Service
	memRPC  *memrpc.Service
}

// New a pendant service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:       c,
		dao:     usersuit.New(c),
		vipDao:  vip.New(c),
		usRPC:   usrpc.New(c.RPCClient2.Usersuit),
		memRPC:  memrpc.New(c.RPCClient2.Member),
		accRPC:  accrpc.New3(c.RPCClient2.Account),
		coinRPC: coinrpc.New(c.RPCClient2.Coin),
	}
	return
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return
}

// Close dao.
func (s *Service) Close() {}
