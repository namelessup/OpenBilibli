package server

import (
	"github.com/namelessup/bilibili/app/service/main/upcredit/conf"
	"github.com/namelessup/bilibili/app/service/main/upcredit/service"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/context"
)

//RPC rpc server
type RPC struct {
	s *service.Service
}

//New create
func New(c *conf.Config, s *service.Service) (svr *rpc.Server) {
	r := &RPC{s: s}
	svr = rpc.NewServer(c.RPCServer)
	if err := svr.Register(r); err != nil {
		panic(err)
	}
	return
}

// Ping check connection success.
func (r *RPC) Ping(c context.Context, arg *struct{}, res *struct{}) (err error) {
	return
}
