package service

import (
	"context"
	"github.com/namelessup/bilibili/app/service/live/wallet/dao"
	"github.com/namelessup/bilibili/app/service/live/wallet/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

type GetTidHandler struct {
}

func (handler *GetTidHandler) NeedCheckUid() bool {
	return false
}
func (handler *GetTidHandler) NeedTransactionMutex() bool {
	return false
}
func (handler *GetTidHandler) BizExecute(ws *WalletService, basicParam *model.BasicParam, uid int64, params ...interface{}) (v interface{}, err error) {
	serviceType, _ := params[0].(int32)
	if !model.IsValidServiceType(serviceType) {
		err = ecode.RequestErr
		return
	}

	callParams, _ := params[1].(string)
	if callParams == "" {
		err = ecode.RequestErr
		return
	}
	log.Info("getTid info : type:%d,callParams:%s", serviceType, callParams)
	tid := dao.GetTid(model.ServiceType(serviceType), callParams)

	v = model.GetTidResp(tid)

	return

}

func (s *Service) GetTid(c context.Context, basicParam *model.BasicParam, uid int64, params ...interface{}) (v interface{}, err error) {
	handler := GetTidHandler{}
	return s.execByHandler(&handler, c, basicParam, uid, params...)
}
