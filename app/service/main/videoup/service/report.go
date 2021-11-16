package service

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/service/main/videoup/model/archive"
	"github.com/namelessup/bilibili/library/log"
	xtime "github.com/namelessup/bilibili/library/time"
)

// ArcReport  add archive report.
func (s *Service) ArcReport(c context.Context, mid, aid int64, tp int8, reason, pics string, now time.Time) (err error) {
	var aa *archive.ArcReport
	if aa, err = s.arc.ArcReport(c, aid, mid); aa != nil {
		return
	}
	xNow := xtime.Time(now.Unix())
	aa = &archive.ArcReport{
		Mid:    mid,
		Aid:    aid,
		Type:   tp,
		Reason: reason,
		Pics:   pics,
		State:  archive.ArcReportNew,
		CTime:  xNow,
		MTime:  xNow,
	}
	if _, err = s.arc.AddArcReport(c, aa); err != nil {
		log.Error("s.arc.AddArcReport() error(%v)", err)
		return
	}
	return
}
