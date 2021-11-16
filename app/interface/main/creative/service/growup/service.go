package growup

import (
	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/archive"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/growup"
	"github.com/namelessup/bilibili/app/interface/main/creative/service"
)

//Service struct.
type Service struct {
	c      *conf.Config
	arc    *archive.Dao
	growup *growup.Dao
	p      *service.Public
}

//New get service.
func New(c *conf.Config, rpcdaos *service.RPCDaos, p *service.Public) *Service {
	s := &Service{
		c:      c,
		arc:    rpcdaos.Arc,
		growup: growup.New(c),
		p:      p,
	}
	return s
}
