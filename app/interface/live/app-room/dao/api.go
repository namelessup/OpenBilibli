package dao

import (
	"github.com/namelessup/bilibili/app/interface/live/app-room/conf"
	userextApi "github.com/namelessup/bilibili/app/service/live/userext/api/liverpc"
	"github.com/namelessup/bilibili/library/net/rpc/liverpc"
)

// InitAPI init all service APIs
func InitAPI(dao *Dao) {
	dao.UserExtAPI = userextApi.New(getConf("userext"))
}

func getConf(appName string) *liverpc.ClientConfig {
	c := conf.Conf.LiveRpc
	if c != nil {
		return c[appName]
	}
	return nil
}
