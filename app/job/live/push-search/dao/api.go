package dao

import (
	"github.com/namelessup/bilibili/app/job/live/push-search/conf"
	userApi "github.com/namelessup/bilibili/app/service/live/user/api/liverpc"
	relationApi "github.com/namelessup/bilibili/app/service/live/relation/api/liverpc"
	roomApi "github.com/namelessup/bilibili/app/service/live/room/api/liverpc"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
)

var UserApi *userApi.Client
var RelationApi *relationApi.Client
var RoomApi *roomApi.Client

// InitAPI init all service APIs
func InitAPI() {
	UserApi = userApi.New(getConf("user"))
	RelationApi = relationApi.New(getConf("relation"))
	RoomApi = roomApi.New(getConf("room"))
}

func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRpc
	if c != nil {
		return c[appName]
	}
	return nil
}