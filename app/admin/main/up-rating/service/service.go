package service

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/up-rating/conf"
	"github.com/namelessup/bilibili/app/admin/main/up-rating/dao"
	"github.com/namelessup/bilibili/app/admin/main/up-rating/dao/global"
)

// Service struct
type Service struct {
	conf  *conf.Config
	dao   *dao.Dao
	cache *Cache
}

// New fn
func New(c *conf.Config) (s *Service) {
	global.Init(c)
	s = &Service{
		conf:  c,
		dao:   dao.New(c),
		cache: NewCache(60),
	}
	return s
}

// Ping fn
func (s *Service) Ping(c context.Context) (err error) {
	return nil
}

// Close dao
func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
}
