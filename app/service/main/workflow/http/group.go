package http

import (
	"github.com/namelessup/bilibili/app/service/main/workflow/model"
	bm "github.com/namelessup/bilibili/library/net/http/blademaster"
)

// deleteGroup delete group .
func deleteGroup(c *bm.Context) {
	p := &model.DeleteGroupParams{}
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(nil, wkfSvc.DeleteGroup(c, p))
}

// pubRefereeGroup delete group .
func pubRefereeGroup(c *bm.Context) {
	p := &model.PublicRefereeGroupParams{}
	if err := c.Bind(p); err != nil {
		return
	}
	c.JSON(nil, wkfSvc.PublicRefereeGroup(c, p))
}
