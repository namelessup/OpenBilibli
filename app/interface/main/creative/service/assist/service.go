package assist

import (
	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/account"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/assist"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/danmu"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/reply"
	"github.com/namelessup/bilibili/app/interface/main/creative/service"
)

// Service assist.
type Service struct {
	c      *conf.Config
	assist *assist.Dao
	reply  *reply.Dao
	dm     *danmu.Dao
	acc    *account.Dao
}

// New get assist service.
func New(c *conf.Config, rpcdaos *service.RPCDaos) *Service {
	s := &Service{
		c:      c,
		assist: assist.New(c),
		reply:  reply.New(c),
		dm:     danmu.New(c),
		acc:    rpcdaos.Acc,
	}
	return s
}
