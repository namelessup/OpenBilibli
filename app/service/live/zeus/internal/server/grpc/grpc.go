package grpc

import (
	v1pb "github.com/namelessup/bilibili/app/service/live/zeus/api/v1"
	"github.com/namelessup/bilibili/app/service/live/zeus/internal/conf"
	v1srv "github.com/namelessup/bilibili/app/service/live/zeus/internal/service/v1"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

func New(c *conf.Config, s *v1srv.ZeusService) *warden.Server {
	ws := warden.NewServer(nil)
	v1pb.RegisterZeusServer(ws.Server(), s)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
