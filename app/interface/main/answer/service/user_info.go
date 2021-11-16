package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/answer/model"
	accoutCli "github.com/namelessup/bilibili/app/service/main/account/api"
	memModel "github.com/namelessup/bilibili/app/service/main/member/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/metadata"
)

// CheckBirthday check user had birthday info
func (s *Service) CheckBirthday(c context.Context, mid int64) (ok bool) {
	var (
		err error
		res *memModel.BaseInfo
		arg = &memModel.ArgMemberMid{Mid: mid, RemoteIP: metadata.String(c, metadata.RemoteIP)}
	)
	if res, err = s.memRPC.Base(c, arg); err != nil {
		log.Error("s.accRPC.Detail(mid:%d) error (%v)", mid, err)
		return
	}
	if res != nil && res.Birthday != 0 {
		birthday := res.Birthday.Time().Format("2006-01-02")
		if birthday != model.DefBirthday1 && birthday != model.DefBirthday2 {
			ok = true
		}
	}
	return
}

func (s *Service) accInfo(c context.Context, mid int64) (*accoutCli.Info, error) {
	accInfo, err := s.accountSvc.Info3(c, &accoutCli.MidReq{Mid: mid})
	if err != nil || accInfo == nil || accInfo.Info == nil {
		log.Error("s.accRPC.Info(%d) error(%v)", mid, err)
		return nil, ecode.AnswerAccCallErr
	}
	return accInfo.Info, nil
}
