package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/feedback/conf"
	"github.com/namelessup/bilibili/app/interface/main/feedback/dao"
	locrpc "github.com/namelessup/bilibili/app/service/main/location/rpc/client"
)

// Service struct.
type Service struct {
	// dao
	dao *dao.Dao
	// conf
	c *conf.Config
	// rpc
	locationRPC *locrpc.Service
}

// New new Tag service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
		// rpc
		locationRPC: locrpc.New(c.LocationRPC),
	}
	// init dao
	s.dao = dao.New(c)
	return
}

// Ping check server ok
func (s *Service) Ping(c context.Context) (err error) {
	return
}
