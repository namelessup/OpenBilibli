package ping

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-resource/conf"
	pgdao "github.com/namelessup/bilibili/app/interface/main/app-resource/dao/plugin"
)

type Service struct {
	pgDao *pgdao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		pgDao: pgdao.New(c),
	}
	return
}

func (s *Service) Ping(c context.Context) (err error) {
	return s.pgDao.PingDB(c)
}
