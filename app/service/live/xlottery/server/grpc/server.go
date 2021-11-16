package grpc

import (
	"fmt"

	"github.com/namelessup/bilibili/app/service/live/xlottery/dao"

	pb "github.com/namelessup/bilibili/app/service/live/xlottery/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/xlottery/conf"
	svc "github.com/namelessup/bilibili/app/service/live/xlottery/service/v1"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// Init .
func Init(c *conf.Config) *warden.Server {
	dao.InitAPI()
	s := warden.NewServer(nil) // 酌情传入config
	gs := svc.NewCapsuleService(c)
	pb.RegisterCapsuleServer(s.Server(), gs)
	pb.RegisterStormServer(s.Server(), svc.NewStromService(c))
	_, err := s.Start()
	if err != nil {
		log.Error("grpc Start error(%v)", err)
		panic(err)
	}
	fmt.Println("start")
	return s
}
