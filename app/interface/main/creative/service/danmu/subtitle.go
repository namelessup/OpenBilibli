package danmu

import (
	"context"
	"github.com/namelessup/bilibili/app/interface/main/creative/model/danmu"

	"github.com/namelessup/bilibili/app/interface/main/dm2/model"
	"github.com/namelessup/bilibili/library/log"
)

// SubView fn
func (s *Service) SubView(c context.Context, aid int64, ip string) (ret *danmu.SubtitleSubjectReply, err error) {
	var sub *model.SubtitleSubjectReply
	if sub, err = s.sub.View(c, aid); err != nil {
		log.Error("s.sub.View err(%v) | aid(%d), ip(%s)", err, aid, ip)
		return
	}
	if sub != nil {
		ret = &danmu.SubtitleSubjectReply{
			AllowSubmit: sub.AllowSubmit,
			Lan:         sub.Lan,
			LanDoc:      sub.LanDoc,
		}
	}
	return
}
