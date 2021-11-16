package dao

import (
	"github.com/namelessup/bilibili/app/admin/live/live-admin/conf"
	avApi "github.com/namelessup/bilibili/app/service/live/av/api/liverpc"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
)

// AvApi liveRpc room-service api
var AvApi *avApi.Client

// InitAPI init all service APIs
func InitAPI() {
	AvApi = avApi.New(getConf("av"))
}

func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRpc
	if c != nil {
		return c[appName]
	}
	return nil
}
