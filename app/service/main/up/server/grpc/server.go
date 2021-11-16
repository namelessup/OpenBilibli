package grpc

import (
	uprpc "github.com/namelessup/bilibili/app/service/main/up/api/v1"
	"github.com/namelessup/bilibili/app/service/main/up/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New new a grpc server.
func New(cfg *warden.ServerConfig, s *service.Service) *warden.Server {
	grpc := warden.NewServer(cfg)
	uprpc.RegisterUpServer(grpc.Server(), s)
	grpc, err := grpc.Start()
	if err != nil {
		panic(err)
	}
	return grpc
}
