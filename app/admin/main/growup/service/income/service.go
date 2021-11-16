package income

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/growup/conf"
	upD "github.com/namelessup/bilibili/app/admin/main/growup/dao"
	incomeD "github.com/namelessup/bilibili/app/admin/main/growup/dao/income"
	"github.com/namelessup/bilibili/app/admin/main/growup/dao/message"
)

// Service struct
type Service struct {
	conf  *conf.Config
	dao   *incomeD.Dao
	msg   *message.Dao
	upDao *upD.Dao
}

// New fn
func New(c *conf.Config) (s *Service) {
	s = &Service{
		conf:  c,
		dao:   incomeD.New(c),
		msg:   message.New(c),
		upDao: upD.New(c),
	}
	return s
}

// Ping check dao health.
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close dao
func (s *Service) Close() {
	s.dao.Close()
}
