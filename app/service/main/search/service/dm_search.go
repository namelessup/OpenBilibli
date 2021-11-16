package service

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/service/main/search/dao"
	"github.com/namelessup/bilibili/app/service/main/search/model"
	"github.com/namelessup/bilibili/library/ecode"
)

func (s *Service) DmSearch(c context.Context, sp *model.DmSearchParams) (res *model.SearchResult, err error) {
	if res, err = s.dao.DmSearch(c, sp); err != nil {
		dao.PromError(fmt.Sprintf("es:%s 搜索dm_search失败", sp.Bsp.AppID), "s.dao.DmSearch(%v) error(%v)", sp, err)
		err = ecode.SearchDmFailed
	}
	return
}
