package grpc

import (
	pb "github.com/namelessup/bilibili/app/service/main/reply-feed/api"
	"github.com/namelessup/bilibili/app/service/main/reply-feed/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New grpc server
func New(cfg *warden.ServerConfig, srv *service.Service) (wsvr *warden.Server) {
	var err error
	wsvr = warden.NewServer(cfg)
	pb.RegisterReplyFeedServer(wsvr.Server(), srv)
	wsvr, err = wsvr.Start()
	if err != nil {
		panic(err)
	}
	return
}
