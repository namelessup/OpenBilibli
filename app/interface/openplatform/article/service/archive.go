package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/openplatform/article/dao"
	"github.com/namelessup/bilibili/app/service/main/archive/api"
	"github.com/namelessup/bilibili/app/service/main/archive/model/archive"
	"github.com/namelessup/bilibili/library/log"
)

// Archives gets archives by aids.
func (s *Service) Archives(c context.Context, aids []int64, ip string) (arcs map[int64]*api.Arc, err error) {
	arg := &archive.ArgAids2{
		Aids:   aids,
		RealIP: ip,
	}
	if arcs, err = s.arcRPC.Archives3(c, arg); err != nil {
		dao.PromError("rpc:获取视频稿件信息")
		log.Error("s.arcRPC.Archives(%v) error(%+v)", aids, err)
		return
	}
	fmtArcs(arcs)
	return
}

func fmtArcs(arcs map[int64]*api.Arc) {
	for id, v := range arcs {
		if !v.IsNormal() {
			delete(arcs, id)
			continue
		}
		// 会员可见 不展示播放数
		if v.Access >= 10000 {
			v.Stat.View = -1
		}
	}
}
