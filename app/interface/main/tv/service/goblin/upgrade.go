package goblin

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/tv/model"
	"github.com/namelessup/bilibili/library/ecode"
)

// VerUpdate .
func (s *Service) VerUpdate(c context.Context, ver *model.VerUpdate) (result *model.HTTPData, errCode ecode.Codes, err error) {
	result, errCode, err = s.dao.VerUpdate(c, ver)
	return
}
