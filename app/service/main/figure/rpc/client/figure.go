package figure

import (
	"context"

	"github.com/namelessup/bilibili/app/service/main/figure/model"
	"github.com/namelessup/bilibili/library/net/rpc"
)

const (
	_userFigure = "RPC.UserFigure"
)

const (
	_appid = "account.service.figure"
)

var (
	_noRes = &struct{}{}
)

// Service struct info.
type Service struct {
	client *rpc.Client2
}

// New create instance of service and return.
func New(c *rpc.ClientConfig) (s *Service) {
	s = &Service{}
	s.client = rpc.NewDiscoveryCli(_appid, c)
	return
}

// UserFigure get user figure & figure rank info.
func (s *Service) UserFigure(c context.Context, arg *model.ArgUserFigure) (res *model.FigureWithRank, err error) {
	res = &model.FigureWithRank{}
	err = s.client.Call(c, _userFigure, arg, res)
	return
}
