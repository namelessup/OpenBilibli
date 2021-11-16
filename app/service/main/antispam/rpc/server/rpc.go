package server

import (
	"log"

	"github.com/namelessup/bilibili/app/service/main/antispam/conf"
	"github.com/namelessup/bilibili/app/service/main/antispam/service"

	"github.com/namelessup/bilibili/library/net/rpc"
)

// New .
func New(config *conf.Config, s service.Service) *rpc.Server {
	rpcSvr := rpc.NewServer(config.RPC)
	if err := rpcSvr.Register(&Filter{svr: s}); err != nil {
		log.Fatalf("%+v", err)
	}
	return rpcSvr
}
