package http

import (
	"net/http"

	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func moPing(c *bm.Context) {
	var err error
	if err = pingSvc.Ping(c); err != nil {
		log.Error("app-resource service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
