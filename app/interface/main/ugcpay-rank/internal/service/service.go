package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/ugcpay-rank/internal/conf"
	"github.com/namelessup/bilibili/app/interface/main/ugcpay-rank/internal/dao"
)

// Service struct
type Service struct {
	dao *dao.Dao
}

// New init
func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(),
	}
	return s
}

// Ping Service
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close Service
func (s *Service) Close() {
	s.dao.Close()
}
