package grpc

import (
	"fmt"
	pb "github.com/namelessup/bilibili/app/service/live/gift/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/gift/conf"
	"github.com/namelessup/bilibili/app/service/live/gift/dao"
	svc "github.com/namelessup/bilibili/app/service/live/gift/service/v1"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

//Init Init
func Init(c *conf.Config) {
	s := warden.NewServer(nil) // 酌情传入config
	dao.InitApi()
	gs := svc.NewGiftService(c)
	pb.RegisterGiftServer(s.Server(), gs)
	_, err := s.Start()
	if err != nil {
		log.Error("grpc Start error(%v)", err)
		panic(err)
	}
	fmt.Println("start")
	gs.TickerReloadGift()
}
