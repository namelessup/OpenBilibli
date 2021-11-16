package http

import (
	"github.com/namelessup/bilibili/app/service/live/xcaptcha/service/v1"
	"net/http"

	"github.com/namelessup/bilibili/app/service/live/xcaptcha/conf"
	"github.com/namelessup/bilibili/app/service/live/xcaptcha/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	srv             *service.Service
	vfy             *verify.Verify
	xCaptchaService *v1.XCaptchaService
)

// Init init
func Init(c *conf.Config) {
	srv = service.New(c)
	vfy = verify.New(c.Verify)
	engine := bm.DefaultServer(c.BM)
	xCaptchaService = v1.NewXCaptchaService(c)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm Start error(%v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	g := e.Group("/x/xcaptcha")
	{
		g.GET("/start", vfy.Verify, howToStart)
	}
	e.GET("/xlive/internal/xcaptcha/v1/xcaptcha/verify", captchaVerify)
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

// example for http request handler
func howToStart(c *bm.Context) {
	c.String(0, "Golang 大法好 !!!")
}
