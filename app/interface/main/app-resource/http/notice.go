package http

import (
	"strconv"

	"github.com/namelessup/bilibili/app/interface/main/app-resource/model"
	"github.com/namelessup/bilibili/library/ecode"
	"github.com/namelessup/bilibili/library/log"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// getNotice get notice data.
func getNotice(c *bm.Context) {
	params := c.Request.Form
	ver := params.Get("ver")
	buildStr := params.Get("build")
	mobiApp := params.Get("mobi_app")
	mobiApp = model.MobiAPPBuleChange(mobiApp)
	typeStr := params.Get("type")
	// check params
	build, err := strconv.Atoi(buildStr)
	if err != nil {
		log.Error("stronv.ParseInt(%s) error(%v)", buildStr, err)
		c.JSON(nil, ecode.RequestErr)
		return
	}
	typeInt, _ := strconv.Atoi(typeStr)
	device := params.Get("device")
	plat := model.Plat(mobiApp, device)
	// get
	data, version, err := ntcSvc.Notice(c, plat, build, typeInt, ver)
	res := map[string]interface{}{
		"data": data,
		"ver":  version,
	}
	c.JSONMap(res, err)
}
