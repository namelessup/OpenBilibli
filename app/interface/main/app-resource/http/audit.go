package http

import (
	"strconv"

	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func audit(c *bm.Context) {
	params := c.Request.Form
	buildStr := params.Get("build")
	mobiApp := params.Get("mobi_app")
	build, err := strconv.Atoi(buildStr)
	if err != nil {
		log.Error("stronv.ParseInt(%s) error(%v)", buildStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, auditSvc.Audit(c, mobiApp, build))
}
