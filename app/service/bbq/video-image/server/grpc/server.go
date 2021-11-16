package grpc

import (
	"context"

	"github.com/namelessup/bilibili/app/service/bbq/video-image/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/bbq/video-image/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"

	"google.golang.org/grpc"
)

//New 生成rpc服务
func New(conf *warden.ServerConfig, srv *service.Service) *warden.Server {
	s := warden.NewServer(conf)
	s.Use(middleware())
	v1.RegisterVideoImageServer(s.Server(), srv)
	_, err := s.Start()
	if err != nil {
		panic("run server failed!" + err.Error())
	}
	return s
}

func middleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//call chain
		resp, err = handler(ctx, req)
		return
	}
}
