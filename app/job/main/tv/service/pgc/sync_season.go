package pgc

import (
	"database/sql"
	"time"

	"github.com/namelessup/bilibili/app/job/main/tv/dao/lic"
	model "github.com/namelessup/bilibili/app/job/main/tv/model/pgc"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
)

// Sync modified season data to the license owner
func (s *Service) syncSeason() {
	defer s.waiter.Done()
	for {
		if s.daoClosed {
			log.Info("syncSeason DB closed!")
			return
		}
		modSeason, err := s.dao.ModSeason(ctx)
		if err == sql.ErrNoRows || len(modSeason) == 0 {
			log.Info("No modified data to pick from Season to audit")
			time.Sleep(time.Duration(s.c.Sync.Frequency.FreModSeason))
			continue
		}
		for _, v := range modSeason {
			if err = s.snSync(v); err != nil {
				s.addRetrySn(v)
			}
			s.dao.AuditSeason(ctx, int(v.ID)) // update season status after succ
		}
		time.Sleep(1 * time.Second) // break after each loop
	}
}

func (s *Service) snSync(sn *model.TVEpSeason) (err error) {
	cfg := s.c.Sync
	data := newLic(sn, cfg)
	data.XMLData.Service.Head.Count = 1
	res, err := s.licDao.CallRetry(ctx, cfg.API.UpdateURL, lic.PrepareXML(data))
	if res == nil {
		err = ecode.TvSyncErr
	}
	return
}
