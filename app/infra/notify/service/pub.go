package service

import (
	"context"

	"github.com/namelessup/bilibili/app/infra/notify/model"
	"github.com/namelessup/bilibili/app/infra/notify/notify"
	"github.com/namelessup/bilibili/library/ecode"
)

// Pub pub message.
func (s *Service) Pub(c context.Context, arg *model.ArgPub) (err error) {
	pc, ok := s.pubConfs[key(arg.Group, arg.Topic)]
	if !ok {
		err = ecode.AccessDenied
		return
	}
	s.plock.RLock()
	pub, ok := s.pubs[key(arg.Group, arg.Topic)]
	s.plock.RUnlock()
	if !ok {
		pub, err = notify.NewPub(pc, s.c)
		if err != nil {
			return
		}
		s.plock.Lock()
		s.pubs[key(arg.Group, arg.Topic)] = pub
		s.plock.Unlock()
	}
	if !pub.Auth(arg.AppSecret) {
		err = ecode.AccessDenied
		return
	}
	err = pub.Send([]byte(arg.Key), []byte(arg.Msg))
	return
}
