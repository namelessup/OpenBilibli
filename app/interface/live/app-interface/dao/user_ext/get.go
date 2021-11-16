package user_ext

import (
	"context"
	"github.com/pkg/errors"
	"github.com/namelessup/bilibili/app/interface/live/app-interface/dao"
	userExV1 "github.com/namelessup/bilibili/app/service/live/userext/api/liverpc/v1"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// GetGrayRule 获取灰度配置
func (d *Dao) GetGrayRule(ctx context.Context, req *userExV1.GrayRuleGetByMarkReq) (extResult *userExV1.GrayRuleGetByMarkResp, err error) {
	extResult = &userExV1.GrayRuleGetByMarkResp{}
	if req == nil {
		return nil, nil
	}
	ret, err := dao.UserExtApi.V1GrayRule.GetByMark(ctx, req)
	if err != nil {
		log.Error("call_userExt_grayRule error,err:%v", err)
		err = errors.WithMessage(ecode.GetGrayRuleError, "GET SEA PATROL FAIL")
		return
	}
	extResult = ret
	return
}
