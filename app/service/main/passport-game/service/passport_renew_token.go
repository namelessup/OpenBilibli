package service

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/service/main/passport-game/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// RenewToken renew token.
func (s *Service) RenewToken(c context.Context, accessKey, from string) (res *model.RenewToken, err error) {
	if accessKey == "" {
		err = ecode.AccessKeyErr
		return
	}
	region, ok := region(accessKey)
	if !ok {
		err = ecode.AccessKeyErr
		return
	}
	if region == s.currentRegion {
		return s.currentRenewToken(c, accessKey)
	}
	if region == _origin {
		if res, err = s.d.RenewToken(context.TODO(), s.renewToken[_origin], accessKey, s.currentRegion); err != nil {
			return
		}
		if _, innerErr := s.currentRenewToken(c, accessKey); innerErr != nil {
			log.Error("renewtoken succeeded on origin but failed on cloud, accessKey(%s) from(%s) error(%v)", accessKey, from, innerErr)
		}
		return
	}
	if from != "" {
		s.dispatcherErrStats.Incr("dispatcher_error")
		err = ErrDispatcherError
		return
	}
	return s.d.RenewToken(c, s.renewToken[region], accessKey, s.currentRegion)
}

func (s *Service) currentRenewToken(c context.Context, accessKey string) (res *model.RenewToken, err error) {
	var tokenInfo *model.Perm
	if tokenInfo, err = s.tokenInfo(c, accessKey); err != nil {
		return
	}
	if tokenInfo == nil {
		err = ecode.AccessKeyErr
		return
	}
	expires := time.Now().Unix() + _expireSeconds
	token := &model.Perm{
		Expires:     expires,
		AccessToken: accessKey,
	}
	if _, err = s.d.UpdateToken(c, token); err != nil {
		return
	}
	s.d.DelTokenCache(c, accessKey)
	res = &model.RenewToken{
		Expires: expires,
	}
	return
}
