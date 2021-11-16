package dao

import (
	banned_api "github.com/namelessup/bilibili/app/service/live/banned_service/api/liverpc"
	fans_medal "github.com/namelessup/bilibili/app/service/live/fans_medal/api/liverpc"
	room_api "github.com/namelessup/bilibili/app/service/live/room/api/liverpc"

	"github.com/namelessup/bilibili/app/service/live/xuser/conf"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
)

// Dao dao
type Dao struct {
	c *conf.Config
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
	}
	return
}

// RoomAPI .
var RoomAPI *room_api.Client

// FansMedalAPI .
var FansMedalAPI *fans_medal.Client

// BannedAPI .
var BannedAPI *banned_api.Client

// InitAPI init all service APIs
func InitAPI() {
	RoomAPI = room_api.New(getConf("room"))
	FansMedalAPI = fans_medal.New(getConf("fans_medal"))
	BannedAPI = banned_api.New(getConf("banned"))
}

func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRPC
	if c != nil {
		return c[appName]
	}
	return nil
}
