package service

import (
	"context"

	"github.com/namelessup/bilibili/app/job/main/coin/dao"
	coinmdl "github.com/namelessup/bilibili/app/service/main/coin/model"
	memmdl "github.com/namelessup/bilibili/app/service/main/member/api"
	"github.com/namelessup/bilibili/library/log"
)

func (s *Service) addExp(c context.Context, mid int64, count float64, reason, ip string) (err error) {
	argExp := &memmdl.AddExpReq{
		Mid:     mid,
		Count:   count,
		Operate: "coin",
		Reason:  reason,
		Ip:      ip,
	}
	if count <= 0 {
		log.Errorv(c, log.KV("log", "add exp count < 0"), log.KV("mid", mid), log.KV("err", err), log.KV("reason", reason), log.KV("count", count))
		dao.PromError("exp:addExp0")
		return
	}
	if _, err = s.memRPC.UpdateExp(c, argExp); err != nil {
		log.Errorv(c, log.KV("log", "s.coinDao.IncrExp()"), log.KV("mid", mid), log.KV("err", err), log.KV("reason", reason), log.KV("count", count))
		dao.PromError("exp:addExp")
		return
	}
	return
}

func (s *Service) addCoinExp(c context.Context, mid, tp, number int64, ip string) (err error) {
	arg := &coinmdl.ArgAddUserCoinExp{Mid: mid, Business: tp, Number: number, RealIP: ip}
	if err = s.coinRPC.AddUserCoinExp(c, arg); err != nil {
		log.Errorv(c, log.KV("log", "AddUserCoinExp"), log.KV("err", err))
		dao.PromError("exp:addCoinExp")
	}
	return
}
