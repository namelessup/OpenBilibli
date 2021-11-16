package grpc

import (
	pb "github.com/namelessup/bilibili/app/service/live/xcaptcha/api/grpc/v1"
	"github.com/namelessup/bilibili/app/service/live/xcaptcha/conf"
	svc "github.com/namelessup/bilibili/app/service/live/xcaptcha/service/v1"
	"github.com/namelessup/bilibili/library/log"
	"github.com/namelessup/bilibili/library/net/rpc/warden"
)

// Init gRpc Init
func Init(c *conf.Config) {
	s := warden.NewServer(nil) // 酌情传入config
	// 每个proto里定义的service添加一行
	pb.RegisterXCaptchaServer(s.Server(), svc.NewXCaptchaService(c))
	_, err := s.Start()
	if err != nil {
		log.Error("grpc Start error(%v)", err)
		panic(err)
	}
}
