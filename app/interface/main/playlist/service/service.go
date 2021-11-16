package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/playlist/conf"
	"github.com/namelessup/bilibili/app/interface/main/playlist/dao"
	accclient "github.com/namelessup/bilibili/app/service/main/account/api"
	accwarden "github.com/namelessup/bilibili/app/service/main/account/api"
	accrpc "github.com/namelessup/bilibili/app/service/main/account/rpc/client"
	arcrpc "github.com/namelessup/bilibili/app/service/main/archive/api/gorpc"
	arcclient "github.com/namelessup/bilibili/app/service/main/archive/api"
	favrpc "github.com/namelessup/bilibili/app/service/main/favorite/api/gorpc"
	"github.com/namelessup/bilibili/app/service/main/filter/rpc/client"
	"github.com/namelessup/bilibili/library/cache"
	"github.com/namelessup/bilibili/library/log"
)

// Service service struct.
type Service struct {
	c   *conf.Config
	dao *dao.Dao
	// rpc
	fav    *favrpc.Service
	arc    *arcrpc.Service2
	acc    *accrpc.Service3
	filter *filter.Service
	// cache proc
	cache *cache.Cache
	// playlist power mids
	allowMids map[int64]struct{}
	maxSort   int64
	arcClient arcclient.ArchiveClient
	accClient accwarden.AccountClient
}

// New new service.
func New(c *conf.Config) *Service {
	s := &Service{
		c:       c,
		dao:     dao.New(c),
		fav:     favrpc.New2(c.FavoriteRPC),
		arc:     arcrpc.New2(c.ArchiveRPC),
		acc:     accrpc.New3(c.AccountRPC),
		filter:  filter.New(c.FilterRPC),
		cache:   cache.New(1, 1024),
		maxSort: c.Rule.MinSort + 4*c.Rule.SortStep*int64(c.Rule.MaxVideoCnt),
	}
	var err error
	if s.arcClient, err = arcclient.NewClient(c.ArcClient); err != nil {
		panic(err)
	}
	if s.accClient, err = accclient.NewClient(c.AccClient); err != nil {
		panic(err)
	}

	s.initMids()
	return s
}

func (s *Service) initMids() {
	tmp := make(map[int64]struct{}, len(s.c.Rule.PowerMids))
	for _, id := range s.c.Rule.PowerMids {
		tmp[id] = struct{}{}
	}
	s.allowMids = tmp
}

// Ping ping service.
func (s *Service) Ping(c context.Context) (err error) {
	if err = s.dao.Ping(c); err != nil {
		log.Error("s.dao.Ping error(%v)", err)
	}
	return
}
