package ping

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/app-wall/conf"
	walldao "github.com/namelessup/bilibili/app/interface/main/app-wall/dao/wall"
)

type Service struct {
	wallDao *walldao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		wallDao: walldao.New(c),
	}
	return
}

// Ping is check server ping.
func (s *Service) Ping(c context.Context) (err error) {
	return s.wallDao.Ping(c)
}
