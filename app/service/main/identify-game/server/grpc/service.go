package grpc

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/identify-game/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/main/identify-game/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New identify game warden rpc server
func New(cfg *warden.ServerConfig, s *service.Service) *warden.Server {
	w := warden.NewServer(cfg)
	v1.RegisterIdentifyGameServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

var _ v1.IdentifyGameServer = &server{}

func (s *server) DelCache(ctx context.Context, req *v1.DelCacheReq) (*v1.DelCacheReply, error) {
	err := s.svr.DelCache(ctx, req.Token)
	return &v1.DelCacheReply{}, err
}

func (s *server) GetCookieByToken(ctx context.Context, req *v1.CreateCookieReq) (*v1.CreateCookieReply, error) {
	cookies, err := s.svr.GetCookieByToken(ctx, req.Token, req.From)
	return cookies, err
}
