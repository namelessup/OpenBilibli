package grpc

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/ep/saga/api/grpc/v1"
	"github.com/namelessup/bilibili/app/admin/ep/saga/service/wechat"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New grpc server
func New(cfg *warden.ServerConfig, chat *wechat.Wechat) (*warden.Server, error) {
	svr := warden.NewServer(cfg)
	v1.RegisterSagaAdminServer(svr.Server(), &server{chat: chat})
	return svr.Start()
}

var _ v1.SagaAdminServer = &server{}

type server struct {
	chat *wechat.Wechat
}

func (s *server) PushMsg(ctx context.Context, req *v1.PushMsgReq) (*v1.PushMsgReply, error) {
	err := s.chat.PushMsg(ctx, req.Username, req.Content)
	return &v1.PushMsgReply{}, err
}
