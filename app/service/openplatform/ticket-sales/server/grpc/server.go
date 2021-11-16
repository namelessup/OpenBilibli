package grpc

import (
	"context"
	rpc "github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"

	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/service/mis"
	"google.golang.org/grpc"
)

//New 生成rpc服务
func New(srv *service.Service, misSrv *mis.Mis) *warden.Server {
	s := warden.NewServer(nil)
	s.Use(middleware())
	rpc.RegisterTradeServer(s.Server(), srv)
	rpc.RegisterPromotionServer(s.Server(), srv)
	rpc.RegisterPromotionMisServer(s.Server(), misSrv)
	rpc.RegisterTicketServer(s.Server(), srv)
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
