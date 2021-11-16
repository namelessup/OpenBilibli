package grpc

import (
	pb "github.com/namelessup/bilibili/app/service/bbq/notice-service/api/v1"
	"github.com/namelessup/bilibili/app/service/bbq/notice-service/internal/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New new warden rpc server
func New(c *warden.ServerConfig, svc *service.Service) *warden.Server {
	ws := warden.NewServer(c)
	pb.RegisterNoticeServer(ws.Server(), svc)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
