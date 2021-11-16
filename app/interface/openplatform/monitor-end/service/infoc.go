package service

import (
	"github.com/namelessup/bilibili/app/interface/openplatform/monitor-end/model"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/log/infoc"
)

// writeInfoc
func (s *Service) infocproc() {
	var (
		collectInfoc = infoc.New(s.c.CollectInfoc)
	)
	for {
		i, ok := <-s.infoCh
		if !ok {
			log.Warn("infoc proc exit")
			return
		}
		switch l := i.(type) {
		case model.CollectParams:
			collectInfoc.Info(l.Source, l.Product, l.Event, l.SubEvent, l.Code, l.ExtJSON, l.Mid, l.IP, l.Buvid, l.UserAgent)
		}
	}
}
