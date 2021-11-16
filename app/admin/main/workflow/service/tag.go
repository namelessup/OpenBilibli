package service

import (
	"github.com/namelessup/bilibili/app/admin/main/workflow/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// tag method find tag meta of bid & tid
func (s *Service) tag(bid int8, tid int64) (tMeta *model.TagMeta, err error) {
	var ok bool
	tMeta = &model.TagMeta{}
	if _, ok = s.tagListCache[bid]; !ok {
		log.Error("wrong bid(%d)", bid)
		err = ecode.WkfBusinessNotFound
		return
	}
	if tMeta, ok = s.tagListCache[bid][tid]; !ok {
		log.Error("TagListCache not found bid(%d) tag_id(%d)", bid, tid)
		err = ecode.WkfTagNotFound
		return
	}
	return
}
