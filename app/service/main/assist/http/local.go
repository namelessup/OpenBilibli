package http

import (
	"net/http"

	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// ping check server ok.
func ping(c *bm.Context) {
	var err error
	if err = assSvc.Ping(c); err != nil {
		log.Error("assist-service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
