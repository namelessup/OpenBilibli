package service

import (
	"context"

	rpcModel "github.com/namelessup/bilibili/app/service/main/tag/model"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
)

// AddCustomSubChannel AddCustomSubChannel.
func (s *Service) addCustomSubChannels(c context.Context, mid int64, tp int, tids []int64) (err error) {
	arg := &rpcModel.ArgCustomSub{
		Mid:    mid,
		Type:   tp,
		Tids:   tids,
		RealIP: metadata.String(c, metadata.RemoteIP),
	}
	if err = s.tagRPC.AddCustomSubChannel(c, arg); err != nil {
		log.Error("s.tagRPC.AddCustomSubTag()ArgID:%+v, error(%v)", arg, err)
	}
	return
}
