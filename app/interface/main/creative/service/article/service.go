package article

import (
	"github.com/namelessup/bilibili/app/interface/main/creative/conf"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/account"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/activity"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/article"
	"github.com/namelessup/bilibili/app/interface/main/creative/dao/bfs"
	"github.com/namelessup/bilibili/app/interface/main/creative/service"
)

//Service struct.
type Service struct {
	c   *conf.Config
	art *article.Dao
	acc *account.Dao
	bfs *bfs.Dao
	act *activity.Dao
}

//New get service.
func New(c *conf.Config, rpcdaos *service.RPCDaos) *Service {
	s := &Service{
		c:   c,
		art: rpcdaos.Art,
		acc: rpcdaos.Acc,
		bfs: bfs.New(c),
		act: activity.New(c),
	}
	return s
}
