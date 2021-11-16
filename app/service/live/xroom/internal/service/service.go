package service

import (
	"context"

	dav1 "github.com/namelessup/bilibili/app/service/live/dao-anchor/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/xroom/internal/dao"
	"github.com/namelessup/bilibili/library/conf/paladin"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// Service service.
type Service struct {
	ac *paladin.Map

	appConf *AppConf

	dao      *dao.Dao
	daClient *dav1.Client
}

// AppConf Conf
type AppConf struct {
}

//GrpcConf conf
type GrpcConf struct {
	Client *warden.ClientConfig
	Server *warden.ServerConfig
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	var gConf *GrpcConf
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&gConf); err != nil {
		panic(err)
	}
	dClient, err := dav1.NewClient(gConf.Client)
	if err != nil {
		panic(err)
	}

	s = &Service{
		ac:       ac,
		dao:      dao.New(),
		daClient: dClient,
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
