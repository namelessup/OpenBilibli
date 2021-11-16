package favorite

import (
	"github.com/namelessup/bilibili/app/interface/main/app-interface/conf"
	artdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/article"
	audiodao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/audio"
	bangumidao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/bangumi"
	bplusdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/bplus"
	favdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/favorite"
	malldao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/mall"
	spdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/sp"
	ticketdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/ticket"
	topicdao "github.com/namelessup/bilibili/app/interface/main/app-interface/dao/topic"
)

// Service is favorite.
type Service struct {
	c *conf.Config
	// dao
	favDao     *favdao.Dao
	artDao     *artdao.Dao
	spDao      *spdao.Dao
	topicDao   *topicdao.Dao
	bplusDao   *bplusdao.Dao
	audioDao   *audiodao.Dao
	bangumiDao *bangumidao.Dao
	ticketDao  *ticketdao.Dao
	mallDao    *malldao.Dao
}

// New new favorite。
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
		// dao
		favDao:     favdao.New(c),
		topicDao:   topicdao.New(c),
		artDao:     artdao.New(c),
		spDao:      spdao.New(c),
		bplusDao:   bplusdao.New(c),
		audioDao:   audiodao.New(c),
		bangumiDao: bangumidao.New(c),
		ticketDao:  ticketdao.New(c),
		mallDao:    malldao.New(c),
	}
	return s
}
