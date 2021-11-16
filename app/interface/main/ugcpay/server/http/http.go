package http

import (
	"github.com/namelessup/bilibili/app/interface/main/ugcpay/conf"
	"github.com/namelessup/bilibili/app/interface/main/ugcpay/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/antispam"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
)

var (
	srv       *service.Service
	authM     *auth.Auth
	antispamM *antispam.Antispam
)

// Init init
func Init(c *conf.Config, s *service.Service) {
	srv = s
	authM = auth.New(c.Auth)
	antispamM = antispam.New(c.Antispam)
	engine := bm.DefaultServer(c.BM)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm Start error(%+v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	web := e.Group("/x/ugcpay", authM.UserWeb)
	{
		web.GET("/trade", tradeQuery)
		web.POST("/trade/confirm", antispamM.Handler(), tradeConfirm)
		web.POST("/trade/create", antispamM.Handler(), tradeCreate)
		web.POST("/trade/cancel", tradeCancel)
		web.GET("/income/asset/overview", antispamM.Handler(), incomeAssetOverview)
		web.GET("/income/asset/monthly", antispamM.Handler(), incomeAssetMonthly)
	}
	app := e.Group("/x/ugcpay/v1", authM.UserMobile)
	{
		app.GET("/trade", tradeQuery)
		app.POST("/trade/confirm", antispamM.Handler(), tradeConfirm)
		app.POST("/trade/create", antispamM.Handler(), tradeCreate)
		app.POST("/trade/cancel", tradeCancel)
		app.GET("/income/asset/overview", antispamM.Handler(), incomeAssetOverview)
		app.GET("/income/asset/monthly", antispamM.Handler(), incomeAssetMonthly)
	}
}

func ping(c *bm.Context) {
}
