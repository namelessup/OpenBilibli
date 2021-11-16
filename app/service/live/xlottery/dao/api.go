package dao

import (
	giftApi "github.com/namelessup/bilibili/app/service/live/gift/api/liverpc"
	rcApi "github.com/namelessup/bilibili/app/service/live/rc/api/liverpc"
	userExtApi "github.com/namelessup/bilibili/app/service/live/userext/api/liverpc"
	"github.com/namelessup/bilibili/app/service/live/xlottery/conf"
	account "github.com/namelessup/bilibili/app/service/main/account/rpc/client"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
)

// AccountApi liverpc user api
var AccountApi *account.Service3

// GiftApi liverpc gift api
var GiftApi *giftApi.Client

// RcApi rc api
var RcApi *rcApi.Client

// UserExtApi userext api
var UserExtApi *userExtApi.Client

// InitAPI init all service APIs
func InitAPI() {
	AccountApi = account.New3(nil)
	GiftApi = giftApi.New(getConf("gift"))
	RcApi = rcApi.New(getConf("rc"))
	UserExtApi = userExtApi.New(getConf("userext"))
}

func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRpc
	if c != nil {
		return c[appName]
	}
	return nil
}
