package account

import (
	"context"
	"time"

	acc "github.com/namelessup/bilibili/app/service/main/account/api"
	mAcc "github.com/namelessup/bilibili/app/service/main/account/model"
	vip "github.com/namelessup/bilibili/app/service/main/vip/model"
	rpc "github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-sales/dao"
	"github.com/namelessup/bilibili/library/ecode"
)

//Checker 检查用户信息
type Checker struct {
	dao   *dao.Dao
	Users *acc.CardsReply
}

//New 新建一个用户检查类
func New(d *dao.Dao) *Checker {
	return &Checker{
		dao: d,
	}
}

//Check 检查用户条件
func (ac *Checker) Check(ctx context.Context, req *rpc.CreateOrdersRequest) (ee []ecode.Codes, err error) {
	l := len(req.Orders)
	ee = make([]ecode.Codes, l)
	uids := make([]int64, l)
	uidMap := make(map[int64]bool, l)
	i := 0
	for _, v := range req.Orders {
		if ok := uidMap[v.UID]; !ok {
			uids[i] = v.UID
			uidMap[v.UID] = true
			i++
		}
	}
	uids = uids[:i]
	if ac.Users, err = ac.dao.GetUserCards(ctx, uids); err != nil {
		return
	}
begin:
	for k, v := range req.Orders {
		if _, ok := ac.Users.Cards[v.UID]; !ok {
			ee[k] = ecode.TicketInvalidUser
			continue begin
		}
	}
	return
}

//GetUser 获取用户信息
func (ac *Checker) GetUser(mid int64) *mAcc.Card {
	u, ok := ac.Users.Cards[mid]
	if !ok {
		return nil
	}
	v := u.Vip
	if v.Type != vip.NotVip && v.Status != vip.VipStatusNotOverTime && v.DueDate <= time.Now().Unix() {
		return u
	}
	u.Vip.Type = vip.NotVip
	return u
}
