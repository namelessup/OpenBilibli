package http

import (
	"github.com/namelessup/bilibili/app/service/main/ugcpay-rank/internal/conf"
	"github.com/namelessup/bilibili/app/service/main/ugcpay-rank/internal/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	svc     *service.Service
	verifyM *verify.Verify
)

// Init init
func Init(s *service.Service) {
	svc = s
	verifyM = verify.New(conf.Conf.Verify)
	engine := bm.DefaultServer(conf.Conf.BM)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm Start error(%+v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(func(ctx *bm.Context) {})
	g := e.Group("/x/internal/ugcpay-rank")
	{
		g1 := g.Group("/elec", verifyM.Verify)
		{
			g1.GET("/month/up", elecMonthUP)
			g1.GET("/month", elecMonth)
			g1.GET("/all/av", elecAllAV)
		}
	}
}
