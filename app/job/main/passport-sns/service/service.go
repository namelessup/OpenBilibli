package service

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/passport-sns/conf"
	"github.com/namelessup/bilibili/app/job/main/passport-sns/dao"
	"github.com/namelessup/bilibili/app/job/main/passport-sns/model"
	"github.com/namelessup/bilibili/library/queue/databus"
	"github.com/namelessup/bilibili/library/queue/databus/databusutil"
	"github.com/namelessup/bilibili/library/sync/pipeline/fanout"
)

// Service service.
type Service struct {
	c                 *conf.Config
	d                 *dao.Dao
	snsLogConsumer    *databus.Databus
	asoBinLogConsumer *databus.Databus
	group             *databusutil.Group
	snsChan           []chan *model.AsoAccountSns
	checkChan         []chan *model.AsoAccountSns
	cache             *fanout.Fanout
}

// New new a service instance.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:              c,
		d:              dao.New(c),
		snsLogConsumer: databus.New(c.DataBus.SnsLogSub),
		snsChan:        make([]chan *model.AsoAccountSns, c.SyncConf.ChanNum),
		checkChan:      make([]chan *model.AsoAccountSns, c.SyncConf.ChanNum),
		cache:          fanout.New("cache", fanout.Worker(10), fanout.Buffer(10240)),
	}
	go s.snsLogConsume()
	if c.SyncConf.IncSwitch {
		s.asoBinLogConsumer = databus.New(c.DataBus.AsoBinLogSub)
		s.group = databusutil.NewGroup(
			c.DatabusUtil,
			s.asoBinLogConsumer.Messages(),
		)
		s.asoBinLogConsume()
	}
	if c.SyncConf.FullSwitch {
		for i := 0; i < c.SyncConf.ChanNum; i++ {
			ch := make(chan *model.AsoAccountSns, c.SyncConf.ChanSize)
			s.snsChan[i] = ch
			go s.fullSyncSnsConsume(ch)
		}
		go s.fullSyncSns()
	}
	if c.SyncConf.CheckSwitch {
		for i := 0; i < c.SyncConf.ChanNum; i++ {
			ch := make(chan *model.AsoAccountSns, c.SyncConf.ChanSize)
			s.checkChan[i] = ch
			go s.checkConsume(ch)
		}
		go s.checkAll()
	}
	return
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return s.d.Ping(c)
}

// Close close service, including databus and outer service.
func (s *Service) Close() (err error) {
	s.d.Close()
	return
}

func parsePlatformStr(platform int) string {
	switch platform {
	case model.PlatformQQ:
		return model.PlatformQQStr
	case model.PlatformWEIBO:
		return model.PlatformWEIBOStr
	}
	return ""
}
