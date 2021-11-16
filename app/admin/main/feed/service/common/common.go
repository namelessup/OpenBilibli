package common

import (
	"github.com/namelessup/bilibili/app/admin/main/feed/conf"
	accdao "github.com/namelessup/bilibili/app/admin/main/feed/dao/account"
	arcdao "github.com/namelessup/bilibili/app/admin/main/feed/dao/archive"
	pgcdao "github.com/namelessup/bilibili/app/admin/main/feed/dao/pgc"
	showdao "github.com/namelessup/bilibili/app/admin/main/feed/dao/show"
	"github.com/namelessup/bilibili/library/log"
	httpx "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Service is search service
type Service struct {
	showDao    *showdao.Dao
	pgcDao     *pgcdao.Dao
	accDao     *accdao.Dao
	arcDao     *arcdao.Dao
	client     *httpx.Client
	managerURL string
}

// New new a search service
func New(c *conf.Config) (s *Service) {
	var (
		pgc *pgcdao.Dao
		err error
	)
	if pgc, err = pgcdao.New(c); err != nil {
		log.Error("pgcdao.New error(%v)", err)
		return
	}
	s = &Service{
		showDao:    showdao.New(c),
		pgcDao:     pgc,
		accDao:     accdao.New(c),
		arcDao:     arcdao.New(c),
		client:     httpx.NewClient(c.HTTPClient),
		managerURL: c.Host.Manager,
	}
	return
}
