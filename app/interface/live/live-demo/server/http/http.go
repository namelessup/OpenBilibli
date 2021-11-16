// Package http is where http server init,
// including routes
package http

import (
	"net/http"

	pb "github.com/namelessup/bilibili/app/interface/live/live-demo/api/http"
	v2pb "github.com/namelessup/bilibili/app/interface/live/live-demo/api/http/v2"
	"github.com/namelessup/bilibili/app/interface/live/live-demo/conf"
	"github.com/namelessup/bilibili/app/interface/live/live-demo/dao"
	svc "github.com/namelessup/bilibili/app/interface/live/live-demo/service"
	"github.com/namelessup/bilibili/app/interface/live/live-demo/service/v2"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/auth"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	vfy     *verify.Verify
	midAuth *auth.Auth
)

// Init init
func Init(c *conf.Config) {
	dao.InitAPI()
	initMiddleware(c)
	engine := bm.DefaultServer(nil)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("bm Start error(%v)", err)
		panic(err)
	}
}

func initMiddleware(c *conf.Config) {
	vfy = verify.New(c.Verify)
	midAuth = auth.New(nil)
}

func route(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	g := e.Group("/x/live-demo")
	{
		g.GET("/start", vfy.Verify, howToStart)
	}
	midMap := map[string]bm.HandlerFunc{
		"auth":   midAuth.User,
		"guest":  midAuth.Guest,
		"verify": vfy.Verify}
	v2pb.RegisterV2FooService(e, v2.NewFooService(conf.Conf), midMap)
	v2pb.RegisterV2Foo2Service(e, v2.NewFoo2Service(conf.Conf), midMap)
	pb.RegisterFooService(e, svc.NewFooService(conf.Conf), midMap)
	pb.RegisterFoo2Service(e, svc.NewFoo2Service(conf.Conf), midMap)

	e.Inject(pb.PathFooGetInfo, midAuth.User)
	pb.RegisterFooBMServer(e, svc.NewFooService(conf.Conf))
}

func ping(c *bm.Context) {
	c.AbortWithStatus(http.StatusOK)
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}

// example for http request handler
func howToStart(c *bm.Context) {
	c.String(0, "Golang 大法好 !!!")
}
