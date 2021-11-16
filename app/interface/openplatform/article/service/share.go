package service

import (
	"context"

	artmdl "github.com/namelessup/bilibili/app/interface/openplatform/article/model"
	"github.com/namelessup/bilibili/library/ecode"
)

// AddShare adds share stats count.
func (s *Service) AddShare(c context.Context, id int64, mid int64, ip string) (err error) {
	var res *artmdl.Meta
	if res, err = s.ArticleMeta(c, id); (err != nil) || (res == nil) || (!res.IsNormal()) {
		err = ecode.NothingFound
		return
	}
	s.dao.PubShare(c, mid, id, ip)
	return
}
