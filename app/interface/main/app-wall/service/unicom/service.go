package unicom

import (
	"sync"
	"time"

	"github.com/namelessup/bilibili/app/interface/main/app-wall/conf"
	accdao "github.com/namelessup/bilibili/app/interface/main/app-wall/dao/account"
	liveDao "github.com/namelessup/bilibili/app/interface/main/app-wall/dao/live"
	seqDao "github.com/namelessup/bilibili/app/interface/main/app-wall/dao/seq"
	shopDao "github.com/namelessup/bilibili/app/interface/main/app-wall/dao/shopping"
	unicomDao "github.com/namelessup/bilibili/app/interface/main/app-wall/dao/unicom"
	"github.com/namelessup/bilibili/app/interface/main/app-wall/model/unicom"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/stat/prom"
)

const (
	_initIPlimitKey = "iplimit_%v_%v"
)

type Service struct {
	c                *conf.Config
	dao              *unicomDao.Dao
	live             *liveDao.Dao
	seqdao           *seqDao.Dao
	accd             *accdao.Dao
	shop             *shopDao.Dao
	tick             time.Duration
	unicomIpCache    []*unicom.UnicomIP
	unicomIpSQLCache map[string]*unicom.UnicomIP
	operationIPlimit map[string]struct{}
	unicomPackCache  []*unicom.UserPack
	// infoc
	logCh      chan interface{}
	packCh     chan interface{}
	packLogCh  chan interface{}
	userBindCh chan interface{}
	// waiter
	waiter sync.WaitGroup
	// databus
	userbindPub *databus.Databus
	// prom
	pHit  *prom.Prom
	pMiss *prom.Prom
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:                c,
		dao:              unicomDao.New(c),
		live:             liveDao.New(c),
		seqdao:           seqDao.New(c),
		accd:             accdao.New(c),
		shop:             shopDao.New(c),
		tick:             time.Duration(c.Tick),
		unicomIpCache:    []*unicom.UnicomIP{},
		unicomIpSQLCache: map[string]*unicom.UnicomIP{},
		operationIPlimit: map[string]struct{}{},
		unicomPackCache:  []*unicom.UserPack{},
		// databus
		userbindPub: databus.New(c.UnicomDatabus),
		// infoc
		logCh:      make(chan interface{}, 1024),
		packCh:     make(chan interface{}, 1024),
		packLogCh:  make(chan interface{}, 1024),
		userBindCh: make(chan interface{}, 1024),
		// prom
		pHit:  prom.CacheHit,
		pMiss: prom.CacheMiss,
	}
	// now := time.Now()
	s.loadIPlimit(c)
	s.loadUnicomIP()
	// s.loadUnicomIPOrder(now)
	s.loadUnicomPacks()
	// s.loadUnicomFlow()
	go s.loadproc()
	go s.unicomInfocproc()
	go s.unicomPackInfocproc()
	go s.addUserPackLogproc()
	s.waiter.Add(1)
	go s.userbindConsumer()
	return
}

// cacheproc load cache
func (s *Service) loadproc() {
	for {
		time.Sleep(s.tick)
		// now := time.Now()
		s.loadUnicomIP()
		// s.loadUnicomIPOrder(now)
		s.loadUnicomPacks()
		// s.loadUnicomFlow()
	}
}
