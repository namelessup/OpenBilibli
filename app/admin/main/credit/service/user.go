package service

import (
	"context"

	account "github.com/namelessup/bilibili/app/service/main/account/api"
	"github.com/namelessup/bilibili/library/log"
)

// UserInfo get account info.
func (s *Service) UserInfo(c context.Context, mid int64) (res *account.InfoReply, err error) {
	if res, err = s.accDao.RPCInfo(c, mid); err != nil {
		log.Error("s.accDao.RPCInfo error (%v)", err)
	}
	return
}
