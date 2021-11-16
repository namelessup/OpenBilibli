package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/service/main/passport/conf"
	"github.com/namelessup/bilibili/app/service/main/passport/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	passportSvc *service.Service
	vfy         *verify.Verify
)

// Init fot init open service
func Init(c *conf.Config, s *service.Service) {
	passportSvc = s
	vfy = verify.New(c.Identify)
	// engine
	engIn := bm.DefaultServer(c.BM)
	innerRouter(c, engIn)
	// init inner server
	if err := engIn.Start(); err != nil {
		log.Error("bm.Start error(%v)", err)
		panic(err)
	}
}

func innerRouter(c *conf.Config, e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	group := e.Group("/x/internal/passport", vfy.Verify)
	{
		group.GET("/records/face", face)
		if c.Switch.LoginLogHBase {
			group.GET("records/loginlog", loginLog)
		}
		group.POST("/history/pwd/check", historyPwdCheck)
	}
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := passportSvc.Ping(c); err != nil {
		log.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// register support discovery.
func register(c *bm.Context) {
	c.JSON(map[string]struct{}{}, nil)
}
