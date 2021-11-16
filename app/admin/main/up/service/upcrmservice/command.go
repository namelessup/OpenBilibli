package upcrmservice

import (
	"context"
	"github.com/namelessup/bilibili/app/admin/main/up/model"
	"time"
)

//CommandRefreshUpRank refresh up rank
func (s *Service) CommandRefreshUpRank(con context.Context, arg *model.CommandCommonArg) (result model.CommandCommonResult, err error) {
	s.RefreshCache(time.Now())
	return
}
