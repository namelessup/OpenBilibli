package service

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/vip/conf"
	"github.com/namelessup/bilibili/app/admin/main/vip/dao"
)

var (
	_maxTipLen     = 28
	_maxTitleLen   = 36
	_maxContentLen = 36
)

// Service struct
type Service struct {
	c         *conf.Config
	dao       *dao.Dao
	sendBcoin chan func()
}

// New init
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:         c,
		dao:       dao.New(c),
		sendBcoin: make(chan func(), 10240),
	}
	go s.bcoinproc()
	return s
}

// Ping check db live
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// func (s *Service) asyncBcoin(f func()) {
// 	select {
// 	case s.sendBcoin <- f:
// 	default:
// 		log.Warn("bcoinproc chan full")
// 	}
// }

func (s *Service) bcoinproc() {
	for {
		f := <-s.sendBcoin
		f()
	}
}

// Close Service
func (s *Service) Close() {
	s.dao.Close()
}
