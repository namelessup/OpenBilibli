package service

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/answer/conf"
	"github.com/namelessup/bilibili/app/admin/main/answer/dao"
	"github.com/namelessup/bilibili/library/cache"
)

// Service struct of service.
type Service struct {
	c         *conf.Config
	dao       *dao.Dao
	eventChan *cache.Cache
}

// New create service instance and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:         c,
		dao:       dao.New(c),
		eventChan: cache.New(1, 10240),
	}
	s.generate(context.Background(), x, 0, len(x)-1)
	return
}

// Close dao.
func (s *Service) Close() {
	s.dao.Close()
}
