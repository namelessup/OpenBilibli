package http

import (
	"net/http"

	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// ping check server ok.
func ping(c *bm.Context) {

	if vdaSvc.Ping(c) != nil {
		log.Error("videoup-admin service ping error")
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
