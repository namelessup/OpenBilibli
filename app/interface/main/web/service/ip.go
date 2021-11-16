package service

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/location/model"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
)

// IPZone get ip zone info by ip
func (s *Service) IPZone(c context.Context) (res *model.Info, err error) {
	ip := metadata.String(c, metadata.RemoteIP)
	if res, err = s.loc.Info(c, &model.ArgIP{IP: ip}); err != nil {
		log.Error("s.loc.Info(%s) error(%v)", ip, err)
	}
	return
}
