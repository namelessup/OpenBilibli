package grpc

import (
	pb "github.com/namelessup/bilibili/app/service/main/history/api/grpc"
	"github.com/namelessup/bilibili/app/service/main/history/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New Coin warden rpc server
func New(c *warden.ServerConfig, svr *service.Service) *warden.Server {
	ws := warden.NewServer(c)
	pb.RegisterHistoryServer(ws.Server(), svr)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
