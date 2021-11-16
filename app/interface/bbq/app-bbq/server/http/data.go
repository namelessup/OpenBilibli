package http

import (
	"github.com/namelessup/bilibili/app/interface/bbq/app-bbq/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

func videoPlay(c *bm.Context) {
	uiLog(c, model.ActionPlay, nil)
}
