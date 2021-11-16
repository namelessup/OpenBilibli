package dao

import (
	"github.com/namelessup/bilibili/app/interface/live/app-interface/conf"
	avApi "github.com/namelessup/bilibili/app/service/live/av/api/liverpc"
	fansMedalApi "github.com/namelessup/bilibili/app/service/live/fans_medal/api/liverpc"
	liveDataApi "github.com/namelessup/bilibili/app/service/live/live_data/api/liverpc"
	liveUserApi "github.com/namelessup/bilibili/app/service/live/live_user/api/liverpc"
	rankdbApi "github.com/namelessup/bilibili/app/service/live/rankdb/api/liverpc"
	relationApi "github.com/namelessup/bilibili/app/service/live/relation/api/liverpc"
	roomApi "github.com/namelessup/bilibili/app/service/live/room/api/liverpc"
	roomExApi "github.com/namelessup/bilibili/app/service/live/room_ex/api/liverpc"
	bvcApi "github.com/namelessup/bilibili/app/service/live/third_api/bvc"
	skyHorseApi "github.com/namelessup/bilibili/app/service/live/third_api/skyhorse"
	userExtApi "github.com/namelessup/bilibili/app/service/live/userext/api/liverpc"
	xuserApi "github.com/namelessup/bilibili/app/service/live/xuser/api/grpc/v1"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
)

// RoomApi liveRpc room-service api
var RoomApi *roomApi.Client

// AvApi liveRpc room-service api
var AvApi *avApi.Client

// RoomRawApi liveRpc room-service api
var RoomRawApi *liverpc.Client

// LiveUserApi liveRpc room-service api
var LiveUserApi *liveUserApi.Client

// RelationApi liveRpc room-service api
var RelationApi *relationApi.Client

// BvcApi liveRpc room-service api
var BvcApi *bvcApi.Client

// SkyHorseApi ... liveRpc room-service api
var SkyHorseApi *skyHorseApi.Client

// UserExtApi liveRpc room-service api
var UserExtApi *userExtApi.Client

// LiveDataApi liveRpc room-service api
var LiveDataApi *liveDataApi.Client

// RoomExtApi liveRpc room-service api
var RoomExtApi *roomExApi.Client

// FansMedalApi liveRpc room-service api
var FansMedalApi *fansMedalApi.Client

// RankdbApi liveRpc rankdb-service api
var RankdbApi *rankdbApi.Client

// RankdbApi liveRpc rankdb-service api
var XuserApi *xuserApi.Client

// InitAPI init all service APIs
func InitAPI() {
	RoomApi = roomApi.New(getConf("room"))
	AvApi = avApi.New(getConf("av"))
	RoomExtApi = roomExApi.New(getConf("roomex"))
	LiveUserApi = liveUserApi.New(getConf("live_user"))
	RelationApi = relationApi.New(getConf("relation"))
	BvcApi = bvcApi.New(conf.Conf.HttpClient, getBvcConf("host"), getBvcConf("mock"))
	SkyHorseApi = skyHorseApi.New(conf.Conf.HttpClient)
	RoomRawApi = liverpc.NewClient(getConf("room"))
	UserExtApi = userExtApi.New(getConf("userext"))
	LiveDataApi = liveDataApi.New(getConf("livedata"))
	FansMedalApi = fansMedalApi.New(getConf("fansmedal"))
	RankdbApi = rankdbApi.New(getConf("rankdb"))
}

func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRpc
	if c != nil {
		return c[appName]
	}
	return nil
}

func getBvcConf(name string) string {
	c := conf.Conf.Bvc
	if c == nil {
		return ""
	}
	if _, ok := c[name]; ok {
		return c[name]
	}
	return ""
}
