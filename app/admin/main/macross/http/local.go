package http

import (
	"net/http"

	"github.com/namelessup/bilibili/app/admin/main/macross/conf"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// ping check server ok.
func ping(c *bm.Context) {
	if err := svr.Ping(c); err != nil {
		log.Error("resource service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// version check server version.
func version(c *bm.Context) {
	data := map[string]interface{}{
		"version": conf.Conf.Version,
	}
	c.JSONMap(data, nil)
}

// register for discovery
func register(c *bm.Context) {
	c.JSON(nil, nil)
}
