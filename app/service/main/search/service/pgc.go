package service

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/service/main/search/dao"
	"github.com/namelessup/bilibili/app/service/main/search/model"
	"github.com/namelessup/bilibili/library/ecode"
)

// PgcMedia .
func (s *Service) PgcMedia(c context.Context, sp *model.PgcMediaParams) (res *model.SearchResult, err error) {
	if res, err = s.dao.PgcMedia(c, sp); err != nil {
		dao.PromError(fmt.Sprintf("es:%s 搜索pgc番剧失败", sp.Bsp.AppID), "s.dao.PgcMedia(%v) error(%v)", sp, err)
		err = ecode.SearchPgcMediaFailed
	}
	return
}
