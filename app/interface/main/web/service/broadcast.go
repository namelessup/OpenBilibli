package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/web/model"
	warden "github.com/namelessup/bilibili/app/service/main/broadcast/api/grpc/v1"
	"github.com/namelessup/bilibili/library/log"
)

// BroadServers broadcast server list.
func (s *Service) BroadServers(c context.Context, platform string) (res *warden.ServerListReply, err error) {
	if res, err = s.broadcastClient.ServerList(c, &warden.ServerListReq{Platform: platform}); err != nil {
		log.Error("s.broadCastClient.ServerList(%s) error(%v)", platform, err)
		res = model.DefaultServer
		err = nil
	}
	return
}
