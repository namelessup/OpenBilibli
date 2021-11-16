package http

import (
	"fmt"

	"github.com/namelessup/bilibili/app/service/main/ugcpay/conf"
	"github.com/namelessup/bilibili/app/service/main/ugcpay/server/grpc"
	"github.com/namelessup/bilibili/app/service/main/ugcpay/service"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	srv *service.Service
	vfy *verify.Verify
)

// Init init
func Init(c *conf.Config, s *service.Service) {
	srv = s
	vfy = verify.New(c.Verify)
	grpc.New(nil, srv)
	engine := bm.DefaultServer(c.BM)
	route(engine)
	if err := engine.Start(); err != nil {
		panic(fmt.Sprintf("BM start err(%+v)", err))
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/x/internal/ugcpay", vfy.Verify)
	{
		g1 := g.Group("/asset")
		{
			g1.GET("", assetQuery)
			g1.POST("/register", assetRegister)
			g1.GET("/relation", assetRelation)
			g1.GET("/relation/detail", assetRelationDetail)
		}
	}
	g = e.Group("/x/internal/ugcpay")
	{
		g1 := g.Group("/trade")
		{
			g1.POST("refund", vfy.Verify, tradePayRefund)
			g1.POST("refunds", vfy.Verify, tradePayRefunds)
			g1.GET("/pay/callback", tradePayCallback)
			g1.GET("/pay/refund/callback", tradePayRefundCallback)
			g1.GET("/pay/recharge/callback", tradePayRechargeCallback)
		}
	}
}

func ping(c *bm.Context) {
}
