package http

import (
	"github.com/namelessup/bilibili/app/service/openplatform/ticket-item/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// @params VersionSearchParam
// @router get /openplatform/internal/ticket/item/version/search
// @response VersionSearchList
func versionSearch(c *bm.Context) {
	req := &model.VersionSearchParam{}
	if err := c.Bind(req); err != nil {
		return
	}
	c.JSON(itemSvc.VersionSearch(c, req))
}
