package dao

import (
	activityService "github.com/namelessup/bilibili/app/service/live/activity/api/liverpc"
	avService "github.com/namelessup/bilibili/app/service/live/av/api/liverpc"
	bannedService "github.com/namelessup/bilibili/app/service/live/banned_service/api/liverpc"
	fansMedalService "github.com/namelessup/bilibili/app/service/live/fans_medal/api/liverpc"
	"github.com/namelessup/bilibili/app/service/live/live-dm/conf"
	liveUserService "github.com/namelessup/bilibili/app/service/live/live_user/api/liverpc"
	rankdbService "github.com/namelessup/bilibili/app/service/live/rankdb/api/liverpc"
	rcService "github.com/namelessup/bilibili/app/service/live/rc/api/liverpc"
	roomService "github.com/namelessup/bilibili/app/service/live/room/api/liverpc"
	liveBroadCast "github.com/namelessup/bilibili/app/service/live/third_api/liveBroadcast"
	userService "github.com/namelessup/bilibili/app/service/live/user/api/liverpc"
	userextService "github.com/namelessup/bilibili/app/service/live/userext/api/liverpc"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
)

var (
	// BannedServiceClient liveRpc banner_service api
	BannedServiceClient *bannedService.Client
	// RoomServiceClient liveRpc room service api
	RoomServiceClient *roomService.Client
	// LiveUserServiceClient liveRpc liveUser service api
	LiveUserServiceClient *liveUserService.Client
	// AvServiceClient  liveRpc av service api
	AvServiceClient *avService.Client
	//FansMedalServiceClient liverpc fansmedal service api
	FansMedalServiceClient *fansMedalService.Client
	//ActivityServiceClient liverpc  activity service api
	ActivityServiceClient *activityService.Client
	//RcServiceClient liverpc rc service api
	RcServiceClient *rcService.Client
	//RankdbServiceClient liverpc rankdb service api
	RankdbServiceClient *rankdbService.Client
	//UserExtServiceClient liverpc userext service api
	UserExtServiceClient *userextService.Client
	//LiveBroadCastClient liverpc thirdApi
	LiveBroadCastClient *liveBroadCast.Client
	//UserClient liveRpc user api
	userClient *userService.Client
)

//InitAPI init all service APIS
func InitAPI() {
	BannedServiceClient = bannedService.New(getConf("banneDService"))
	RoomServiceClient = roomService.New(getConf("room"))
	LiveUserServiceClient = liveUserService.New(getConf("liveUser"))
	AvServiceClient = avService.New(getConf("avService"))
	FansMedalServiceClient = fansMedalService.New(getConf("fansMedal"))
	ActivityServiceClient = activityService.New(getConf("activity"))
	RcServiceClient = rcService.New(getConf("rc"))
	RankdbServiceClient = rankdbService.New(getConf("rankdbService"))
	UserExtServiceClient = userextService.New(getConf("userext"))
	LiveBroadCastClient = liveBroadCast.New(conf.Conf.HTTPClient)
	userClient = userService.New(getConf("user"))
}

func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRPC
	if c != nil {
		return c[appName]
	}
	return nil
}
