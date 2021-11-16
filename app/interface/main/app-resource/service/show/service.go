package show

import (
	"time"

	"github.com/namelessup/bilibili/app/interface/main/app-resource/conf"
	adtdao "github.com/namelessup/bilibili/app/interface/main/app-resource/dao/audit"
	resdao "github.com/namelessup/bilibili/app/interface/main/app-resource/dao/resource"
	tabdao "github.com/namelessup/bilibili/app/interface/main/app-resource/dao/tab"
	"github.com/namelessup/bilibili/app/interface/main/app-resource/model/show"
	"github.com/namelessup/bilibili/app/interface/main/app-resource/model/tab"
	resource "github.com/namelessup/bilibili/app/service/main/resource/model"
)

// Service is showtab service.
type Service struct {
	c *conf.Config
	//dao
	rdao        *resdao.Dao
	tdao        *tabdao.Dao
	adt         *adtdao.Dao
	tick        time.Duration
	tabCache    map[string][]*show.Tab
	limitsCahce map[int64][]*resource.SideBarLimit
	menuCache   []*tab.Menu
	abtestCache map[string]*resource.AbTest
	showTabMids map[int64]struct{}
	auditCache  map[string]map[int]struct{} // audit mobi_app builds
}

// New new a showtab service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:           c,
		rdao:        resdao.New(c),
		tdao:        tabdao.New(c),
		adt:         adtdao.New(c),
		tick:        time.Duration(c.Tick),
		tabCache:    map[string][]*show.Tab{},
		limitsCahce: map[int64][]*resource.SideBarLimit{},
		menuCache:   []*tab.Menu{},
		abtestCache: map[string]*resource.AbTest{},
		showTabMids: map[int64]struct{}{},
		auditCache:  map[string]map[int]struct{}{},
	}
	if err := s.loadCache(); err != nil {
		panic(err)
	}
	s.loadShowTabAids()
	go s.loadCacheproc()
	return
}
