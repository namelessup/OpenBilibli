package gorpc

import (
	pb "github.com/namelessup/bilibili/app/service/main/sms/api"
	"github.com/namelessup/bilibili/app/service/main/sms/conf"
	"github.com/namelessup/bilibili/app/service/main/sms/model"
	"github.com/namelessup/bilibili/app/service/main/sms/service"
	"github.com/namelessup/bilibili/library/net/rpc"
	"github.com/namelessup/bilibili/library/net/rpc/context"
	"github.com/namelessup/bilibili/library/net/rpc/interceptor"
)

// RPC rpc server
type RPC struct {
	s *service.Service
}

// New new rpc server.
func New(c *conf.Config, s *service.Service) (svr *rpc.Server) {
	r := &RPC{s: s}
	svr = rpc.NewServer(c.RPCServer)
	in := interceptor.NewInterceptor("")
	svr.Interceptor = in
	if err := svr.Register(r); err != nil {
		panic(err)
	}
	return
}

// Send rpc send.
func (r *RPC) Send(c context.Context, a *model.ArgSend, res *struct{}) (err error) {
	req := &pb.SendReq{Mid: a.Mid, Mobile: a.Mobile, Country: a.Country, Tcode: a.Tcode, Tparam: a.Tparam}
	_, err = r.s.Send(c, req)
	return
}

// SendBatch rpc sendbatch.
func (r *RPC) SendBatch(c context.Context, a *model.ArgSendBatch, res *struct{}) (err error) {
	req := &pb.SendBatchReq{Mids: a.Mids, Mobiles: a.Mobiles, Tcode: a.Tcode, Tparam: a.Tparam}
	_, err = r.s.SendBatch(c, req)
	return
}
