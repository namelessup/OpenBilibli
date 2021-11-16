package http

import (
	"encoding/json"

	"github.com/namelessup/bilibili/app/interface/main/app-card/model/card/ai"
	"github.com/namelessup/bilibili/library/ecode"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func upRcmd(c *bm.Context) {
	params := c.Request.Form
	item := params.Get("item")
	var is []*ai.Item
	if err := json.Unmarshal([]byte(item), &is); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, feedSvc.UpRcmdCache(c, is))
}
