package service

import (
	"context"
	"net/url"

	adminmodel "github.com/namelessup/bilibili/app/admin/main/mcn/model"
	"github.com/namelessup/bilibili/app/interface/main/mcn/model/mcnmodel"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// CreativeHandle .
func (s *Service) CreativeHandle(c context.Context, arg *mcnmodel.CreativeCommonReq, params url.Values, key string) (res interface{}, err error) {
	if !s.checkPermission(c, arg.McnMid, arg.UpMid, adminmodel.AttrDataPermitBit) {
		log.Warn("mcn permission insufficient, upmid=%d, mcnmid=%d", arg.UpMid, arg.McnMid)
		err = ecode.MCNPermissionInsufficient
		return
	}
	return s.datadao.HTTPDataHandle(c, params, key)
}
