package http

import (
	"net/http"

	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func sitemap(c *bm.Context) {
	logUA(c)

	res, err := srv.Sitemap(c, c.Request.Host)
	if err != nil {
		log.Error(err.Error())
		return
	}
	c.Writer.Header().Set("Content-Type", "text/xml;charset=utf-8")
	c.String(http.StatusOK, string(res))
}
