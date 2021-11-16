package grpc

import (
	pb "github.com/namelessup/bilibili/app/service/live/grpc-demo/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/grpc-demo/conf"
	svc "github.com/namelessup/bilibili/app/service/live/grpc-demo/service/v1"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// TODO

func Init(c *conf.Config) {
	s := warden.NewServer(nil)
	pb.RegisterGreeterServer(s.Server(), svc.NewGreeterService(c))
	_, err := s.Start()
	if err != nil {
		log.Error("grpc Start error(%v)", err)
		panic(err)
	}
}
