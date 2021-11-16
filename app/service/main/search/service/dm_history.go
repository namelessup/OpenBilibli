package service

import (
	"context"
	"fmt"

	"github.com/namelessup/bilibili/app/service/main/search/dao"
	"github.com/namelessup/bilibili/app/service/main/search/model"
	"github.com/namelessup/bilibili/library/ecode"
)

// DMHistory .
func (s *Service) DmHistory(c context.Context, sp *model.DmHistoryParams) (res *model.SearchResult, err error) {
	if res, err = s.dao.DmHistory(c, sp); err != nil {
		dao.PromError(fmt.Sprintf("es:%s 搜索dm_history失败", sp.Bsp.AppID), "s.dao.DmHistory(%v) error(%v)", sp, err)
		err = ecode.SearchDmFailed
	}
	return
}
