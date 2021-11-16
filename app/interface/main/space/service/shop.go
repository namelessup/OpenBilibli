package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/space/model"
)

// ShopInfo get shop info.
func (s *Service) ShopInfo(c context.Context, mid int64) (data *model.ShopInfo, err error) {
	return s.dao.ShopInfo(c, mid)
}
