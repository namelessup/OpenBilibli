package audit

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	auditDao "github.com/namelessup/bilibili/app/interface/main/tv/dao/audit"
	"github.com/namelessup/bilibili/app/interface/main/tv/dao/cms"
)

// Service .
type Service struct {
	conf     *conf.Config
	auditDao *auditDao.Dao
	cmsDao   *cms.Dao
}

// New .
func New(c *conf.Config) *Service {
	srv := &Service{
		conf:     c,
		auditDao: auditDao.New(c),
		cmsDao:   cms.New(c),
	}
	return srv
}
