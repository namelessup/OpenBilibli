package wechat

import (
	"github.com/namelessup/bilibili/app/interface/main/web-goblin/conf"
	"github.com/namelessup/bilibili/app/interface/main/web-goblin/dao/wechat"
)

// Service struct .
type Service struct {
	c   *conf.Config
	dao *wechat.Dao
}

// New init wechat service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: wechat.New(c),
	}
	return s
}
