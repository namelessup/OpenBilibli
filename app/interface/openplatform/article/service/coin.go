package service

import (
	"context"

	"github.com/namelessup/bilibili/app/interface/openplatform/article/dao"
	coin "github.com/namelessup/bilibili/app/service/main/coin/model"
	"github.com/namelessup/bilibili/library/log"
)

// Coin get user coin number
func (s *Service) Coin(c context.Context, mid, aid int64, ip string) (res int64, err error) {
	var coins *coin.ArchiveUserCoins
	arg := coin.ArgCoinInfo{Mid: mid, Aid: aid, RealIP: ip, AvType: 2}
	if coins, err = s.coinRPC.ArchiveUserCoins(c, &arg); err != nil {
		dao.PromError("coin:获取投币数量")
		log.Error("s.coinRPC.ArchiveUserCoins(%+v) error(%+v)", arg, err)
		return
	}
	res = coins.Multiply
	return
}
