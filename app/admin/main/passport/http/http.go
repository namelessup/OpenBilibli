package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/admin/main/passport/conf"
	"github.com/namelessup/bilibili/app/admin/main/passport/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/permit"
)

var (
	srv       *service.Service
	permitSvr *permit.Permit
)

// Init init
func Init(c *conf.Config) {
	srv = service.New(c)
	engine := bm.DefaultServer(c.BM)
	permitSvr = permit.New(c.Permit)
	router(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

func router(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	g := e.Group("/x/admin/passport")
	{
		g.GET("/userBindLog", permitSvr.Permit("LOG_USER_CONTACT_CHANGE"), userBindLog)
		g.GET("/user_bind_log/decrypt", decryptBindLog)
	}
}

func ping(c *bm.Context) {
	if err := srv.Ping(c); err != nil {
		log.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
