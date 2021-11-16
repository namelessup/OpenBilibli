package space

import (
	"context"
	"runtime"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/app-interface/conf"
	accdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/account"
	arcdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/archive"
	artdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/article"
	audiodao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/audio"
	bgmdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/bangumi"
	bplusdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/bplus"
	coindao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/coin"
	commdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/community"
	elecdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/elec"
	favdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/favorite"
	livedao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/live"
	memberdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/member"
	paydao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/pay"
	reldao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/relation"
	srchdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/search"
	shopdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/shop"
	spcdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/space"
	tagdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/tag"
	thumbupdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/thumbup"
	"github.com/namelessup/bilibili/library/log"
)

// Service is space service
type Service struct {
	c          *conf.Config
	arcDao     *arcdao.Dao
	spcDao     *spcdao.Dao
	accDao     *accdao.Dao
	coinDao    *coindao.Dao
	commDao    *commdao.Dao
	srchDao    *srchdao.Dao
	favDao     *favdao.Dao
	bgmDao     *bgmdao.Dao
	tagDao     *tagdao.Dao
	liveDao    *livedao.Dao
	elecDao    *elecdao.Dao
	artDao     *artdao.Dao
	audioDao   *audiodao.Dao
	relDao     *reldao.Dao
	bplusDao   *bplusdao.Dao
	shopDao    *shopdao.Dao
	thumbupDao *thumbupdao.Dao
	payDao     *paydao.Dao
	memberDao  *memberdao.Dao
	// chan
	mCh       chan func()
	tick      time.Duration
	BlackList map[int64]struct{}
}

// New new space
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:          c,
		arcDao:     arcdao.New(c),
		spcDao:     spcdao.New(c),
		accDao:     accdao.New(c),
		coinDao:    coindao.New(c),
		commDao:    commdao.New(c),
		srchDao:    srchdao.New(c),
		favDao:     favdao.New(c),
		bgmDao:     bgmdao.New(c),
		tagDao:     tagdao.New(c),
		liveDao:    livedao.New(c),
		elecDao:    elecdao.New(c),
		artDao:     artdao.New(c),
		audioDao:   audiodao.New(c),
		relDao:     reldao.New(c),
		bplusDao:   bplusdao.New(c),
		shopDao:    shopdao.New(c),
		thumbupDao: thumbupdao.New(c),
		payDao:     paydao.New(c),
		memberDao:  memberdao.New(c),
		// mc proc
		mCh:       make(chan func(), 1024),
		tick:      time.Duration(c.Tick),
		BlackList: make(map[int64]struct{}),
	}
	// video db
	for i := 0; i < runtime.NumCPU(); i++ {
		go s.cacheproc()
	}
	if c != nil && c.Space != nil {
		for _, mid := range c.Space.ForbidMid {
			s.BlackList[mid] = struct{}{}
		}
	}
	s.loadBlacklist()
	go s.blacklistproc()
	return
}

// addCache add archive to mc or redis
func (s *Service) addCache(f func()) {
	select {
	case s.mCh <- f:
	default:
		log.Warn("cacheproc chan full")
	}
}

// cacheproc write memcache and stat redis use goroutine
func (s *Service) cacheproc() {
	for {
		f := <-s.mCh
		f()
	}
}

// Ping check server ok
func (s *Service) Ping(c context.Context) (err error) {
	return
}

// loadBlacklist
func (s *Service) loadBlacklist() {
	list, err := s.spcDao.Blacklist(context.Background())
	if err != nil {
		log.Error("%+v", err)
		return
	}
	s.BlackList = list
}

func (s *Service) blacklistproc() {
	for {
		time.Sleep(s.tick)
		s.loadBlacklist()
	}
}
