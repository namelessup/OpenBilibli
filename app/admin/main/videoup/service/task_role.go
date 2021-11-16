package service

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/videoup/model/manager"
	"github.com/namelessup/bilibili/library/log"
)

func (s *Service) isLeader(c context.Context, uid int64) bool {
	role, e := s.mng.GetUserRole(c, uid)
	if e != nil {
		log.Error("s.mng.GetUserRole(%d) error(%v)", uid, e)
		return false
	}
	if role == manager.TaskLeader {
		return true
	}
	return false
}
