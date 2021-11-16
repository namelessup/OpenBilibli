package http

import (
	"time"

	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func reddot(c *bm.Context) {
	res := map[string]interface{}{
		"data": operatorSvc.Reddot(time.Now()),
	}
	returnDataJSON(c, res, nil)
}
