package grpc

import (
	v1pb "github.com/namelessup/bilibili/app/service/live/xuserex/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/xuserex/service"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// New new grpc server
func New(svc *service.Service) (wsvr *warden.Server, err error) {
	wsvr = warden.NewServer(nil)

	v1pb.RegisterRoomNoticeServer(wsvr.Server(), svc.RoomNoticeV1Svc())

	if wsvr, err = wsvr.Start(); err != nil {
		return
	}
	return
}
