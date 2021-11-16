// Package server generate by warden_gen
package server

import (
	pb "github.com/namelessup/bilibili/app/service/main/coin/api"
	"github.com/namelessup/bilibili/app/service/main/coin/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New Coin warden rpc server
func New(c *warden.ServerConfig, svr *service.Service) *warden.Server {
	ws := warden.NewServer(c)
	pb.RegisterCoinServer(ws.Server(), svr)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
