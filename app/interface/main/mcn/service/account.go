package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/mcn/dao/global"
	"github.com/namelessup/bilibili/app/interface/main/mcn/model/mcnmodel"
	"github.com/namelessup/bilibili/library/log"
)

//GetUpAccountInfo get account info
func (s *Service) GetUpAccountInfo(c context.Context, arg *mcnmodel.McnGetAccountReq) (result *mcnmodel.McnGetAccountReply, err error) {
	var data, e = global.GetInfo(c, arg.Mid)
	err = e
	if err != nil || data == nil {
		log.Error("get info fail, req=%+v, err=%+v", arg, err)
		return
	}

	result = new(mcnmodel.McnGetAccountReply)
	result.Mid = data.Mid
	result.Name = data.Name
	log.Info("query acount info ok, req=%+v, result=%+v", arg, result)
	return
}
