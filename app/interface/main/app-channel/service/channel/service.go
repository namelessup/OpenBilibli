package channel

import (
	"context"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/app-card/model/card/live"
	"github.com/namelessup/bilibili/app/interface/main/app-card/model/card/operate"
	"github.com/namelessup/bilibili/app/interface/main/app-channel/conf"
	accdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/account"
	actdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/activity"
	arcdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/archive"
	artdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/article"
	audiodao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/audio"
	adtdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/audit"
	bgmdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/bangumi"
	carddao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/card"
	convergedao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/converge"
	gamedao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/game"
	livdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/live"
	locdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/location"
	rgdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/region"
	reldao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/relation"
	shopdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/shopping"
	specialdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/special"
	tabdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/tab"
	tagdao "github.com/namelessup/bilibili/app/interface/main/app-channel/dao/tag"
	"github.com/namelessup/bilibili/app/interface/main/app-channel/model/card"
	"github.com/namelessup/bilibili/app/interface/main/app-channel/model/channel"
	"github.com/namelessup/bilibili/app/interface/main/app-channel/model/tab"
)

// Service channel
type Service struct {
	c *conf.Config
	// dao
	acc   *accdao.Dao
	arc   *arcdao.Dao
	act   *actdao.Dao
	art   *artdao.Dao
	adt   *adtdao.Dao
	bgm   *bgmdao.Dao
	audio *audiodao.Dao
	rel   *reldao.Dao
	sp    *shopdao.Dao
	tg    *tagdao.Dao
	cd    *carddao.Dao
	ce    *convergedao.Dao
	g     *gamedao.Dao
	sl    *specialdao.Dao
	rg    *rgdao.Dao
	lv    *livdao.Dao
	loc   *locdao.Dao
	tab   *tabdao.Dao
	// tick
	tick time.Duration
	// cache
	cardCache         map[int64][]*card.Card
	cardPlatCache     map[string][]*card.CardPlat
	upCardCache       map[int64]*operate.Follow
	convergeCardCache map[int64]*operate.Converge
	gameDownloadCache map[int64]*operate.Download
	specialCardCache  map[int64]*operate.Special
	liveCardCache     map[int64][]*live.Card
	cardSetCache      map[int64]*operate.CardSet
	menuCache         map[int64][]*tab.Menu
	// new region list cache
	cachelist   map[string][]*channel.Region
	limitCache  map[int64][]*channel.RegionLimit
	configCache map[int64][]*channel.RegionConfig
	// audit cache
	auditCache map[string]map[int]struct{} // audit mobi_app builds
	// infoc
	logCh chan interface{}
}

// New channel
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:     c,
		arc:   arcdao.New(c),
		acc:   accdao.New(c),
		adt:   adtdao.New(c),
		art:   artdao.New(c),
		act:   actdao.New(c),
		bgm:   bgmdao.New(c),
		sp:    shopdao.New(c),
		tg:    tagdao.New(c),
		cd:    carddao.New(c),
		ce:    convergedao.New(c),
		g:     gamedao.New(c),
		sl:    specialdao.New(c),
		rg:    rgdao.New(c),
		audio: audiodao.New(c),
		lv:    livdao.New(c),
		rel:   reldao.New(c),
		loc:   locdao.New(c),
		tab:   tabdao.New(c),
		// tick
		tick: time.Duration(c.Tick),
		// cache
		cardCache:         map[int64][]*card.Card{},
		cardPlatCache:     map[string][]*card.CardPlat{},
		upCardCache:       map[int64]*operate.Follow{},
		convergeCardCache: map[int64]*operate.Converge{},
		gameDownloadCache: map[int64]*operate.Download{},
		specialCardCache:  map[int64]*operate.Special{},
		cachelist:         map[string][]*channel.Region{},
		limitCache:        map[int64][]*channel.RegionLimit{},
		configCache:       map[int64][]*channel.RegionConfig{},
		liveCardCache:     map[int64][]*live.Card{},
		cardSetCache:      map[int64]*operate.CardSet{},
		menuCache:         map[int64][]*tab.Menu{},
		// audit cache
		auditCache: map[string]map[int]struct{}{},
		// infoc
		logCh: make(chan interface{}, 1024),
	}
	s.loadCache()
	go s.loadCacheproc()
	go s.infocproc()
	return
}

func (s *Service) loadCacheproc() {
	for {
		time.Sleep(s.tick)
		s.loadCache()
	}
}

func (s *Service) loadCache() {
	now := time.Now()
	s.loadAuditCache()
	s.loadRegionlist()
	// card
	s.loadCardCache(now)
	s.loadConvergeCache()
	s.loadSpecialCache()
	s.loadLiveCardCache()
	s.loadGameDownloadCache()
	s.loadCardSetCache()
	s.loadMenusCache(now)
}

// Ping is check server ping.
func (s *Service) Ping(c context.Context) (err error) {
	if err = s.cd.PingDB(c); err != nil {
		return
	}
	return
}
