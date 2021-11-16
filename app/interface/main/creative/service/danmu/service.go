package danmu

import (
	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/account"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/archive"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/danmu"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/elec"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/subtitle"
	"github.com/namelessup/bilibili/app/interface/main/creative/service"
)

// Service danmu.
type Service struct {
	c    *conf.Config
	dm   *danmu.Dao
	arc  *archive.Dao
	acc  *account.Dao
	sub  *subtitle.Dao
	elec *elec.Dao
}

// New get danmu service.
func New(c *conf.Config, rpcdaos *service.RPCDaos) *Service {
	s := &Service{
		c:    c,
		dm:   danmu.New(c),
		acc:  rpcdaos.Acc,
		arc:  rpcdaos.Arc,
		sub:  rpcdaos.Sub,
		elec: elec.New(c),
	}
	return s
}
