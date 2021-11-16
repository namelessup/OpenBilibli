package dao

import (
	"github.com/namelessup/bilibili/app/job/live/xlottery/internal/conf"
	rcApi "github.com/namelessup/bilibili/app/service/live/rc/api/liverpc"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
)

// RcApi liverpc reward-service api
var RcApi *rcApi.Client

// InitAPI init all service APIs
func InitAPI() {
	RcApi = rcApi.New(getConf("rc"))
}

func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRpc
	if c != nil {
		return c[appName]
	}
	return nil
}
