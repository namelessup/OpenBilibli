package grpc

import (
	v1pb "github.com/namelessup/bilibili/app/service/live/xanchor/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/xanchor/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New new grpc server
func New(svc *service.Service) (wsvr *warden.Server, err error) {
	wsvr = warden.NewServer(nil)
	v1pb.RegisterXAnchorServer(wsvr.Server(), svc.V1Svc())
	if wsvr, err = wsvr.Start(); err != nil {
		return
	}
	return
}
