package dao

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/tv/model"
	"github.com/namelessup/bilibili/library/log"
)

const (
	_userTableName = "tv_user_info"
)

// GetByMId  select user info by mid
func (d *Dao) GetByMId(c context.Context, mid int64) (userInfo *model.TvUserInfoResp, err error) {
	userInfo = &model.TvUserInfoResp{}

	if err = d.DB.Table(_userTableName).Where("mid = ?", mid).First(userInfo).Error; err != nil {
		log.Error("GetByMId (%v) error(%v)", mid, err)
	}

	return
}
