package http

import (
	"github.com/namelessup/bilibili/app/admin/ep/saga/conf"
	"github.com/namelessup/bilibili/app/admin/ep/saga/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// @params Empty
// @router get /ep/admin/saga/v1/data/teams
// @response TeamInfoResp
func queryTeams(c *bm.Context) {

	if _, err := getUsername(c); err != nil {
		c.JSON(nil, err)
		return
	}
	resp := &model.TeamInfoResp{
		Department: conf.Conf.Property.DeInfo,
		Business:   conf.Conf.Property.BuInfo,
	}
	c.JSON(resp, nil)
}
