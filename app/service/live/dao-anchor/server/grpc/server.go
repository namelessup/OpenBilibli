package grpc

import (
	"time"

	v0pb "github.com/namelessup/bilibili/app/service/live/dao-anchor/api/grpc/v0"
	v1pb "github.com/namelessup/bilibili/app/service/live/dao-anchor/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/dao-anchor/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
	xtime "github.com/namelessup/bilibili/library/time"

	"google.golang.org/grpc"
)

// New new grpc server
func New(svc *service.Service) (wsvr *warden.Server, err error) {
	wsvr = warden.NewServer(&warden.ServerConfig{Timeout: xtime.Duration(time.Second * 10)}, grpc.MaxRecvMsgSize(1024*1024*1024), grpc.MaxSendMsgSize(1024*1024*1024))
	v0pb.RegisterCreateDataServer(wsvr.Server(), svc.CreateDataSvc())
	//	v0pb.RegisterPopularityServer(wsvr.Server(), svc.PopularitySvc())
	v1pb.RegisterDaoAnchorServer(wsvr.Server(), svc.V1Svc())
	if wsvr, err = wsvr.Start(); err != nil {
		return
	}
	return
}
