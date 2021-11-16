package service

import (
	"context"
	"github.com/namelessup/bilibili/app/job/main/figure/model"
)

const (
	_live = 11
)

// PayOrderInfo handle user coin chenage message
func (s *Service) PayOrderInfo(c context.Context, mid, money int64, merchant int8) (err error) {
	column := model.ACColumnPayMoney
	if merchant == _live {
		column = model.ACColumnPayLiveMoney
	}
	return s.figureDao.PayOrderInfo(c, column, mid, money)
}
