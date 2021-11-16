package history

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/app/interface/main/tv/dao/archive"
	"github.com/namelessup/bilibili/app/interface/main/tv/dao/cms"
	"github.com/namelessup/bilibili/app/interface/main/tv/dao/history"
)

// Service .
type Service struct {
	conf   *conf.Config
	dao    *history.Dao
	cmsDao *cms.Dao
	arcDao *archive.Dao
}

// New .
func New(c *conf.Config) *Service {
	srv := &Service{
		conf:   c,
		dao:    history.New(c),
		cmsDao: cms.New(c),
		arcDao: archive.New(c),
	}
	return srv
}
