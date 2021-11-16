package service

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/ep/tapd/conf"
	"github.com/namelessup/bilibili/app/admin/ep/tapd/dao"
	"github.com/namelessup/bilibili/library/sync/pipeline/fanout"

	"github.com/robfig/cron"
)

// Service struct
type Service struct {
	c            *conf.Config
	dao          *dao.Dao
	transferChan *fanout.Fanout
	cron         *cron.Cron
}

// New init.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:            c,
		dao:          dao.New(c),
		transferChan: fanout.New("cache", fanout.Worker(5), fanout.Buffer(10240)),
	}

	if s.c.Scheduler.Active {
		s.cron = cron.New()
		if err := s.cron.AddFunc(c.Scheduler.UpdateHookURLCacheTask, func() { s.dao.SaveEnableHookURLToCache() }); err != nil {
			panic(err)
		}
		s.cron.Start()
	}

	return
}

// Close Service.
func (s *Service) Close() {
	s.dao.Close()
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	err = s.dao.Ping(c)
	return
}
