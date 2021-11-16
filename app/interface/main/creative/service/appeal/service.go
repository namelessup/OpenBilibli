package appeal

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/account"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/appeal"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/archive"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/tag"
	"github.com/namelessup/bilibili/app/interface/main/creative/service"
)

//Service struct
type Service struct {
	c         *conf.Config
	ap        *appeal.Dao
	arc       *archive.Dao
	acc       *account.Dao
	tag       *tag.Dao
	appealTag int64
}

//New get service
func New(c *conf.Config, rpcdaos *service.RPCDaos) *Service {
	s := &Service{
		c:         c,
		ap:        appeal.New(c),
		arc:       rpcdaos.Arc,
		acc:       rpcdaos.Acc,
		tag:       tag.New(c),
		appealTag: c.AppealTag,
	}
	return s
}

// Ping service
func (s *Service) Ping(c context.Context) (err error) {
	return
}

// Close dao
func (s *Service) Close() {
}
