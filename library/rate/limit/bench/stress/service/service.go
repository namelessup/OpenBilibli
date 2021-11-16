package service

import (
	"github.com/namelessup/bilibili/library/rate/limit/bench/stress/conf"
)

// Service struct
type Service struct {
	c *conf.Config
}

// New init
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
	}
	return s
}
