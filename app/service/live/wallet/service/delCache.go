package service

import (
	"context"
	"github.com/namelessup/bilibili/app/service/live/wallet/model"
	mc "github.com/namelessup/bilibili/library/cache/memcache"
	"github.com/namelessup/bilibili/library/ecode"
)

// DelCacheHandler del cache handler
type DelCacheHandler struct {
}

// NeedCheckUid  need check uid
func (handler *DelCacheHandler) NeedCheckUid() bool {
	return true
}

// NeedTransactionMutex need mutex
func (handler *DelCacheHandler) NeedTransactionMutex() bool {
	return false
}

// BizExecute biz execute
func (handler *DelCacheHandler) BizExecute(ws *WalletService, basicParam *model.BasicParam, uid int64, params ...interface{}) (v interface{}, err error) {

	err = ws.s.dao.DelWalletCache(ws.c, uid)
	if err == mc.ErrNotFound {
		err = ecode.NothingFound
	}

	return
}

// DelCache del cache
func (s *Service) DelCache(c context.Context, basicParam *model.BasicParam, uid int64, params ...interface{}) (v interface{}, err error) {
	handler := DelCacheHandler{}
	return s.execByHandler(&handler, c, basicParam, uid, params...)
}
