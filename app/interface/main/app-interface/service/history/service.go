package history

import (
	"github.com/namelessup/bilibili/app/interface/main/app-interface/conf"
	historydao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/history"
	livedao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/live"
)

// Service service struct
type Service struct {
	historyDao *historydao.Dao
	liveDao    *livedao.Dao
}

// New new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		historyDao: historydao.New(c),
		liveDao:    livedao.New(c),
	}
	return
}
