package service

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/vip/model"
)

// MailCouponCodeCreate salary mail coupon.
func (s *Service) MailCouponCodeCreate(c context.Context, mid int64) (err error) {
	if err = s.dao.MailCouponCodeCreate(c, &model.ArgMailCouponCodeCreate{Mid: mid, CouponID: s.c.AssociateConf.MailCouponID1}); err != nil {
		return
	}
	return s.dao.MailCouponCodeCreate(c, &model.ArgMailCouponCodeCreate{Mid: mid, CouponID: s.c.AssociateConf.MailCouponID2})
}
