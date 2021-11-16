package passport

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/account/conf"
	"github.com/namelessup/bilibili/app/interface/main/account/dao/passport"
	"github.com/namelessup/bilibili/library/net/metadata"
)

// Service struct of service.
type Service struct {
	// conf
	c *conf.Config

	passDao *passport.Dao
}

// New create service instance and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,

		passDao: passport.New(c),
	}
	return
}

// TestUserName is.
func (s *Service) TestUserName(ctx context.Context, name string, mid int64) error {
	ip := metadata.String(ctx, metadata.RemoteIP)
	return s.passDao.TestUserName(ctx, name, mid, ip)
}
