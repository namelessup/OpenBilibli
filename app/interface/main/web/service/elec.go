package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/web/model"
	arcmdl "github.com/namelessup/bilibili/app/service/main/archive/api"
	"github.com/namelessup/bilibili/app/service/main/archive/model/archive"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// ElecShow elec show.
func (s *Service) ElecShow(c context.Context, mid, aid, loginID int64) (rs *model.ElecShow, err error) {
	var arcReply *arcmdl.ArcReply
	if arcReply, err = s.arcClient.Arc(c, &arcmdl.ArcRequest{Aid: aid}); err != nil {
		log.Error("s.arcClient.Arc(%d) error(%v)", aid, err)
		return
	}
	arc := arcReply.Arc
	if arc.Copyright != int32(archive.CopyrightOriginal) || !arc.IsNormal() {
		err = ecode.ElecDenied
		return
	}
	if _, ok := s.elecShowTypeIds[arc.TypeID]; !ok {
		err = ecode.ElecDenied
		return
	}
	if rs, err = s.dao.ElecShow(c, mid, aid, loginID); err != nil {
		log.Error("s.dao.ElecShow(%d,%d,%d) error(%v)", mid, aid, loginID, err)
	}
	return
}
