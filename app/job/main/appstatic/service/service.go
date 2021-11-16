package service

import (
	"context"
	"sync"

	"github.com/namelessup/bilibili/app/job/main/appstatic/conf"
	"github.com/namelessup/bilibili/app/job/main/appstatic/dao/caldiff"
	"github.com/namelessup/bilibili/app/job/main/appstatic/dao/push"
	"github.com/namelessup/bilibili/library/log"
)

var ctx = context.Background()

// Service .
type Service struct {
	c         *conf.Config
	dao       *caldiff.Dao
	pushDao   *push.Dao
	waiter    *sync.WaitGroup
	daoClosed bool
}

// New creates a Service instance.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:       c,
		dao:     caldiff.New(c),
		pushDao: push.New(c),
		waiter:  new(sync.WaitGroup),
	}
	s.waiter.Add(1)
	go s.pushproc()
	s.waiter.Add(1)
	go s.calDiffproc()
	return
}

// Close releases resources which owned by the Service instance.
func (s *Service) Close() (err error) {
	log.Info("Close dao!")
	s.daoClosed = true
	log.Info("Wait waiter!")
	s.waiter.Wait()
	log.Info("appstatic-job has been closed.")
	return
}
