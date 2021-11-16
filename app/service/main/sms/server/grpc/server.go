package grpc

import (
	pb "github.com/namelessup/bilibili/app/service/main/sms/api"
	"github.com/namelessup/bilibili/app/service/main/sms/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New Sms warden rpc server
func New(c *warden.ServerConfig, svr *service.Service) *warden.Server {
	ws := warden.NewServer(c)
	pb.RegisterSmsServer(ws.Server(), svr)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
