package service

import (
	"context"
	"github.com/namelessup/bilibili/app/interface/main/shorturl/conf"
	shortdao "github.com/namelessup/bilibili/app/interface/main/shorturl/dao"
	"github.com/namelessup/bilibili/library/log"
)

// Service service struct
type Service struct {
	shortd *shortdao.Dao
}

// New new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		shortd: shortdao.New(c),
	}
	return
}

// Ping ping service.
func (s *Service) Ping(c context.Context) (err error) {
	if err = s.shortd.Ping(c); err != nil {
		log.Error("s.dao.Ping error(%v)", err)
	}
	return
}
