package service

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/admin/main/mcn/model"
	"github.com/namelessup/bilibili/app/interface/main/mcn/conf"
	"github.com/namelessup/bilibili/app/interface/main/mcn/dao/bfs"
	"github.com/namelessup/bilibili/app/interface/main/mcn/dao/cache"
	"github.com/namelessup/bilibili/app/interface/main/mcn/dao/global"
	"github.com/namelessup/bilibili/app/interface/main/mcn/dao/mcndao"
	"github.com/namelessup/bilibili/app/interface/main/mcn/dao/msg"
	"github.com/namelessup/bilibili/app/interface/main/mcn/tool/worker"

	"github.com/namelessup/bilibili/app/interface/main/mcn/dao/datadao"

	"github.com/bluele/gcache"
)

// Service struct
type Service struct {
	c             *conf.Config
	mcndao        *mcndao.Dao
	bfsdao        *bfs.Dao
	notifych      chan func()
	msg           *msg.Dao
	msgMap        map[model.MSGType]*model.MSG
	worker        *worker.Pool
	uniqueChecker *UniqueCheck
	datadao       *datadao.Dao
}

// New init
func New(c *conf.Config) (s *Service) {
	var localcache = gcache.New(c.RankCache.Size).Simple().Build()
	global.Init(c)
	s = &Service{
		c:             c,
		mcndao:        mcndao.New(c, localcache),
		bfsdao:        bfs.New(c),
		notifych:      make(chan func(), 10240),
		msg:           msg.New(c),
		worker:        worker.New(nil),
		uniqueChecker: NewUniqueCheck(),
		datadao:       datadao.New(c),
	}
	s.datadao.Client.Debug = true
	s.refreshCache()
	s.setMsgTypeMap()
	go s.cacheproc()
	return s
}

// Ping Service
func (s *Service) Ping(c context.Context) (err error) {
	return nil
}

// Close Service
func (s *Service) Close() {
	s.worker.Close()
	s.worker.Wait()
	s.mcndao.Close()
}

func (s *Service) refreshCache() {
	cache.LoadCache()
	s.loadMcnUniqueCache()
}

func (s *Service) cacheproc() {
	for {
		time.Sleep(5 * time.Minute)
		s.refreshCache()
	}
}
