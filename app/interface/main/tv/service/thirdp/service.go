package thirdp

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/app/interface/main/tv/dao/archive"
	cmsDao "github.com/namelessup/bilibili/app/interface/main/tv/dao/cms"
	"github.com/namelessup/bilibili/app/interface/main/tv/dao/thirdp"
	tpMdl "github.com/namelessup/bilibili/app/interface/main/tv/model/thirdp"
	xcache "github.com/namelessup/bilibili/library/cache"
)

var (
	ctx   = context.Background()
	cache *xcache.Cache
)

func init() {
	cache = xcache.New(1, 1024)
}

// Service .
type Service struct {
	dao        *thirdp.Dao
	cmsDao     *cmsDao.Dao
	arcDao     *archive.Dao
	conf       *conf.Config
	mangoRecom []*tpMdl.MangoParams // mango recom data
}

// New .
func New(c *conf.Config) *Service {
	srv := &Service{
		// dao
		dao:    thirdp.New(c),
		cmsDao: cmsDao.New(c),
		arcDao: archive.New(c),
		// config
		conf: c,
	}
	go srv.mangorproc() // load mango recom data
	srv.mangoR()
	return srv
}
