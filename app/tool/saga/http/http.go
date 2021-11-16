package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/tool/saga/conf"
	"github.com/namelessup/bilibili/app/tool/saga/service"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
	"github.com/namelessup/bilibili/library/net/http/blademaster/middleware/verify"
)

var (
	svc    *service.Service
	idfSvc *verify.Verify
)

// Init init http sever instance.
func Init(s *service.Service) {
	svc = s
	idfSvc = verify.New(nil)
	e := bm.DefaultServer(conf.Conf.BM)
	internalRouter(e)
	if err := e.Start(); err != nil {
		log.Error("xhttp.Serve error(%v)", err)
		panic(err)
	}
}

// internalRouter init internal router.
func internalRouter(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)

	group1 := e.Group("/x/internal/v2/saga/gitlab")
	{
		group1.POST("/comment", gitlabComment)
		group1.POST("/pipeline", gitlabPipeline)
		group1.POST("/mr", gitlabMR)
	}
	group2 := e.Group("/x/internal/v2/saga/api")
	{
		group2.POST("/contributors", buildContributors)
	}
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := svc.Ping(c); err != nil {
		log.Error("saga ping error(%+v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func register(c *bm.Context) {
	c.JSON(nil, nil)
}
