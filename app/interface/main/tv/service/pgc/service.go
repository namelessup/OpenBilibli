package pgc

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	appDao "github.com/namelessup/bilibili/app/interface/main/tv/dao/app"
	"github.com/namelessup/bilibili/app/interface/main/tv/dao/cms"
	"github.com/namelessup/bilibili/app/interface/main/tv/dao/pgc"
	"github.com/namelessup/bilibili/app/interface/main/tv/model"
)

var ctx = context.Background()

// Service .
type Service struct {
	appDao     *appDao.Dao
	cmsDao     *cms.Dao
	dao        *pgc.Dao
	conf       *conf.Config
	styleLabel map[int64][]*model.ParamStyle // style label
}

// New .
func New(c *conf.Config) *Service {
	srv := &Service{
		conf:       c,
		appDao:     appDao.New(c),
		cmsDao:     cms.New(c),
		dao:        pgc.New(c),
		styleLabel: make(map[int64][]*model.ParamStyle),
	}
	srv.styleCache()
	go srv.upStyleCache() // style label cache
	return srv
}
