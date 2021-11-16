package tvvip

import (
	"github.com/namelessup/bilibili/app/interface/main/tv/conf"
	"github.com/namelessup/bilibili/app/service/main/tv/api"
	"github.com/namelessup/bilibili/library/log"
)

// Service .
type Service struct {
	conf        *conf.Config
	tvVipClient api.TVServiceClient
}

// New .
func New(c *conf.Config) *Service {
	tvVipClient, err := api.NewClient(c.TvVipClient)
	if err != nil {
		log.Error("client.Dial err(%v)", err)
		panic(err)
	}
	srv := &Service{
		conf:        c,
		tvVipClient: tvVipClient,
	}
	return srv
}
