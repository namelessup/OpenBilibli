package service

import (
	"context"

	"github.com/namelessup/bilibili/app/admin/main/mcn/conf"
	bfs "github.com/namelessup/bilibili/app/admin/main/mcn/dao/bfs"
	msg "github.com/namelessup/bilibili/app/admin/main/mcn/dao/msg"
	dao "github.com/namelessup/bilibili/app/admin/main/mcn/dao/up"
	videoup "github.com/namelessup/bilibili/app/admin/main/mcn/dao/videoup"
	"github.com/namelessup/bilibili/app/admin/main/mcn/model"
	"github.com/namelessup/bilibili/app/interface/main/mcn/tool/worker"
	accgrpc "github.com/namelessup/bilibili/app/service/main/account/api"
	arcgrpc "github.com/namelessup/bilibili/app/service/main/archive/api"
	memgrpc "github.com/namelessup/bilibili/app/service/main/member/api"

	"github.com/pkg/errors"
)

// Service struct
type Service struct {
	c       *conf.Config
	dao     *dao.Dao
	bfs     *bfs.Dao
	msg     *msg.Dao
	videoup *videoup.Dao
	msgMap  map[model.MSGType]*model.MSG
	memGRPC memgrpc.MemberClient
	accGRPC accgrpc.AccountClient
	arcGRPC arcgrpc.ArchiveClient
	worker  *worker.Pool
}

// New init
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:       c,
		dao:     dao.New(c),
		bfs:     bfs.New(c),
		msg:     msg.New(c),
		videoup: videoup.New(c),
		worker:  worker.New(nil),
	}
	var err error
	if s.accGRPC, err = accgrpc.NewClient(c.GRPCClient.Account); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	}
	if s.arcGRPC, err = arcgrpc.NewClient(c.GRPCClient.Archive); err != nil {
		panic(errors.WithMessage(err, "Failed to dial archive service"))
	}
	if s.memGRPC, err = memgrpc.NewClient(c.GRPCClient.Member); err != nil {
		panic(errors.WithMessage(err, "Failed to dial member service"))
	}
	s.setMsgTypeMap()
	return s
}

// Ping Service
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close Service
func (s *Service) Close() {
	s.dao.Close()
	s.worker.Close()
	s.worker.Wait()
}
