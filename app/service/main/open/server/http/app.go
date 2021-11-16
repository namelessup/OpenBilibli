package http

import (
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// Secret .
func secret(c *bm.Context) {
	sappKey := c.Request.Form.Get("sappkey")
	if sappKey == "" {
		c.JSON(nil, ecode.RequestErr)
		log.Error("sappkey is empty")
		return
	}
	appSecret, err := openSvc.Secret(c, sappKey)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(map[string]interface{}{
		"app_secret": appSecret,
	}, err)
}
