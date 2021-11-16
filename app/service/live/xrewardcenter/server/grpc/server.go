package grpc

import (
	pb "github.com/namelessup/bilibili/app/service/live/xrewardcenter/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/xrewardcenter/conf"
	"github.com/namelessup/bilibili/app/service/live/xrewardcenter/dao"
	svc "github.com/namelessup/bilibili/app/service/live/xrewardcenter/service/v1"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// TODO

// Init .
func Init(c *conf.Config) {
	s := warden.NewServer(nil)
	dao.InitAPI()
	pb.RegisterAnchorRewardServer(s.Server(), svc.NewAnchorTaskService(c))
	_, err := s.Start()
	if err != nil {
		log.Error("grpc Start error(%v)", err)
		panic(err)
	}
}
