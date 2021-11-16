package service

import (
	"github.com/namelessup/bilibili/app/interface/live/lottery-interface/internal/conf"
	risk "github.com/namelessup/bilibili/app/service/live/live_riskcontrol/api/grpc/v1"
	storm "github.com/namelessup/bilibili/app/service/live/xlottery/api/grpc/v1"
	"github.com/namelessup/bilibili/library/log/infoc"
)

// Service struct
type Service struct {
	c                 *conf.Config
	Infoc             *infoc.Infoc
	StormClient       storm.StormClient
	IsForbiddenClient risk.IsForbiddenClient
}

// New init
func New(c *conf.Config) (s *Service) {
	sc, err := storm.NewClient(c.LongClient)
	if err != nil {
		panic(err)
	}
	isForbiddenClient, err := risk.NewClient(c.ShortClient)
	if err != nil {
		panic(err)
	}
	s = &Service{
		c:                 c,
		Infoc:             infoc.New(c.Infoc),
		StormClient:       sc.StormClient,
		IsForbiddenClient: isForbiddenClient,
	}
	return s
}

// Close Service
func (s *Service) Close() {
	s.Infoc.Close()
}

// ServiceInstance instance
var ServiceInstance *Service

// Init init
func Init(c *conf.Config) {
	ServiceInstance = New(c)
}
