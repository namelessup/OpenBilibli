package grpc

import (
	dmg "github.com/namelessup/bilibili/app/service/live/live-dm/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/live-dm/conf"
	"github.com/namelessup/bilibili/app/service/live/live-dm/dao"
	dms "github.com/namelessup/bilibili/app/service/live/live-dm/service/v1"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

//Init 弹幕grpc 初始化
func Init(c *conf.Config) {
	dao.InitAPI()
	dao.InitGrpc(c)
	dao.InitIPdb()
	dao.InitDatabus(c)
	dao.InitLancer(c)
	dao.InitTitan()
	s := warden.NewServer(nil)
	dmg.RegisterDMServer(s.Server(), dms.NewDMService(c))
	_, err := s.Start()
	if err != nil {
		panic(err)
	}
}
