package wechat

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/web-goblin/model/wechat"
	"github.com/namelessup/bilibili/library/log"
)

// Qrcode get qrcode from wechat.
func (s *Service) Qrcode(c context.Context, arg string) (qrcode []byte, err error) {
	var accessToken *wechat.AccessToken
	if accessToken, err = s.dao.AccessToken(c); err != nil {
		log.Error("Qrcode s.dao.AccessToken error(%v) arg(%s)", err, arg)
		return
	}
	qrcode, err = s.dao.Qrcode(c, accessToken.AccessToken, arg)
	return
}
